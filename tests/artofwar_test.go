package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestArtOfWarHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "IRON_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					ArtOfWarCount: 1,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"THE_ART_OF_WAR": 20000000},
			shouldApply:         true,
			expectedPriceChange: 20000000 * constants.APPLICATION_WORTH["artOfWar"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "THE_ART_OF_WAR",
					Type:  "THE_ART_OF_WAR",
					Price: 20000000 * constants.APPLICATION_WORTH["artOfWar"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "IRON_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.ArtOfWarHandler{}, testCases)
}
