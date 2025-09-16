package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func TestFarmingForDummiesHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "THEORETICAL_HOE_CARROT_3",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					FarmingForDummies: 5,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"FARMING_FOR_DUMMIES": 2000000},
			shouldApply:         true,
			expectedPriceChange: 5 * 2000000 * constants.APPLICATION_WORTH["farmingForDummies"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "FARMING_FOR_DUMMIES",
					Type:  "FARMING_FOR_DUMMIES",
					Price: 5 * 2000000 * constants.APPLICATION_WORTH["farmingForDummies"],
					Count: 5,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "THEORETICAL_HOE_CARROT_3",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.FarmingForDummiesHandler{}, testCases)
}
