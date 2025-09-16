package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func TestDyeHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "POWER_WITHER_LEGGINGS",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					DyeItem: "DYE_WARDEN",
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"DYE_WARDEN": 90000000},
			shouldApply:         true,
			expectedPriceChange: 90000000 * constants.APPLICATION_WORTH["dye"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "DYE_WARDEN",
					Type:  "DYE",
					Price: 90000000 * constants.APPLICATION_WORTH["dye"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "POWER_WITHER_LEGGINGS",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.DyeHandler{}, testCases)
}
