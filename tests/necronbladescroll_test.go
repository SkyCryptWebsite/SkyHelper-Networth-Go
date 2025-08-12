package tests

import (
	"testing"

	"duckysolucky/skyhelper-networth-go/internal/calculators/handlers"
	"duckysolucky/skyhelper-networth-go/internal/constants"
	"duckysolucky/skyhelper-networth-go/internal/models"
)

func TestNecronBladeScrollHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "HYPERION",
				ExtraAttributes: &models.ExtraAttributes{
					AbilityScroll: []string{"WITHER_SHIELD_SCROLL", "IMPLOSION_SCROLL"},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"WITHER_SHIELD_SCROLL": 280000000, "IMPLOSION_SCROLL": 300000000},
			shouldApply:         true,
			expectedPriceChange: 280000000*constants.APPLICATION_WORTH["necronBladeScroll"] + 300000000*constants.APPLICATION_WORTH["necronBladeScroll"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "WITHER_SHIELD_SCROLL",
					Type:  "NECRON_SCROLL",
					Price: 280000000 * constants.APPLICATION_WORTH["necronBladeScroll"],
					Count: 1,
				},
				{
					Id:    "IMPLOSION_SCROLL",
					Type:  "NECRON_SCROLL",
					Price: 300000000 * constants.APPLICATION_WORTH["necronBladeScroll"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "IRON_SWORD",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.NecronBladeScrollHandler{}, testCases)
}
