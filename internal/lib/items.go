package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

var cachedItems map[string]models.HypixelItem
var lastCached int64
var alreadyFetching bool

func GetItems(cache bool, cacheTime int64, retries int) (map[string]models.HypixelItem, error) {
	if cacheTime == 0 {
		cacheTime = 12 * 60 * 60
	}

	if retries == 0 {
		retries = 3
	}

	if lastCached != 0 && cache && (time.Now().Unix()-lastCached) < cacheTime {
		return cachedItems, nil
	}

	if alreadyFetching {
		for lastCached == 0 {
			time.Sleep(time.Millisecond * 100)
		}

		return cachedItems, nil
	}

	alreadyFetching = true
	defer func() { alreadyFetching = false }() // Ensure we reset the flag after fetching
	resp, err := http.Get("https://api.hypixel.net/v2/resources/skyblock/items")
	if err != nil {
		if retries > 0 {
			return GetItems(cache, cacheTime, retries-1)
		}

		return nil, fmt.Errorf("failed to fetch items: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if retries > 0 {
			return GetItems(cache, cacheTime, retries-1)
		}

		return nil, fmt.Errorf("failed to fetch prices: %s", resp.Status)
	}

	response := models.HypixelItemsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		// ? NOTE: Is there even point of this? Won't this fail every time?
		if retries > 0 {
			return GetItems(cache, cacheTime, retries-1)
		}

		return nil, fmt.Errorf("failed to decode prices: %v", err)
	}

	items := make(map[string]models.HypixelItem)
	for _, item := range response.Items {
		items[item.SkyBlockID] = item
	}

	cachedItems = items
	lastCached = time.Now().Unix()

	return items, nil
}

func GetItem(itemId string) *models.HypixelItem {
	items, err := GetItems(true, 69420, 3)
	if err != nil {
		return nil
	}

	item, exists := items[itemId]
	if !exists {
		return nil
	}

	return &item
}

func init() {
	// go GetItems(true, 69420, 3) // Initial fetch to populate cache
}
