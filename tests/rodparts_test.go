package tests

import (
	"testing"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestRodPartsHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Line:   models.RodPart{Part: "titan_line"},
					Hook:   models.RodPart{Part: "hotspot_hook"},
					Sinker: models.RodPart{Part: "hotspot_sinker"},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"TITAN_LINE": 220000000, "HOTSPOT_HOOK": 16000000, "HOTSPOT_SINKER": 16000000},
			shouldApply:         true,
			expectedPriceChange: 220000000*constants.APPLICATION_WORTH["rodPart"] + 16000000*constants.APPLICATION_WORTH["rodPart"] + 16000000*constants.APPLICATION_WORTH["rodPart"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "TITAN_LINE",
					Type:  "ROD_PART",
					Price: 220000000 * constants.APPLICATION_WORTH["rodPart"],
					Count: 1,
				},
				{
					Id:    "HOTSPOT_HOOK",
					Type:  "ROD_PART",
					Price: 16000000 * constants.APPLICATION_WORTH["rodPart"],
					Count: 1,
				},
				{
					Id:    "HOTSPOT_SINKER",
					Type:  "ROD_PART",
					Price: 16000000 * constants.APPLICATION_WORTH["rodPart"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with soulbound parts",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Line: models.RodPart{Part: "titan_line"},
					Hook: models.RodPart{Part: "hotspot_hook", Soulbound: true},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"TITAN_LINE": 220000000, "HOTSPOT_HOOK": 16000000},
			shouldApply:         true,
			expectedPriceChange: 220000000*constants.APPLICATION_WORTH["rodPart"] + 16000000*constants.APPLICATION_WORTH["rodPart"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "TITAN_LINE",
					Type:  "ROD_PART",
					Price: 220000000 * constants.APPLICATION_WORTH["rodPart"],
					Count: 1,
				},
				{
					Id:        "HOTSPOT_HOOK",
					Type:      "ROD_PART",
					Price:     16000000 * constants.APPLICATION_WORTH["rodPart"],
					Count:     1,
					Soulbound: true,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.RodPartsHandler{}, testCases)
}
