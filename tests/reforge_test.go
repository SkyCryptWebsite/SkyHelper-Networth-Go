package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func TestReforgeHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "SUPERIOR_DRAGON_HELMET",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Modifier: "renowned",
				},
				Price:        100,
				Calculation:  []models.CalculationData{},
				SkyblockItem: &models.HypixelItem{},
			},
			prices:              map[string]float64{"DRAGON_HORN": 10000000},
			shouldApply:         true,
			expectedPriceChange: 10000000 * constants.APPLICATION_WORTH["reforge"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "DRAGON_HORN",
					Type:  "REFORGE",
					Price: 10000000 * constants.APPLICATION_WORTH["reforge"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply with accessory",
			item: &models.NetworthItem{
				ItemId: "BAT_TALISMAN",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Modifier: "strong",
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				SkyblockItem: &models.HypixelItem{
					Category: "ACCESSORY",
				},
			},
			prices:              map[string]float64{"DRAGON_HORN": 10000000},
			shouldApply:         false,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "LEATHER_CHESTPLATE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
				SkyblockItem:    &models.HypixelItem{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.ReforgeHandler{}, testCases)
}
