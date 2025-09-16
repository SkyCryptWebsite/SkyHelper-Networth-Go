package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func TestPotatoBookHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "IRON_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					HotPotatoCount: 10,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"HOT_POTATO_BOOK": 80000},
			shouldApply:         true,
			expectedPriceChange: 10 * 80000 * constants.APPLICATION_WORTH["hotPotatoBook"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "HOT_POTATO_BOOK",
					Type:  "HOT_POTATO_BOOK",
					Price: 10 * 80000 * constants.APPLICATION_WORTH["hotPotatoBook"],
					Count: 10,
				},
			},
		},
		{
			description: "Applies correctly with Fuming Potato Books",
			item: &models.NetworthItem{
				ItemId: "IRON_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					HotPotatoCount: 15,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"HOT_POTATO_BOOK": 80000, "FUMING_POTATO_BOOK": 1400000},
			shouldApply:         true,
			expectedPriceChange: 10*80000*constants.APPLICATION_WORTH["hotPotatoBook"] + 5*1400000*constants.APPLICATION_WORTH["fumingPotatoBook"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "HOT_POTATO_BOOK",
					Type:  "HOT_POTATO_BOOK",
					Price: 10 * 80000 * constants.APPLICATION_WORTH["hotPotatoBook"],
					Count: 10,
				},
				{
					Id:    "FUMING_POTATO_BOOK",
					Type:  "FUMING_POTATO_BOOK",
					Price: 5 * 1400000 * constants.APPLICATION_WORTH["fumingPotatoBook"],
					Count: 5,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "TITANIUM_DRILL_2",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.PotatoBookHandler{}, testCases)
}
