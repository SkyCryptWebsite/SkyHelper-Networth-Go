package tests

import (
	"testing"

	"duckysolucky/skyhelper-networth-go/internal/calculators/handlers"
	"duckysolucky/skyhelper-networth-go/internal/constants"
	"duckysolucky/skyhelper-networth-go/internal/models"
)

func TestPolarvoidBookHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "TITANIUM_DRILL_2",
				ExtraAttributes: &models.ExtraAttributes{
					Polarvoid: 5,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"POLARVOID_BOOK": 2500000},
			shouldApply:         true,
			expectedPriceChange: 5 * 2500000 * constants.APPLICATION_WORTH["polarvoidBook"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "POLARVOID_BOOK",
					Type:  "POLARVOID_BOOK",
					Price: 5 * 2500000 * constants.APPLICATION_WORTH["polarvoidBook"],
					Count: 5,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "TITANIUM_DRILL_2",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.PolarvoidBookHandler{}, testCases)
}
