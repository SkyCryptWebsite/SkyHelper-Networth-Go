package tests

import (
	"testing"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestArtOfPeaceHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "LEATHER_CHESTPLATE",
				ExtraAttributes: &models.ExtraAttributes{
					ArtOfPeaceApplied: 1,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"THE_ART_OF_PEACE": 50000000},
			shouldApply:         true,
			expectedPriceChange: 50000000 * constants.APPLICATION_WORTH["artOfPeace"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "THE_ART_OF_PEACE",
					Type:  "THE_ART_OF_PEACE",
					Price: 50000000 * constants.APPLICATION_WORTH["artOfPeace"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "LEATHER_CHESTPLATE",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.ArtOfPeaceHandler{}, testCases)
}
