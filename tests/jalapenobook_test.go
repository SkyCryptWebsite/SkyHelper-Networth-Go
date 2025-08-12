package tests

import (
	"testing"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestJalapenoBookHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "SOS_FLARE",
				ExtraAttributes: &models.ExtraAttributes{
					JalapenoCount: 1,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"JALAPENO_BOOK": 31000000},
			shouldApply:         true,
			expectedPriceChange: 31000000 * constants.APPLICATION_WORTH["jalapenoBook"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "JALAPENO_BOOK",
					Type:  "JALAPENO_BOOK",
					Price: 31000000 * constants.APPLICATION_WORTH["jalapenoBook"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "SOS_FLARE",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.JalapenoBookHandler{}, testCases)
}
