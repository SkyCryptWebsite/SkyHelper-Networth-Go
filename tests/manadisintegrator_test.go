package tests

import (
	"testing"

	"duckysolucky/skyhelper-networth-go/internal/calculators/handlers"
	"duckysolucky/skyhelper-networth-go/internal/constants"
	"duckysolucky/skyhelper-networth-go/internal/models"
)

func TestManaDisintegratorHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "WAND_OF_ATONEMENT",
				ExtraAttributes: &models.ExtraAttributes{
					ManaDisintegrator: 10,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"MANA_DISINTEGRATOR": 35000},
			shouldApply:         true,
			expectedPriceChange: 10 * 35000 * constants.APPLICATION_WORTH["manaDisintegrator"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "MANA_DISINTEGRATOR",
					Type:  "MANA_DISINTEGRATOR",
					Price: 10 * 35000 * constants.APPLICATION_WORTH["manaDisintegrator"],
					Count: 10,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "WAND_OF_ATONEMENT",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.ManaDisintegratorHandler{}, testCases)
}
