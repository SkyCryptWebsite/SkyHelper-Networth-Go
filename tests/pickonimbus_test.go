package tests

import (
	"testing"

	skycrypttypes "github.com/SkyCryptWebsite/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func TestPickonimbusHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "PICKONIMBUS",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					PickonimbusDurability: 1000,
				},
				Price:       50000,
				BasePrice:   50000,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         true,
			expectedPriceChange: -25000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "PICKONIMBUS_DURABLITY",
					Type:  "PICKONIMBUS",
					Price: -25000,
					Count: 1000,
				},
			},
		},
		{
			description: "Does not apply at full durability",
			item: &models.NetworthItem{
				ItemId:          "PICKONIMBUS",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{PickonimbusDurability: 2000},
				Price:           50000,
				BasePrice:       50000,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "PICKONIMBUS",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.PickonimbusHandler{}, testCases)
}
