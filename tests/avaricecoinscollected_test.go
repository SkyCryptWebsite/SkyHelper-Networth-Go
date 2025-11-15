package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func TestAvariceCoinsCollectedHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "CROWN_OF_AVARICE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					CollectedCoins: 500000000,
				},
				BasePrice:   100,
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:               map[string]float64{"CROWN_OF_AVARICE": 250000000, "CROWN_OF_AVARICE_1B": 4500000000},
			shouldApply:          true,
			expectedNewBasePrice: 2375000000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "CROWN_OF_AVARICE",
					Type:  "CROWN_OF_AVARICE",
					Price: 2375000000,
					Count: 500000000,
				},
			},
		},
		{
			description: "Applies correctly when maxed",
			item: &models.NetworthItem{
				ItemId: "CROWN_OF_AVARICE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					CollectedCoins: 1000000000,
				},
				BasePrice:   100,
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:               map[string]float64{"CROWN_OF_AVARICE": 250000000, "CROWN_OF_AVARICE_1B": 4500000000},
			shouldApply:          true,
			expectedNewBasePrice: 4500000000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "CROWN_OF_AVARICE",
					Type:  "CROWN_OF_AVARICE",
					Price: 4500000000,
					Count: 1000000000,
				},
			},
		},
		{
			description: "Applies correctly when over max",
			item: &models.NetworthItem{
				ItemId: "CROWN_OF_AVARICE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					CollectedCoins: 10000000000,
				},
				BasePrice:   100,
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:               map[string]float64{"CROWN_OF_AVARICE": 250000000, "CROWN_OF_AVARICE_1B": 4500000000},
			shouldApply:          true,
			expectedNewBasePrice: 4500000000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "CROWN_OF_AVARICE",
					Type:  "CROWN_OF_AVARICE",
					Price: 4500000000,
					Count: 1000000000,
				},
			},
		},
		{
			description: "Applies correctly BigInt",
			item: &models.NetworthItem{
				ItemId: "CROWN_OF_AVARICE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					CollectedCoins: 1000000000000,
				},
				BasePrice:   100,
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:               map[string]float64{"CROWN_OF_AVARICE": 250000000, "CROWN_OF_AVARICE_1B": 4500000000},
			shouldApply:          true,
			expectedNewBasePrice: 4500000000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "CROWN_OF_AVARICE",
					Type:  "CROWN_OF_AVARICE",
					Price: 4500000000,
					Count: 1000000000,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "IRON_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply with 0 coins collected",
			item: &models.NetworthItem{
				ItemId: "CROWN_OF_AVARICE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					CollectedCoins: 0,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.AvariceCoinsCollectedHandler{}, testCases)
}
