package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

var itemsURL = "https://api.hypixel.net/v2/resources/skyblock/items"

type itemsSnapshot struct {
	items map[string]models.HypixelItem
	// byID points into one shared copy per item, so GetItem needn't copy the struct on every lookup.
	byID      map[string]*models.HypixelItem
	fetchedAt time.Time
}

// A published snapshot is never written to again, so readers need no lock.
var cachedItems atomic.Pointer[itemsSnapshot]
var itemsFetchMu sync.Mutex
var lastItemsFailure atomic.Int64

func GetItems(cache bool, cacheTime int64, retries int) (map[string]models.HypixelItem, error) {
	if cacheTime == 0 {
		cacheTime = 12 * 60 * 60
	}

	if retries == 0 {
		retries = 3
	}

	current := cachedItems.Load()
	if cache && current != nil && time.Since(current.fetchedAt) < time.Duration(cacheTime)*time.Second {
		return current.items, nil
	}

	// With a stale copy in hand, serve it and refresh in the background, so no request waits on Hypixel.
	if current != nil && cache {
		refreshItemsInBackground(retries)
		return current.items, nil
	}

	itemsFetchMu.Lock()
	defer itemsFetchMu.Unlock()

	if latest := cachedItems.Load(); latest != nil && latest != current {
		return latest.items, nil
	}

	items, err := fetchItems(retries)
	if err != nil {
		if current != nil {
			return current.items, nil
		}

		return nil, err
	}

	storeItems(items)
	return items, nil
}

func storeItems(items map[string]models.HypixelItem) {
	byID := make(map[string]*models.HypixelItem, len(items))
	for id, item := range items {
		byID[id] = &item
	}
	cachedItems.Store(&itemsSnapshot{items: items, byID: byID, fetchedAt: time.Now()})
}

// refreshItemsInBackground starts one refresh at a time, and waits refreshFailureBackoff after a failed one.
func refreshItemsInBackground(retries int) {
	if time.Since(time.Unix(0, lastItemsFailure.Load())) < refreshFailureBackoff {
		return
	}
	if !itemsFetchMu.TryLock() {
		return
	}

	go func() {
		defer itemsFetchMu.Unlock()
		items, err := fetchItems(retries)
		if err != nil {
			lastItemsFailure.Store(time.Now().UnixNano())
			return
		}
		storeItems(items)
	}()
}

func fetchItems(retries int) (map[string]models.HypixelItem, error) {
	var lastErr error
	for attempt := 0; attempt <= retries; attempt++ {
		items, err := fetchItemsOnce()
		if err == nil {
			return items, nil
		}
		lastErr = err
	}
	return nil, lastErr
}

func fetchItemsOnce() (map[string]models.HypixelItem, error) {
	resp, err := httpClient.Get(itemsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch items: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch items: %s", resp.Status)
	}

	response := models.HypixelItemsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode items: %w", err)
	}

	items := make(map[string]models.HypixelItem, len(response.Items))
	for _, item := range response.Items {
		items[item.SkyBlockID] = item
	}

	return items, nil
}

// GetItem returns a shared, read-only item.
func GetItem(itemId string) *models.HypixelItem {
	if _, err := GetItems(true, 69420, 3); err != nil {
		return nil
	}

	snapshot := cachedItems.Load()
	if snapshot == nil {
		return nil
	}
	return snapshot.byID[itemId]
}
