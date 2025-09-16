package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func TestShensAuctionHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "CLOVER_HELMET",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Auction: 6,
					Bid:     6,
					Price:   2500000000,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:               map[string]float64{},
			shouldApply:          true,
			expectedNewBasePrice: 2500000000 * constants.APPLICATION_WORTH["shensAuctionPrice"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "CLOVER_HELMET",
					Type:  "SHENS_AUCTION",
					Price: 2500000000 * constants.APPLICATION_WORTH["shensAuctionPrice"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "LEATHER_CHESTPLATE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.ShensAuctionHandler{}, testCases)
}
