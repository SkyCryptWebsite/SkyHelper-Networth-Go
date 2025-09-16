package tests

import (
	"encoding/json"
	"os"
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestEssenceStarsHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "HYPERION",
				SkyblockItem: &models.HypixelItem{
					UpgradeCosts: [][]models.UpgradeCost{
						{
							{
								Type:        "ESSENCE",
								EssenceType: "WITHER",
								Amount:      10,
							},
						},
						{
							{
								Type:        "ESSENCE",
								EssenceType: "WITHER",
								Amount:      20,
							},
						},
						{
							{
								Type:        "ESSENCE",
								EssenceType: "WITHER",
								Amount:      30,
							},
						},
						{
							{
								Type:        "ESSENCE",
								EssenceType: "WITHER",
								Amount:      40,
							},
						},
						{
							{
								Type:        "ESSENCE",
								EssenceType: "WITHER",
								Amount:      50,
							},
						},
					},
				},
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					DungeonItemLevel: "3b",
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"ESSENCE_WITHER": 100},
			shouldApply:         true,
			expectedPriceChange: (10 + 20 + 30) * 100 * constants.APPLICATION_WORTH["essence"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "WITHER_ESSENCE",
					Type:  "STAR",
					Price: 10 * 100 * constants.APPLICATION_WORTH["essence"],
					Count: 10,
					Star:  1,
				},
				{
					Id:    "WITHER_ESSENCE",
					Type:  "STAR",
					Price: 20 * 100 * constants.APPLICATION_WORTH["essence"],
					Count: 20,
					Star:  2,
				},
				{
					Id:    "WITHER_ESSENCE",
					Type:  "STAR",
					Price: 30 * 100 * constants.APPLICATION_WORTH["essence"],
					Count: 30,
					Star:  3,
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

	f, err := os.Create("testData.json")
	if err != nil {
		panic("Failed to create file: " + err.Error())
	}

	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(testCases[0].item.SkyblockItem); err != nil {
		panic("Failed to encode JSON: " + err.Error())
	}

	runHandlerTests(t, &handlers.EssenceStarsHandler{}, testCases)
}
