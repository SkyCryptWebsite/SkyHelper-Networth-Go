package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func TestGemstonePowerScrollHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "FLORID_ZOMBIE_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					PowerAbilityScroll: "RUBY_POWER_SCROLL",
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"RUBY_POWER_SCROLL": 650000},
			shouldApply:         true,
			expectedPriceChange: 650000 * constants.APPLICATION_WORTH["gemstonePowerScroll"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "RUBY_POWER_SCROLL",
					Type:  "GEMSTONE_POWER_SCROLL",
					Price: 650000 * constants.APPLICATION_WORTH["gemstonePowerScroll"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "FLORID_ZOMBIE_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.GemstonePowerScrollHandler{}, testCases)
}
