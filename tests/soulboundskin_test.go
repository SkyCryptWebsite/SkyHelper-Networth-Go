package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestSoulboundSkinHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "DIAMOND_NECRON_HEAD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Skin: "NECRON_DIAMOND_KNIGHT",
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				ItemLore:    []string{"§8§l* §8Co-op Soulbound §8§l*"},
			},
			prices:              map[string]float64{"NECRON_DIAMOND_KNIGHT": 60000000},
			shouldApply:         true,
			expectedPriceChange: 60000000 * constants.APPLICATION_WORTH["soulboundPetSkins"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "NECRON_DIAMOND_KNIGHT",
					Type:  "SOULBOUND_SKIN",
					Price: 60000000 * constants.APPLICATION_WORTH["soulboundPetSkins"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply when not soulbound",
			item: &models.NetworthItem{
				ItemId: "DIAMOND_NECRON_HEAD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Skin: "NECRON_DIAMOND_KNIGHT",
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				ItemLore:    []string{""},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply when already has skin value",
			item: &models.NetworthItem{
				ItemId: "WITHER_GOGGLES_SKINNED_WITHER_GOGGLES_CELESTIAL",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					Skin: "WITHER_GOGGLES_CELESTIAL",
				},
				Price:       100,
				Calculation: []models.CalculationData{},
				ItemLore:    []string{""},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId:          "LEATHER_CHESTPLATE",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.SoulboundSkinHandler{}, testCases)
}
