package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func TestOverclocker3000Handler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "THEORETICAL_HOE_POTATO_3",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Overclocker3000: 5,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"OVERCLOCKER_3000": 250000},
			shouldApply:         true,
			expectedPriceChange: 5 * 250000 * constants.APPLICATION_WORTH["overclocker3000"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "OVERCLOCKER_3000",
					Type:  "OVERCLOCKER_3000",
					Price: 5 * 250000 * constants.APPLICATION_WORTH["overclocker3000"],
					Count: 5,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "THEORETICAL_HOE_POTATO_3",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.Overclocker3000Handler{}, testCases)
}
