package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestMasterStarsHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "HYPERION",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					UpgradeLevel: 10,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
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
			},
			prices: map[string]float64{
				"FIRST_MASTER_STAR":  15000000,
				"SECOND_MASTER_STAR": 25000000,
				"THIRD_MASTER_STAR":  50000000,
				"FOURTH_MASTER_STAR": 90000000,
				"FIFTH_MASTER_STAR":  100000000,
			},
			shouldApply: true,
			expectedPriceChange: 15000000*constants.APPLICATION_WORTH["masterStar"] +
				25000000*constants.APPLICATION_WORTH["masterStar"] +
				50000000*constants.APPLICATION_WORTH["masterStar"] +
				90000000*constants.APPLICATION_WORTH["masterStar"] +
				100000000*constants.APPLICATION_WORTH["masterStar"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "FIRST_MASTER_STAR",
					Type:  "MASTER_STAR",
					Price: 15000000 * constants.APPLICATION_WORTH["masterStar"],
					Count: 1,
				},
				{
					Id:    "SECOND_MASTER_STAR",
					Type:  "MASTER_STAR",
					Price: 25000000 * constants.APPLICATION_WORTH["masterStar"],
					Count: 1,
				},
				{
					Id:    "THIRD_MASTER_STAR",
					Type:  "MASTER_STAR",
					Price: 50000000 * constants.APPLICATION_WORTH["masterStar"],
					Count: 1,
				},
				{
					Id:    "FOURTH_MASTER_STAR",
					Type:  "MASTER_STAR",
					Price: 90000000 * constants.APPLICATION_WORTH["masterStar"],
					Count: 1,
				},
				{
					Id:    "FIFTH_MASTER_STAR",
					Type:  "MASTER_STAR",
					Price: 100000000 * constants.APPLICATION_WORTH["masterStar"],
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly with dungeon_item_level",
			item: &models.NetworthItem{
				ItemId: "HYPERION",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					DungeonItemLevel: "6b",
				},
				Price:       100,
				Calculation: []models.CalculationData{},
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
			},
			prices: map[string]float64{
				"FIRST_MASTER_STAR":  15000000,
				"SECOND_MASTER_STAR": 25000000,
				"THIRD_MASTER_STAR":  50000000,
				"FOURTH_MASTER_STAR": 90000000,
				"FIFTH_MASTER_STAR":  100000000,
			},
			shouldApply:         true,
			expectedPriceChange: 15000000 * constants.APPLICATION_WORTH["masterStar"],
			expectedCalculation: []models.CalculationData{
				{
					Id:    "FIRST_MASTER_STAR",
					Type:  "MASTER_STAR",
					Price: 15000000 * constants.APPLICATION_WORTH["masterStar"],
					Count: 1,
				},
			},
		},
	}

	runHandlerTests(t, &handlers.MasterStarsHandler{}, testCases)
}
