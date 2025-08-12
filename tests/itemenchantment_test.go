package tests

import (
	"testing"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestItemEnchantmentHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "ROTTEN_LEGGINGS",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"true_protection": 1,
						"ultimate_legion": 5,
						"rejuvenate":      5,
						"growth":          6,
						"protection":      5,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:      map[string]float64{"ENCHANTMENT_TRUE_PROTECTION_1": 1000000, "ENCHANTMENT_ULTIMATE_LEGION_5": 40000000, "ENCHANTMENT_REJUVENATE_5": 450000, "ENCHANTMENT_GROWTH_6": 3000000},
			shouldApply: true,
			expectedPriceChange: 1000000*constants.APPLICATION_WORTH["enchantments"] +
				40000000*constants.APPLICATION_WORTH["enchantments"] +
				450000*constants.APPLICATION_WORTH["enchantments"] +
				3000000*constants.APPLICATION_WORTH["enchantments"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "TRUE_PROTECTION_1",
					Type:  "ENCHANTMENT",
					Price: 1000000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				}, {
					Id:    "ULTIMATE_LEGION_5",
					Type:  "ENCHANTMENT",
					Price: 40000000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				}, {
					Id:    "REJUVENATE_5",
					Type:  "ENCHANTMENT",
					Price: 450000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				}, {
					Id:    "GROWTH_6",
					Type:  "ENCHANTMENT",
					Price: 3000000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with blocked item-specific enchantment",
			item: &models.NetworthItem{
				ItemId: "ADVANCED_GARDENING_HOE",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"replenish":  1,
						"turbo_cane": 1,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"ENCHANTMENT_REPLENISH_1": 1500000, "ENCHANTMENT_TURBO_CANE_1": 5000},
			shouldApply:         true,
			expectedPriceChange: 5000 * constants.APPLICATION_WORTH["enchantments"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "TURBO_CANE_1",
					Type:  "ENCHANTMENT",
					Price: 5000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with ignored enchantment",
			item: &models.NetworthItem{
				ItemId: "IRON_SWORD",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"scavenger": 5,
						"smite":     6,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"ENCHANTMENT_SCAVENGER_5": 300000, "ENCHANTMENT_SMITE_6": 10},
			shouldApply:         true,
			expectedPriceChange: 10 * constants.APPLICATION_WORTH["enchantments"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "SMITE_6",
					Type:  "ENCHANTMENT",
					Price: 10 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "ARTIFACT_OF_CONTROL",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly with stacking enchantment",
			item: &models.NetworthItem{
				ItemId: "DIVAN_DRILL",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"compact": 10,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"ENCHANTMENT_COMPACT_1": 6000000},
			shouldApply:         true,
			expectedPriceChange: 6000000 * constants.APPLICATION_WORTH["enchantments"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "COMPACT_1",
					Type:  "ENCHANTMENT",
					Price: 6000000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly without silex",
			item: &models.NetworthItem{
				ItemId: "DIAMOND_PICKAXE",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"efficiency": 5,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"SIL_EX": 4500000},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly with silex",
			item: &models.NetworthItem{
				ItemId: "DIAMOND_PICKAXE",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"efficiency": 10,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"SIL_EX": 4500000},
			shouldApply:         true,
			expectedPriceChange: 5 * 4500000 * constants.APPLICATION_WORTH["silex"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "SIL_EX",
					Type:  "SILEX",
					Price: 5 * 4500000 * constants.APPLICATION_WORTH["silex"],
					Count: 5,
				},
			},
		},
		{
			description: "Applies correctly stonk without silex",
			item: &models.NetworthItem{
				ItemId: "STONK_PICKAXE",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"efficiency": 6,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"SIL_EX": 4500000},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly stonk with silex",
			item: &models.NetworthItem{
				ItemId: "STONK_PICKAXE",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"efficiency": 10,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"SIL_EX": 4500000},
			shouldApply:         true,
			expectedPriceChange: 4 * 4500000 * constants.APPLICATION_WORTH["silex"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "SIL_EX",
					Type:  "SILEX",
					Price: 4 * 4500000 * constants.APPLICATION_WORTH["silex"],
					Count: 4,
				},
			},
		},
		{
			description: "Applies correctly promising spade without silex",
			item: &models.NetworthItem{
				ItemId: "PROMISING_SPADE",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"efficiency": 10,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"SIL_EX": 4500000},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},

		{
			description: "Applies correctly without golden bounty",
			item: &models.NetworthItem{
				ItemId: "IRON_SWORD",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"scavenger": 5,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"GOLDEN_BOUNTY": 30000000},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly with golden bounty",
			item: &models.NetworthItem{
				ItemId: "IRON_SWORD",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"scavenger": 6,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"GOLDEN_BOUNTY": 30000000},
			shouldApply:         true,
			expectedPriceChange: 30000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "GOLDEN_BOUNTY",
					Type:  "ENCHANTMENT_UPGRADE",
					Price: 30000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly without pesthunting guide",
			item: &models.NetworthItem{
				ItemId: "FERMENTO_LEGGINGS",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"pesterminator": 5,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"PESTHUNTING_GUIDE": 10000000},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly with pesthunting guide",
			item: &models.NetworthItem{
				ItemId: "FERMENTO_LEGGINGS",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"pesterminator": 6,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"PESTHUNTING_GUIDE": 10000000},
			shouldApply:         true,
			expectedPriceChange: 10000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "PESTHUNTING_GUIDE",
					Type:  "ENCHANTMENT_UPGRADE",
					Price: 10000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly without gold bottle cap",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"luck_of_the_sea": 6,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"GOLD_BOTTLE_CAP": 28000000},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly with gold bottle cap",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"luck_of_the_sea": 7,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"GOLD_BOTTLE_CAP": 28000000},
			shouldApply:         true,
			expectedPriceChange: 28000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "GOLD_BOTTLE_CAP",
					Type:  "ENCHANTMENT_UPGRADE",
					Price: 28000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly without troubled bubble",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"piscary": 6,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"TROUBLED_BUBBLE": 150000000},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly with troubled bubble",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"piscary": 7,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"TROUBLED_BUBBLE": 150000000},
			shouldApply:         true,
			expectedPriceChange: 150000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "TROUBLED_BUBBLE",
					Type:  "ENCHANTMENT_UPGRADE",
					Price: 150000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly without severed pincer",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"frail": 6,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"SEVERED_PINCER": 4000000},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly with severed pincer",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"frail": 7,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"SEVERED_PINCER": 4000000},
			shouldApply:         true,
			expectedPriceChange: 4000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "SEVERED_PINCER",
					Type:  "ENCHANTMENT_UPGRADE",
					Price: 4000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly without octopus tendril",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"spiked_hook": 6,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"OCTOPUS_TENDRIL": 4500000},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly with octopus tendril",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"spiked_hook": 7,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"OCTOPUS_TENDRIL": 4500000},
			shouldApply:         true,
			expectedPriceChange: 4500000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "OCTOPUS_TENDRIL",
					Type:  "ENCHANTMENT_UPGRADE",
					Price: 4500000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly without chain of the end times",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"charm": 5,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"CHAIN_END_TIMES": 2000000},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly with chain of the end times",
			item: &models.NetworthItem{
				ItemId: "ROD_OF_THE_SEA",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"charm": 6,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"CHAIN_END_TIMES": 2000000},
			shouldApply:         true,
			expectedPriceChange: 2000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "CHAIN_END_TIMES",
					Type:  "ENCHANTMENT_UPGRADE",
					Price: 2000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
					Count: 1,
				},
			},
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
		{
			description: "Does not apply enchantment",
			item: &models.NetworthItem{
				ItemId: "ENCHANTED_BOOK",
				ExtraAttributes: &models.ExtraAttributes{
					Enchantments: map[string]int{
						"fire_protection": 6,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"ENCHANTMENT_FIRE_PROTECTION_6": 1500},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.ItemEnchantments{}, testCases)

}
