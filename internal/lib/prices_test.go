package lib

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func usePricesServer(t *testing.T, handler http.HandlerFunc) {
	t.Helper()
	server := httptest.NewServer(handler)
	// Holding the fetch lock waits out any background refresh a previous test left running.
	pricesFetchMu.Lock()
	previousURL := pricesURL
	pricesURL = server.URL
	cachedPrices.Store(nil)
	lastPricesFailure.Store(0)
	pricesFetchMu.Unlock()
	t.Cleanup(func() {
		pricesFetchMu.Lock()
		defer pricesFetchMu.Unlock()
		server.Close()
		pricesURL = previousURL
		cachedPrices.Store(nil)
		lastPricesFailure.Store(0)
	})
}

func TestGetPricesWithoutCacheReturnsErrorOnFailure(t *testing.T) {
	usePricesServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	})

	prices, err := GetPrices(true, 60, 1)
	if !errors.Is(err, ErrPricesUnavailable) {
		t.Fatalf("expected ErrPricesUnavailable, got prices=%v err=%v", prices, err)
	}
}

func TestGetPricesServesStaleCopyWhenRefreshFails(t *testing.T) {
	var fail atomic.Bool
	usePricesServer(t, func(w http.ResponseWriter, r *http.Request) {
		if fail.Load() {
			// Close the connection without a response, like a network error.
			hijacker, _ := w.(http.Hijacker)
			conn, _, _ := hijacker.Hijack()
			_ = conn.Close()
			return
		}
		_, _ = w.Write([]byte(`{"HYPERION": 1000000000}`))
	})

	if _, err := GetPrices(true, 60, 1); err != nil {
		t.Fatal(err)
	}

	fail.Store(true)
	cachedPrices.Load().fetchedAt = time.Now().Add(-time.Hour)

	prices, err := GetPrices(true, 60, 1)
	if err != nil {
		t.Fatal(err)
	}
	if prices["HYPERION"] != 1000000000 {
		t.Fatalf("expected the stale price to be served, got %v", prices["HYPERION"])
	}
}

func TestGetPricesConcurrentRefresh(t *testing.T) {
	var requests atomic.Int32
	usePricesServer(t, func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		_, _ = w.Write([]byte(`{"HYPERION": 1000000000}`))
	})

	var wg sync.WaitGroup
	for range 50 {
		wg.Go(func() {
			prices, err := GetPrices(true, 60, 1)
			if err != nil {
				t.Error(err)
				return
			}
			if prices["HYPERION"] != 1000000000 {
				t.Errorf("unexpected price %v", prices["HYPERION"])
			}
		})
	}
	wg.Wait()

	if got := requests.Load(); got != 1 {
		t.Fatalf("expected one upstream fetch, got %d", got)
	}
}

func TestGetPricesRefreshesStaleCopyInBackground(t *testing.T) {
	var price atomic.Int64
	price.Store(1)
	release := make(chan struct{})
	var blocked atomic.Bool
	usePricesServer(t, func(w http.ResponseWriter, r *http.Request) {
		if blocked.Load() {
			<-release
		}
		_, _ = w.Write([]byte(`{"HYPERION": ` + strconv.FormatInt(price.Load(), 10) + `}`))
	})

	if _, err := GetPrices(true, 60, 1); err != nil {
		t.Fatal(err)
	}
	cachedPrices.Load().fetchedAt = time.Now().Add(-time.Hour)
	price.Store(2)
	blocked.Store(true)

	// The refresh is stuck upstream, yet the caller gets the stale copy straight away.
	start := time.Now()
	prices, err := GetPrices(true, 60, 1)
	if err != nil || prices["HYPERION"] != 1 {
		t.Fatalf("expected the stale price at once, got %v (%v)", prices["HYPERION"], err)
	}
	if waited := time.Since(start); waited > time.Second {
		t.Fatalf("caller waited %s for the background refresh", waited)
	}

	close(release)
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if cachedPrices.Load().prices["HYPERION"] == 2 {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("the background refresh never stored the new prices")
}

func TestGetPricesBacksOffAfterFailedFirstFetch(t *testing.T) {
	var requests atomic.Int32
	usePricesServer(t, func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		w.WriteHeader(http.StatusBadGateway)
	})

	for range 5 {
		if _, err := GetPrices(true, 60, 1); !errors.Is(err, ErrPricesUnavailable) {
			t.Fatalf("expected ErrPricesUnavailable, got %v", err)
		}
	}
	// One fetch with one retry; the next four calls wait out the backoff instead of fetching again.
	if got := requests.Load(); got != 2 {
		t.Fatalf("expected 2 upstream requests, got %d", got)
	}
}
