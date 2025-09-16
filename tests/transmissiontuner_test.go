package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestTransmissionTunerHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "ASPECT_OF_THE_END",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					TunedTransmission: 4,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"TRANSMISSION_TUNER": 50000},
			shouldApply:         true,
			expectedPriceChange: 4 * 50000 * constants.APPLICATION_WORTH["tunedTransmission"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "TRANSMISSION_TUNER",
					Type:  "TRANSMISSION_TUNER",
					Price: 4 * 50000 * constants.APPLICATION_WORTH["tunedTransmission"],
					Count: 4,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "ASPECT_OF_THE_END",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.TransmissionTunerHandler{}, testCases)
}
