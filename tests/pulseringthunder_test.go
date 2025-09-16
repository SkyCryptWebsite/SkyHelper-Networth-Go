package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestPulseRingHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "PULSE_RING",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					ThunderCharge: 100000,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"THUNDER_IN_A_BOTTLE": 3000000},
			shouldApply:         true,
			expectedPriceChange: 2 * 3000000 * constants.APPLICATION_WORTH["thunderInABottle"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "THUNDER_IN_A_BOTTLE",
					Type:  "THUNDER_CHARGE",
					Price: 2 * 3000000 * constants.APPLICATION_WORTH["thunderInABottle"],
					Count: 2,
				},
			},
		},
		{
			description: "Applies correctly when above max",
			item: &models.NetworthItem{
				ItemId: "PULSE_RING",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					ThunderCharge: 5050000,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"THUNDER_IN_A_BOTTLE": 3000000},
			shouldApply:         true,
			expectedPriceChange: 100 * 3000000 * constants.APPLICATION_WORTH["thunderInABottle"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "THUNDER_IN_A_BOTTLE",
					Type:  "THUNDER_CHARGE",
					Price: 100 * 3000000 * constants.APPLICATION_WORTH["thunderInABottle"],
					Count: 100,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "PULSE_RING",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.PulseRingThunderHandler{}, testCases)
}
