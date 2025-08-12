package tests

import (
	"testing"

	"duckysolucky/skyhelper-networth-go/internal/calculators/handlers"
	"duckysolucky/skyhelper-networth-go/internal/constants"
	"duckysolucky/skyhelper-networth-go/internal/models"
)

func TestRuneHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "SUPERIOR_DRAGON_HELMET",
				ExtraAttributes: &models.ExtraAttributes{
					Runes: map[string]int{
						"GRAND_SEARING": 3,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"RUNE_GRAND_SEARING_3": 1200000000},
			shouldApply:         true,
			expectedPriceChange: 1200000000 * constants.APPLICATION_WORTH["runes"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "RUNE_GRAND_SEARING_3",
					Type:  "RUNE",
					Price: 1200000000 * constants.APPLICATION_WORTH["runes"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply with rune",
			item: &models.NetworthItem{
				ItemId: "RUNE",
				ExtraAttributes: &models.ExtraAttributes{
					Runes: map[string]int{
						"GRAND_SEARING": 3,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "LEATHER_CHESTPLATE",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.RuneHandler{}, testCases)
}
