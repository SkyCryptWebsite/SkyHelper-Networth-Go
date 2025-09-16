package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestDivanPowderCoatingHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "DIVAN_DRILL",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					DivanPowderCoating: 1,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"DIVAN_POWDER_COATING": 100000000},
			shouldApply:         true,
			expectedPriceChange: 100000000 * constants.APPLICATION_WORTH["divanPowderCoating"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "DIVAN_POWDER_COATING",
					Type:  "DIVAN_POWDER_COATING",
					Price: 100000000 * constants.APPLICATION_WORTH["divanPowderCoating"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "DIVAN_DRILL",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.DivanPowderCoatingHandler{}, testCases)
}
