package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func itemEnchantmentUpgradeCalculation(id string, price float64) models.CalculationData {
	return models.CalculationData{
		Id:    id,
		Type:  "ENCHANTMENT_UPGRADE",
		Price: price * constants.APPLICATION_WORTH["enchantmentUpgrades"],
		Count: 1,
	}
}

func itemEnchantmentUpgradePriceChange(calculation []models.CalculationData) float64 {
	var price float64
	for _, calculationData := range calculation {
		price += calculationData.Price
	}

	return price
}

func itemEnchantmentUpgradeTestCase(description string, enchantments map[string]int, prices map[string]float64, expectedCalculation []models.CalculationData) TestCase {
	return TestCase{
		description: description,
		item: &models.NetworthItem{
			ItemId: "IRON_SWORD",
			ExtraAttributes: &skycrypttypes.ExtraAttributes{
				Enchantments: enchantments,
			},
			Price:       100,
			Calculation: []models.CalculationData{},
		},
		prices:              prices,
		shouldApply:         true,
		expectedPriceChange: itemEnchantmentUpgradePriceChange(expectedCalculation),
		expectedCalculation: expectedCalculation,
	}
}

func TestItemEnchantmentHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "ROTTEN_LEGGINGS",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
					Id:    "GROWTH_6",
					Type:  "ENCHANTMENT",
					Price: 3000000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				}, {
					Id:    "REJUVENATE_5",
					Type:  "ENCHANTMENT",
					Price: 450000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				}, {
					Id:    "TRUE_PROTECTION_1",
					Type:  "ENCHANTMENT",
					Price: 1000000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				}, {
					Id:    "ULTIMATE_LEGION_5",
					Type:  "ENCHANTMENT",
					Price: 40000000 * constants.APPLICATION_WORTH["enchantments"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with blocked item-specific enchantment",
			item: &models.NetworthItem{
				ItemId: "ADVANCED_GARDENING_HOE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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
			description: "Applies correctly without fateful stinger",
			item: &models.NetworthItem{
				ItemId: "HYPERION",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Enchantments: map[string]int{
						"venomous": 6,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"FATEFUL_STINGER": 1000000},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly with fateful stinger",
			item: &models.NetworthItem{
				ItemId: "HYPERION",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Enchantments: map[string]int{
						"venomous": 7,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"FATEFUL_STINGER": 1000000},
			shouldApply:         true,
			expectedPriceChange: 1000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "FATEFUL_STINGER",
					Type:  "ENCHANTMENT_UPGRADE",
					Price: 1000000 * constants.APPLICATION_WORTH["enchantmentUpgrades"],
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
		{
			description: "Does not apply enchantment",
			item: &models.NetworthItem{
				ItemId: "ENCHANTED_BOOK",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
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

func TestItemEnchantmentEndcapUpgrades(t *testing.T) {
	turboGourd := itemEnchantmentUpgradeCalculation("TURBO_GOURD", 10000000)
	enchantedTurboGourd := itemEnchantmentUpgradeCalculation("ENCHANTED_TURBO_GOURD", 25000000)

	testCases := []TestCase{
		itemEnchantmentUpgradeTestCase(
			"smite 6 does not add severed hand",
			map[string]int{"smite": 6},
			map[string]float64{"SEVERED_HAND": 12000000},
			[]models.CalculationData{},
		),
		itemEnchantmentUpgradeTestCase(
			"smite 7 adds severed hand",
			map[string]int{"smite": 7},
			map[string]float64{"SEVERED_HAND": 12000000},
			[]models.CalculationData{itemEnchantmentUpgradeCalculation("SEVERED_HAND", 12000000)},
		),
		itemEnchantmentUpgradeTestCase(
			"ender slayer 6 does not add endstone idol",
			map[string]int{"ender_slayer": 6},
			map[string]float64{"ENDSTONE_IDOL": 20000000},
			[]models.CalculationData{},
		),
		itemEnchantmentUpgradeTestCase(
			"ender slayer 7 adds endstone idol",
			map[string]int{"ender_slayer": 7},
			map[string]float64{"ENDSTONE_IDOL": 20000000},
			[]models.CalculationData{itemEnchantmentUpgradeCalculation("ENDSTONE_IDOL", 20000000)},
		),
		itemEnchantmentUpgradeTestCase(
			"bane of arthropods 6 does not add ensnared snail",
			map[string]int{"bane_of_arthropods": 6},
			map[string]float64{"ENSNARED_SNAIL": 8000000},
			[]models.CalculationData{},
		),
		itemEnchantmentUpgradeTestCase(
			"bane of arthropods 7 adds ensnared snail",
			map[string]int{"bane_of_arthropods": 7},
			map[string]float64{"ENSNARED_SNAIL": 8000000},
			[]models.CalculationData{itemEnchantmentUpgradeCalculation("ENSNARED_SNAIL", 8000000)},
		),
		itemEnchantmentUpgradeTestCase(
			"turbo wheat 5 does not add turbo gourd",
			map[string]int{"turbo_wheat": 5},
			map[string]float64{"TURBO_GOURD": 10000000, "ENCHANTED_TURBO_GOURD": 25000000},
			[]models.CalculationData{},
		),
		itemEnchantmentUpgradeTestCase(
			"turbo wheat 6 adds turbo gourd",
			map[string]int{"turbo_wheat": 6},
			map[string]float64{"TURBO_GOURD": 10000000, "ENCHANTED_TURBO_GOURD": 25000000},
			[]models.CalculationData{turboGourd},
		),
		itemEnchantmentUpgradeTestCase(
			"turbo wheat 7 adds both turbo upgrades",
			map[string]int{"turbo_wheat": 7},
			map[string]float64{"TURBO_GOURD": 10000000, "ENCHANTED_TURBO_GOURD": 25000000},
			[]models.CalculationData{turboGourd, enchantedTurboGourd},
		),
		itemEnchantmentUpgradeTestCase(
			"multiple turbo crop enchantments at tier 7 only add each turbo upgrade once",
			map[string]int{"turbo_wheat": 7, "turbo_cane": 7},
			map[string]float64{"TURBO_GOURD": 10000000, "ENCHANTED_TURBO_GOURD": 25000000},
			[]models.CalculationData{turboGourd, enchantedTurboGourd},
		),
		itemEnchantmentUpgradeTestCase(
			"thorns 3 does not add prickly creeper",
			map[string]int{"thorns": 3},
			map[string]float64{"PRICKLY_CREEPER": 4000000},
			[]models.CalculationData{},
		),
		itemEnchantmentUpgradeTestCase(
			"thorns 4 adds prickly creeper",
			map[string]int{"thorns": 4},
			map[string]float64{"PRICKLY_CREEPER": 4000000},
			[]models.CalculationData{itemEnchantmentUpgradeCalculation("PRICKLY_CREEPER", 4000000)},
		),
		itemEnchantmentUpgradeTestCase(
			"scuba 5 does not add vibrant coral",
			map[string]int{"scuba": 5},
			map[string]float64{"VIBRANT_CORAL": 6000000},
			[]models.CalculationData{},
		),
		itemEnchantmentUpgradeTestCase(
			"scuba 6 adds vibrant coral",
			map[string]int{"scuba": 6},
			map[string]float64{"VIBRANT_CORAL": 6000000},
			[]models.CalculationData{itemEnchantmentUpgradeCalculation("VIBRANT_CORAL", 6000000)},
		),
	}

	runHandlerTests(t, &handlers.ItemEnchantments{}, testCases)
}
