package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestEnrichmentHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "ARTIFACT_OF_CONTROL",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					TalismanEnrichment: "magic_find",
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"TALISMAN_ENRICHMENT_MAGIC_FIND": 9000000, "TALISMAN_ENRICHMENT_CRITICAL_CHANCE": 8000000},
			shouldApply:         true,
			expectedPriceChange: 8000000 * constants.APPLICATION_WORTH["enrichment"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "MAGIC_FIND",
					Type:  "TALISMAN_ENRICHMENT",
					Price: 8000000 * constants.APPLICATION_WORTH["enrichment"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "ARTIFACT_OF_CONTROL",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.EnrichmentHandler{}, testCases)
}
