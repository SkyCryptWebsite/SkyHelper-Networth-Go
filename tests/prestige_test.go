package tests

import (
	"testing"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestPrestigeHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId:          "INFERNAL_FERVOR_LEGGINGS",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices: map[string]float64{
				"HEAVY_PEARL":             220000,
				"ESSENCE_CRIMSON":         1500,
				"KUUDRA_TEETH":            12000,
				"BURNING_FERVOR_LEGGINGS": 30000000,
				"HOT_FERVOR_LEGGINGS":     2000000,
				"FERVOR_LEGGINGS":         1000000,
			},
			shouldApply:         true,
			expectedPriceChange: (90100+25500+16345+4500)*1500*constants.APPLICATION_WORTH["essence"] + (12+12)*220000 + (80+50)*12000 + 30000000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "FIERY_FERVOR_LEGGINGS",
					Type:  "STARS",
					Price: 90100*1500*constants.APPLICATION_WORTH["essence"] + 12*220000,
					Count: 10,
				},
				{
					Id:    "FIERY_FERVOR_LEGGINGS",
					Type:  "PRESTIGE",
					Price: 25500*1500*constants.APPLICATION_WORTH["essence"] + 80*12000,
					Count: 1,
				},
				{
					Id:    "BURNING_FERVOR_LEGGINGS",
					Type:  "STARS",
					Price: 16345*1500*constants.APPLICATION_WORTH["essence"] + 12*220000,
					Count: 10,
				},
				{
					Id:    "BURNING_FERVOR_LEGGINGS",
					Type:  "PRESTIGE",
					Price: 4500*1500*constants.APPLICATION_WORTH["essence"] + 50*12000,
					Count: 1,
				},
				{
					Id:    "BURNING_FERVOR_LEGGINGS",
					Type:  "BASE_PRESTIGE_ITEM",
					Price: 30000000,
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly when only base item has price",
			item: &models.NetworthItem{
				ItemId:          "INFERNAL_FERVOR_LEGGINGS",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices: map[string]float64{
				"HEAVY_PEARL":     220000,
				"ESSENCE_CRIMSON": 1500,
				"KUUDRA_TEETH":    12000,
				"FERVOR_LEGGINGS": 1000000,
			},
			shouldApply: true,
			expectedPriceChange: (90100+25500+16345+4500+3055+800+555+150)*1500*constants.APPLICATION_WORTH["essence"] +
				(12+12+12+9)*220000 +
				(80+50+20+10)*12000 +
				1000000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "FIERY_FERVOR_LEGGINGS",
					Type:  "STARS",
					Price: 90100*1500*constants.APPLICATION_WORTH["essence"] + 12*220000,
					Count: 10,
				},
				{
					Id:    "FIERY_FERVOR_LEGGINGS",
					Type:  "PRESTIGE",
					Price: 25500*1500*constants.APPLICATION_WORTH["essence"] + 80*12000,
					Count: 1,
				},
				{
					Id:    "BURNING_FERVOR_LEGGINGS",
					Type:  "STARS",
					Price: 16345*1500*constants.APPLICATION_WORTH["essence"] + 12*220000,
					Count: 10,
				},
				{
					Id:    "BURNING_FERVOR_LEGGINGS",
					Type:  "PRESTIGE",
					Price: 4500*1500*constants.APPLICATION_WORTH["essence"] + 50*12000,
					Count: 1,
				},

				{
					Id:    "HOT_FERVOR_LEGGINGS",
					Type:  "STARS",
					Price: 3055*1500*constants.APPLICATION_WORTH["essence"] + 12*220000,
					Count: 10,
				},
				{
					Id:    "HOT_FERVOR_LEGGINGS",
					Type:  "PRESTIGE",
					Price: 800*1500*constants.APPLICATION_WORTH["essence"] + 20*12000,
					Count: 1,
				},

				{
					Id:    "FERVOR_LEGGINGS",
					Type:  "STARS",
					Price: 555*1500*constants.APPLICATION_WORTH["essence"] + 9*220000,
					Count: 10,
				},
				{
					Id:    "FERVOR_LEGGINGS",
					Type:  "PRESTIGE",
					Price: 150*1500*constants.APPLICATION_WORTH["essence"] + 10*12000,
					Count: 1,
				},
				{
					Id:    "FERVOR_LEGGINGS",
					Type:  "BASE_PRESTIGE_ITEM",
					Price: 1000000,
					Count: 1,
				},
			},
		},

		{
			description: "Applies correctly item when has price",
			item: &models.NetworthItem{
				ItemId:          "INFERNAL_CRIMSON_BOOTS",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           350000000,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{"INFERNAL_CRIMSON_BOOTS": 350000000},
			shouldApply:         true,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "IRON_SWORD",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.PrestigeHandler{}, testCases)
}
