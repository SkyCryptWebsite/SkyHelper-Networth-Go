package tests

import (
	"testing"

	"github.com/duckysolucky/skyhelper-networth-go/internal/calculators/handlers"
	"github.com/duckysolucky/skyhelper-networth-go/internal/constants"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

func TestGemstonesHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "HYPERION",
				ExtraAttributes: &models.ExtraAttributes{
					Gems: map[string]any{
						"COMBAT_0":       map[string]interface{}{"quality": "PERFECT"},
						"unlocked_slots": []string{"SAPPHIRE_0", "COMBAT_0"},
						"COMBAT_0_gem":   "SAPPHIRE",
						"SAPPHIRE_0":     map[string]interface{}{"quality": "PERFECT"},
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				SkyblockItem: &models.HypixelItem{
					GemstoneSlots: []models.GemstoneSlot{
						{
							SlotType: "SAPPHIRE",
							Costs: []models.GemstoneCost{
								{Type: "COINS", Coins: 250000},
								{Type: "ITEM", ItemId: "FLAWLESS_SAPPHIRE_GEM", Amount: 4},
							},
						},
						{
							SlotType: "COMBAT",
							Costs: []models.GemstoneCost{
								{Type: "COINS", Coins: 250000},
								{Type: "ITEM", ItemId: "FLAWLESS_JASPER_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_SAPPHIRE_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_RUBY_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_AMETHYST_GEM", Amount: 1},
							},
						},
					},
				},
			},
			prices:              map[string]float64{"PERFECT_SAPPHIRE_GEM": 16000000},
			shouldApply:         true,
			expectedPriceChange: 2 * 16000000 * constants.APPLICATION_WORTH["gemstone"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "PERFECT_SAPPHIRE_GEM",
					Type:  "GEMSTONE",
					Price: 16000000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				},
				{
					Id:    "PERFECT_SAPPHIRE_GEM",
					Type:  "GEMSTONE",
					Price: 16000000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly v2",
			item: &models.NetworthItem{
				ItemId: "HYPERION",
				ExtraAttributes: &models.ExtraAttributes{
					Gems: map[string]any{
						"COMBAT_0":        "FINE",
						"COMBAT_0_gem":    "SAPPHIRE",
						"UNIVERSAL_0":     "FLAWLESS",
						"UNIVERSAL_0_gem": "SAPPHIRE",
						"SAPPHIRE_0":      "FINE",
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				SkyblockItem: &models.HypixelItem{
					GemstoneSlots: []models.GemstoneSlot{
						{
							SlotType: "SAPPHIRE",
							Costs: []models.GemstoneCost{
								{Type: "COINS", Coins: 250000},
								{Type: "ITEM", ItemId: "FLAWLESS_SAPPHIRE_GEM", Amount: 4},
							},
						},
						{
							SlotType: "COMBAT",
							Costs: []models.GemstoneCost{
								{Type: "COINS", Coins: 250000},
								{Type: "ITEM", ItemId: "FLAWLESS_JASPER_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_SAPPHIRE_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_RUBY_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_AMETHYST_GEM", Amount: 1},
							},
						},
					},
				},
			},
			prices:              map[string]float64{"FINE_SAPPHIRE_GEM": 30000},
			shouldApply:         true,
			expectedPriceChange: 2 * 30000 * constants.APPLICATION_WORTH["gemstone"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "FINE_SAPPHIRE_GEM",
					Type:  "GEMSTONE",
					Price: 30000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				},
				{
					Id:    "FINE_SAPPHIRE_GEM",
					Type:  "GEMSTONE",
					Price: 30000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly simple",
			item: &models.NetworthItem{
				ItemId: "ADAPTIVE_BOOTS",
				ExtraAttributes: &models.ExtraAttributes{
					Gems: map[string]any{
						"COMBAT_0":     "FINE",
						"COMBAT_0_gem": "JASPER",
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				SkyblockItem: &models.HypixelItem{
					GemstoneSlots: []models.GemstoneSlot{
						{
							SlotType: "COMBAT",
						},
					},
				},
			},
			prices:              map[string]float64{"FINE_JASPER_GEM": 90000},
			shouldApply:         true,
			expectedPriceChange: 90000 * constants.APPLICATION_WORTH["gemstone"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "FINE_JASPER_GEM",
					Type:  "GEMSTONE",
					Price: 90000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with divan",
			item: &models.NetworthItem{
				ItemId: "DIVAN_CHESTPLATE",
				ExtraAttributes: &models.ExtraAttributes{
					Gems: map[string]any{
						"JADE_1":         map[string]interface{}{"quality": "PERFECT"},
						"JADE_0":         map[string]interface{}{"quality": "PERFECT"},
						"unlocked_slots": []string{"TOPAZ_0", "JADE_1", "JADE_0", "AMBER_0", "AMBER_1"},
						"AMBER_0":        map[string]interface{}{"quality": "PERFECT"},
						"AMBER_1":        map[string]interface{}{"quality": "PERFECT"},
						"TOPAZ_0":        map[string]interface{}{"quality": "PERFECT"},
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				SkyblockItem: &models.HypixelItem{
					GemstoneSlots: []models.GemstoneSlot{
						{
							SlotType: "AMBER",
							Costs: []models.GemstoneCost{
								{Type: "ITEM", ItemId: "GEMSTONE_CHAMBER", Amount: 1},
							},
						}, {
							SlotType: "JADE",
							Costs: []models.GemstoneCost{
								{Type: "ITEM", ItemId: "GEMSTONE_CHAMBER", Amount: 1},
							},
						}, {
							SlotType: "AMBER",
							Costs: []models.GemstoneCost{
								{Type: "ITEM", ItemId: "GEMSTONE_CHAMBER", Amount: 1},
							},
						}, {
							SlotType: "JADE",
							Costs: []models.GemstoneCost{
								{Type: "ITEM", ItemId: "GEMSTONE_CHAMBER", Amount: 1},
							},
						}, {
							SlotType: "TOPAZ",
							Costs: []models.GemstoneCost{
								{Type: "ITEM", ItemId: "GEMSTONE_CHAMBER", Amount: 1},
							},
						},
					},
				},
			},
			prices:      map[string]float64{"GEMSTONE_CHAMBER": 7000000, "PERFECT_AMBER_GEM": 15000000, "PERFECT_JADE_GEM": 16000000, "PERFECT_TOPAZ_GEM": 17500000},
			shouldApply: true,
			expectedPriceChange: 5*7000000*constants.APPLICATION_WORTH["gemstoneChambers"] +
				2*16000000*constants.APPLICATION_WORTH["gemstone"] +
				2*15000000*constants.APPLICATION_WORTH["gemstone"] +
				17500000*constants.APPLICATION_WORTH["gemstone"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "AMBER",
					Type:  "GEMSTONE_SLOT",
					Price: 7000000 * constants.APPLICATION_WORTH["gemstoneChambers"],
					Count: 1,
				}, {
					Id:    "JADE",
					Type:  "GEMSTONE_SLOT",
					Price: 7000000 * constants.APPLICATION_WORTH["gemstoneChambers"],
					Count: 1,
				}, {
					Id:    "AMBER",
					Type:  "GEMSTONE_SLOT",
					Price: 7000000 * constants.APPLICATION_WORTH["gemstoneChambers"],
					Count: 1,
				}, {
					Id:    "JADE",
					Type:  "GEMSTONE_SLOT",
					Price: 7000000 * constants.APPLICATION_WORTH["gemstoneChambers"],
					Count: 1,
				}, {
					Id:    "TOPAZ",
					Type:  "GEMSTONE_SLOT",
					Price: 7000000 * constants.APPLICATION_WORTH["gemstoneChambers"],
					Count: 1,
				}, {
					Id:    "PERFECT_AMBER_GEM",
					Type:  "GEMSTONE",
					Price: 15000000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				}, {
					Id:    "PERFECT_JADE_GEM",
					Type:  "GEMSTONE",
					Price: 16000000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				}, {
					Id:    "PERFECT_AMBER_GEM",
					Type:  "GEMSTONE",
					Price: 15000000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				}, {
					Id:    "PERFECT_JADE_GEM",
					Type:  "GEMSTONE",
					Price: 16000000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				}, {
					Id:    "PERFECT_TOPAZ_GEM",
					Type:  "GEMSTONE",
					Price: 17500000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with divan unlocked and no gems",
			item: &models.NetworthItem{
				ItemId: "DIVAN_CHESTPLATE",
				ExtraAttributes: &models.ExtraAttributes{
					Gems: map[string]any{
						"unlocked_slots": []string{"TOPAZ_0", "JADE_1", "JADE_0", "AMBER_0", "AMBER_1"},
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				SkyblockItem: &models.HypixelItem{
					GemstoneSlots: []models.GemstoneSlot{
						{
							SlotType: "AMBER",
							Costs: []models.GemstoneCost{
								{Type: "ITEM", ItemId: "GEMSTONE_CHAMBER", Amount: 1},
							},
						}, {
							SlotType: "JADE",
							Costs: []models.GemstoneCost{
								{Type: "ITEM", ItemId: "GEMSTONE_CHAMBER", Amount: 1},
							},
						}, {
							SlotType: "AMBER",
							Costs: []models.GemstoneCost{
								{Type: "ITEM", ItemId: "GEMSTONE_CHAMBER", Amount: 1},
							},
						}, {
							SlotType: "JADE",
							Costs: []models.GemstoneCost{
								{Type: "ITEM", ItemId: "GEMSTONE_CHAMBER", Amount: 1},
							},
						}, {
							SlotType: "TOPAZ",
							Costs: []models.GemstoneCost{
								{Type: "ITEM", ItemId: "GEMSTONE_CHAMBER", Amount: 1},
							},
						},
					},
				},
			},
			prices:              map[string]float64{"GEMSTONE_CHAMBER": 7000000, "PERFECT_AMBER_GEM": 15000000, "PERFECT_JADE_GEM": 16000000, "PERFECT_TOPAZ_GEM": 17500000},
			shouldApply:         true,
			expectedPriceChange: 5 * 7000000 * constants.APPLICATION_WORTH["gemstoneChambers"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "AMBER",
					Type:  "GEMSTONE_SLOT",
					Price: 7000000 * constants.APPLICATION_WORTH["gemstoneChambers"],
					Count: 1,
				}, {
					Id:    "JADE",
					Type:  "GEMSTONE_SLOT",
					Price: 7000000 * constants.APPLICATION_WORTH["gemstoneChambers"],
					Count: 1,
				}, {
					Id:    "AMBER",
					Type:  "GEMSTONE_SLOT",
					Price: 7000000 * constants.APPLICATION_WORTH["gemstoneChambers"],
					Count: 1,
				}, {
					Id:    "JADE",
					Type:  "GEMSTONE_SLOT",
					Price: 7000000 * constants.APPLICATION_WORTH["gemstoneChambers"],
					Count: 1,
				}, {
					Id:    "TOPAZ",
					Type:  "GEMSTONE_SLOT",
					Price: 7000000 * constants.APPLICATION_WORTH["gemstoneChambers"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with kuudra",
			item: &models.NetworthItem{
				ItemId: "INFERNAL_AURORA_CHESTPLATE",
				ExtraAttributes: &models.ExtraAttributes{
					Gems: map[string]any{
						"COMBAT_0":       "PERFECT",
						"unlocked_slots": []string{"COMBAT_0", "COMBAT_1"},
						"COMBAT_1_gem":   "SAPPHIRE",
						"COMBAT_0_gem":   "SAPPHIRE",
						"COMBAT_1":       "PERFECT",
					},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				SkyblockItem: &models.HypixelItem{
					GemstoneSlots: []models.GemstoneSlot{
						{
							SlotType: "COMBAT",
							Costs: []models.GemstoneCost{
								{Type: "COINS", Coins: 250000},
								{Type: "ITEM", ItemId: "FLAWLESS_JASPER_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_SAPPHIRE_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_RUBY_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_AMETHYST_GEM", Amount: 1},
							},
						}, {
							SlotType: "COMBAT",
							Costs: []models.GemstoneCost{
								{Type: "COINS", Coins: 250000},
								{Type: "ITEM", ItemId: "FLAWLESS_JASPER_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_SAPPHIRE_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_RUBY_GEM", Amount: 1},
								{Type: "ITEM", ItemId: "FLAWLESS_AMETHYST_GEM", Amount: 1},
							},
						},
					},
				},
			},
			prices:              map[string]float64{"FLAWLESS_JASPER_GEM": 7500000, "FLAWLESS_SAPPHIRE_GEM": 2500000, "FLAWLESS_RUBY_GEM": 2000000, "FLAWLESS_AMETHYST_GEM": 2250000, "PERFECT_SAPPHIRE_GEM": 16000000},
			shouldApply:         true,
			expectedPriceChange: 2*(250000+7500000+2500000+2000000+2250000)*constants.APPLICATION_WORTH["gemstoneSlots"] + 2*16000000*constants.APPLICATION_WORTH["gemstone"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "COMBAT",
					Type:  "GEMSTONE_SLOT",
					Price: (250000 + 7500000 + 2500000 + 2000000 + 2250000) * constants.APPLICATION_WORTH["gemstoneSlots"],
					Count: 1,
				}, {
					Id:    "COMBAT",
					Type:  "GEMSTONE_SLOT",
					Price: (250000 + 7500000 + 2500000 + 2000000 + 2250000) * constants.APPLICATION_WORTH["gemstoneSlots"],
					Count: 1,
				}, {
					Id:    "PERFECT_SAPPHIRE_GEM",
					Type:  "GEMSTONE",
					Price: 16000000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				}, {
					Id:    "PERFECT_SAPPHIRE_GEM",
					Type:  "GEMSTONE",
					Price: 16000000 * constants.APPLICATION_WORTH["gemstone"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "HYPERION",
				ExtraAttributes: &models.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.GemstonesHandler{}, testCases)
}
