package tests

import (
	"testing"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestEnchantedBookandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly with single enchantment",
			item: &models.NetworthItem{
				ItemId: "ENCHANTED_BOOK",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"ultimate_legion": 7,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:               map[string]float64{"ENCHANTMENT_ULTIMATE_LEGION_7": 50000000},
			shouldApply:          true,
			expectedNewBasePrice: 50000000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "ULTIMATE_LEGION_7",
					Type:  "ENCHANT",
					Price: 50000000,
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with mutliple enchantment",
			item: &models.NetworthItem{
				ItemId: "ENCHANTED_BOOK",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"ultimate_legion": 7,
						"smite":           7,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:               map[string]float64{"ENCHANTMENT_ULTIMATE_LEGION_7": 50000000, "ENCHANTMENT_SMITE_7": 4000000},
			shouldApply:          true,
			expectedNewBasePrice: 50000000*constants.APPLICATION_WORTH["enchantments"] + 4000000*constants.APPLICATION_WORTH["enchantments"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "ULTIMATE_LEGION_7",
					Type:  "ENCHANT",
					Price: 50000000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				},
				{
					Id:    "SMITE_7",
					Type:  "ENCHANT",
					Price: 4000000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with no price",
			item: &models.NetworthItem{
				ItemId: "ENCHANTED_BOOK",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"smite": 5,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         true,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply on items",
			item: &models.NetworthItem{
				ItemId: "IRON_SWORD",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"sharpness": 5,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "ENCHANTED_BOOK",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.EnchantedBookHandler{}, testCases)
}
