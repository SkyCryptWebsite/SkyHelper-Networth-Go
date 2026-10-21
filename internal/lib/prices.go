package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// ErrPricesUnavailable means prices could not be fetched and no earlier copy is cached.
var ErrPricesUnavailable = errors.New("prices unavailable")

var pricesURL = "https://raw.githubusercontent.com/SkyHelperBot/Prices/main/pricesV2.json"

var httpClient = &http.Client{Timeout: 15 * time.Second}

type priceSnapshot struct {
	prices    map[string]float64
	fetchedAt time.Time
}

// A published snapshot is never written to again, so readers need no lock.
var cachedPrices atomic.Pointer[priceSnapshot]
var pricesFetchMu sync.Mutex
var lastPricesFailure atomic.Int64

// refreshFailureBackoff spaces out background refreshes while the upstream is failing.
const refreshFailureBackoff = 30 * time.Second

func GetPrices(cache bool, cacheTimeSeconds int64, retries int) (map[string]float64, error) {
	if cacheTimeSeconds == 0 {
		cacheTimeSeconds = 5 * 60
	}

	if retries == 0 {
		retries = 3
	}

	current := cachedPrices.Load()
	if cache && current != nil && time.Since(current.fetchedAt) < time.Duration(cacheTimeSeconds)*time.Second {
		return current.prices, nil
	}

	// With a stale copy in hand, serve it and refresh in the background, so no request waits on GitHub.
	if current != nil && cache {
		refreshPricesInBackground(retries)
		return current.prices, nil
	}

	pricesFetchMu.Lock()
	defer pricesFetchMu.Unlock()

	if latest := cachedPrices.Load(); latest != nil && latest != current {
		return latest.prices, nil
	}
	if current == nil && time.Since(time.Unix(0, lastPricesFailure.Load())) < refreshFailureBackoff {
		return nil, fmt.Errorf("%w: the last fetch failed less than %s ago", ErrPricesUnavailable, refreshFailureBackoff)
	}

	prices, err := fetchPrices(retries)
	if err != nil {
		lastPricesFailure.Store(time.Now().UnixNano())
		if current != nil {
			return current.prices, nil
		}

		return nil, fmt.Errorf("%w: %v", ErrPricesUnavailable, err)
	}

	cachedPrices.Store(&priceSnapshot{prices: prices, fetchedAt: time.Now()})
	return prices, nil
}

// refreshPricesInBackground starts one refresh at a time, and waits refreshFailureBackoff after a failed one.
func refreshPricesInBackground(retries int) {
	if time.Since(time.Unix(0, lastPricesFailure.Load())) < refreshFailureBackoff {
		return
	}
	if !pricesFetchMu.TryLock() {
		return
	}

	go func() {
		defer pricesFetchMu.Unlock()
		prices, err := fetchPrices(retries)
		if err != nil {
			lastPricesFailure.Store(time.Now().UnixNano())
			return
		}
		cachedPrices.Store(&priceSnapshot{prices: prices, fetchedAt: time.Now()})
	}()
}

func fetchPrices(retries int) (map[string]float64, error) {
	var lastErr error
	for attempt := 0; attempt <= retries; attempt++ {
		prices, err := fetchPricesOnce()
		if err == nil {
			return prices, nil
		}
		lastErr = err
	}

	return nil, lastErr
}

func fetchPricesOnce() (map[string]float64, error) {
	resp, err := httpClient.Get(pricesURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch prices: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch prices: %s", resp.Status)
	}

	prices := make(map[string]float64)
	if err := json.NewDecoder(resp.Body).Decode(&prices); err != nil {
		return nil, fmt.Errorf("failed to decode prices: %w", err)
	}

	prices["lastUpdated"] = float64(time.Now().Unix())
	return prices, nil
}
