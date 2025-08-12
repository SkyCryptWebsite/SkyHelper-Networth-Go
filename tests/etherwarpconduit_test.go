package tests

import (
	"testing"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestEtherWarpHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "ASPECT_OF_THE_VOID",
				ExtraAttributes: &models.ExtraAttributes{
					Ethermerge: 1,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"ETHERWARP_CONDUIT": 15000000},
			shouldApply:         true,
			expectedPriceChange: 15000000 * constants.APPLICATION_WORTH["etherwarp"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "ETHERWARP_CONDUIT",
					Type:  "ETHERWARP_CONDUIT",
					Price: 15000000 * constants.APPLICATION_WORTH["etherwarp"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "ASPECT_OF_THE_VOID",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.EtherwarpConduitHandler{}, testCases)
}
