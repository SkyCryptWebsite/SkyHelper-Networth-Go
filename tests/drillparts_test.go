package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestDrillPartsHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "TITANIUM_DRILL_1",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					DrillPartEngine:   "amber_polished_drill_engine",
					DrillPartFuelTank: "perfectly_cut_fuel_tank",
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"AMBER_POLISHED_DRILL_ENGINE": 250000000, "PERFECTLY_CUT_FUEL_TANK": 100000000},
			shouldApply:         true,
			expectedPriceChange: 100000000*constants.APPLICATION_WORTH["drillPart"] + 250000000*constants.APPLICATION_WORTH["drillPart"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "PERFECTLY_CUT_FUEL_TANK",
					Type:  "DRILL_PART",
					Price: 100000000 * constants.APPLICATION_WORTH["drillPart"],
					Count: 1,
				},
				{
					Id:    "AMBER_POLISHED_DRILL_ENGINE",
					Type:  "DRILL_PART",
					Price: 250000000 * constants.APPLICATION_WORTH["drillPart"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "TITANIUM_DRILL_1",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.DrillPartsHandler{}, testCases)
}
