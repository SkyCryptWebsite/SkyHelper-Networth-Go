package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestRecombobulatorHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "IRON_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Recombobulated: 1,
					Enchantments: map[string]int{
						"enchantment": 1,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				ItemLore:    []string{},
			},
			prices:              map[string]float64{"RECOMBOBULATOR_3000": 10000000},
			shouldApply:         true,
			expectedPriceChange: 10000000 * constants.APPLICATION_WORTH["recombobulator"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "RECOMBOBULATOR_3000",
					Type:  "RECOMBOBULATOR_3000",
					Price: 10000000 * constants.APPLICATION_WORTH["recombobulator"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with accessory via category",
			item: &models.NetworthItem{
				ItemId: "HEGEMONY_ARTIFACT",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Recombobulated: 1,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				ItemLore:    []string{},
				SkyblockItem: &models.HypixelItem{
					Category: "ACCESSORY",
				},
			},
			prices:              map[string]float64{"RECOMBOBULATOR_3000": 10000000},
			shouldApply:         true,
			expectedPriceChange: 10000000 * constants.APPLICATION_WORTH["recombobulator"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "RECOMBOBULATOR_3000",
					Type:  "RECOMBOBULATOR_3000",
					Price: 10000000 * constants.APPLICATION_WORTH["recombobulator"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with accessory via accessory",
			item: &models.NetworthItem{
				ItemId: "TEST_ACCESSORY_WITHOUT_SKYBLOCK_ITEM",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Recombobulated: 1,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				ItemLore:    []string{"MYTHIC ACCESSORY"},
			},
			prices:              map[string]float64{"RECOMBOBULATOR_3000": 10000000},
			shouldApply:         true,
			expectedPriceChange: 10000000 * constants.APPLICATION_WORTH["recombobulator"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "RECOMBOBULATOR_3000",
					Type:  "RECOMBOBULATOR_3000",
					Price: 10000000 * constants.APPLICATION_WORTH["recombobulator"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with accessory via hatcessory",
			item: &models.NetworthItem{
				ItemId: "TEST_HATCESSORY_WITHOUT_SKYBLOCK_ITEM",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Recombobulated: 1,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				ItemLore:    []string{"MYTHIC HATCESSORY"},
			},
			prices:              map[string]float64{"RECOMBOBULATOR_3000": 10000000},
			shouldApply:         true,
			expectedPriceChange: 10000000 * constants.APPLICATION_WORTH["recombobulator"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "RECOMBOBULATOR_3000",
					Type:  "RECOMBOBULATOR_3000",
					Price: 10000000 * constants.APPLICATION_WORTH["recombobulator"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly due to specific item exception",
			item: &models.NetworthItem{
				ItemId: "DIVAN_CHESTPLATE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Recombobulated: 1,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				ItemLore:    []string{"MYTHIC HATCESSORY"},
			},
			prices:              map[string]float64{"RECOMBOBULATOR_3000": 10000000},
			shouldApply:         true,
			expectedPriceChange: 10000000 * constants.APPLICATION_WORTH["recombobulator"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "RECOMBOBULATOR_3000",
					Type:  "RECOMBOBULATOR_3000",
					Price: 10000000 * constants.APPLICATION_WORTH["recombobulator"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with bonemerang",
			item: &models.NetworthItem{
				ItemId: "BONE_BOOMERANG",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Recombobulated: 1,
					Enchantments: map[string]int{
						"power": 5,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				ItemLore:    []string{},
			},
			prices:              map[string]float64{"RECOMBOBULATOR_3000": 10000000},
			shouldApply:         true,
			expectedPriceChange: 10000000 * 0.5 * constants.APPLICATION_WORTH["recombobulator"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "RECOMBOBULATOR_3000",
					Type:  "RECOMBOBULATOR_3000",
					Price: 10000000 * 0.5 * constants.APPLICATION_WORTH["recombobulator"],
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
				ItemLore:        []string{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply due to category",
			item: &models.NetworthItem{
				ItemId: "RADIANT_POWER_ORB",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Recombobulated: 1,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				ItemLore:    []string{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply due to dungeon drop",
			item: &models.NetworthItem{
				ItemId: "MACHINE_GUN_BOW",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Recombobulated: 1,
					ItemTier:       1,
					Enchantments: map[string]int{
						"power": 5,
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				ItemLore:    []string{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.RecombobulatorHandler{}, testCases)
}
