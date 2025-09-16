package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func TestWoodSingularityHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "TACTICIAN_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					WoodSingularityCount: 1,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"WOOD_SINGULARITY": 7000000},
			shouldApply:         true,
			expectedPriceChange: 7000000 * constants.APPLICATION_WORTH["woodSingularity"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "WOOD_SINGULARITY",
					Type:  "WOOD_SINGULARITY",
					Price: 7000000 * constants.APPLICATION_WORTH["woodSingularity"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "TACTICIAN_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.WoodSingularityHandler{}, testCases)
}
