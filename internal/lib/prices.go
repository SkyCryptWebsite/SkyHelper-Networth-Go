package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

var cachedPrices models.Prices
var alreadyFetchingPrice bool

func GetPrices(cache bool, cacheTimeSeconds int64, retries int) (*models.Prices, error) {
	if cacheTimeSeconds == 0 {
		cacheTimeSeconds = 5 * 60
	}

	if retries == 0 {
		retries = 3
	}

	prices := make(models.Prices)
	if cachedPrices["lastUpdated"] != 0 && cache && (time.Now().Unix()-int64(cachedPrices["lastUpdated"])) < cacheTimeSeconds {
		return &cachedPrices, nil
	}

	if alreadyFetchingPrice {
		for cachedPrices["lastUpdated"] == 0 {
			time.Sleep(time.Millisecond * 100)
		}

		return &cachedPrices, nil
	}

	alreadyFetchingPrice = true
	defer func() { alreadyFetchingPrice = false }() // Ensure we reset the flag after fetching
	resp, err := http.Get("https://raw.githubusercontent.com/SkyHelperBot/Prices/main/pricesV2.json")
	if err != nil {
		if retries > 0 {
			return GetPrices(cache, cacheTimeSeconds, retries-1)
		}

		return &prices, nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if retries > 0 {
			return GetPrices(cache, cacheTimeSeconds, retries-1)
		}

		return nil, fmt.Errorf("failed to fetch prices: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&prices)
	if err != nil {
		// ? NOTE: Is there even point of this? Won't this fail every time?
		if retries > 0 {
			return GetPrices(cache, cacheTimeSeconds, retries-1)
		}

		return nil, fmt.Errorf("failed to decode prices: %v", err)
	}

	cachedPrices = prices
	cachedPrices["lastUpdated"] = float64(time.Now().Unix())

	return &cachedPrices, nil
}

func init() {
	go GetPrices(true, 69420, 3) // Initial fetch to populate cache in background
}
