package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestBoosterHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "FIGSTONE_AXE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Boosters: []string{"sweep"},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"SWEEP_BOOSTER": 100000},
			shouldApply:         true,
			expectedPriceChange: 100000 * constants.APPLICATION_WORTH["booster"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "SWEEP_BOOSTER",
					Type:  "BOOSTER",
					Price: 100000 * constants.APPLICATION_WORTH["booster"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "FIGSTONE_AXE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.BoosterHandler{}, testCases)
}
