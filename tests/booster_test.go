package tests

import (
	"testing"

	skycrypttypes "github.com/SkyCryptWebsite/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
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
			description: "Applies correctly with booster tiers",
			item: &models.NetworthItem{
				ItemId: "FIGSTONE_AXE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					BoosterTiers: map[string]int{"sweep": 2},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"SWEEP_BOOSTER": 100000, "SWEEP_BOOSTER_UNCOMMON": 1000000},
			shouldApply:         true,
			expectedPriceChange: 1100000 * constants.APPLICATION_WORTH["booster"],
			expectedCalculation: []models.CalculationData{
				{Id: "SWEEP_BOOSTER", Type: "BOOSTER", Price: 100000 * constants.APPLICATION_WORTH["booster"], Count: 1},
				{Id: "SWEEP_BOOSTER_UNCOMMON", Type: "BOOSTER", Price: 1000000 * constants.APPLICATION_WORTH["booster"], Count: 1},
			},
		},
		{
			description: "Booster tiers fall back to the common booster ID",
			item: &models.NetworthItem{
				ItemId: "HUNTING_SCYTHE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					BoosterTiers: map[string]int{"hunting_wisdom": 1},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"HUNTING_WISDOM_BOOSTER_COMMON": 50000},
			shouldApply:         true,
			expectedPriceChange: 50000 * constants.APPLICATION_WORTH["booster"],
			expectedCalculation: []models.CalculationData{
				{Id: "HUNTING_WISDOM_BOOSTER_COMMON", Type: "BOOSTER", Price: 50000 * constants.APPLICATION_WORTH["booster"], Count: 1},
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
