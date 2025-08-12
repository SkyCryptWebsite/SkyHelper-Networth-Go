package tests

import (
	"testing"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestPocketSackInASackHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "LARGE_HUSBANDRY_SACK",
				ExtraAttributes: &models.ExtraAttributes{
					SackPss: 3,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"POCKET_SACK_IN_A_SACK": 12000000},
			shouldApply:         true,
			expectedPriceChange: 3 * 12000000 * constants.APPLICATION_WORTH["pocketSackInASack"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "POCKET_SACK_IN_A_SACK",
					Type:  "POCKET_SACK_IN_A_SACK",
					Price: 3 * 12000000 * constants.APPLICATION_WORTH["pocketSackInASack"],
					Count: 3,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "LARGE_HUSBANDRY_SACK",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.PocketSackInASackHandler{}, testCases)
}
