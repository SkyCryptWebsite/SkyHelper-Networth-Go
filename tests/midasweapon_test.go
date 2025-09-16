package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func TestMidasWeaponHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly less than max price paid",
			item: &models.NetworthItem{
				ItemId: "MIDAS_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					WinningBid: 10000000,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly less than max price paid with additonal coins",
			item: &models.NetworthItem{
				ItemId: "MIDAS_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					WinningBid:      10000000,
					AdditionalCoins: 25000000,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         true,
			expectedPriceChange: 0,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Applies correctly max price paid",
			item: &models.NetworthItem{
				ItemId: "MIDAS_SWORD",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					WinningBid: 50000000,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:               map[string]float64{"MIDAS_SWORD_50M": 300000000},
			shouldApply:          true,
			expectedNewBasePrice: 300000000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "MIDAS_SWORD",
					Type:  "MIDAS_SWORD_50M",
					Price: 300000000,
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly max price paid + additional coins",
			item: &models.NetworthItem{
				ItemId: "MIDAS_STAFF",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					WinningBid:      50000000,
					AdditionalCoins: 50000000,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:               map[string]float64{"MIDAS_STAFF_100M": 400000000},
			shouldApply:          true,
			expectedNewBasePrice: 400000000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "MIDAS_STAFF",
					Type:  "MIDAS_STAFF_100M",
					Price: 400000000,
					Count: 1,
				},
			},
		},
		{
			description: "Applies correctly max price paid",
			item: &models.NetworthItem{
				ItemId: "STARRED_MIDAS_STAFF",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					WinningBid:      50000000,
					AdditionalCoins: 1000000000000,
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:               map[string]float64{"STARRED_MIDAS_STAFF_500M": 580000000},
			shouldApply:          true,
			expectedNewBasePrice: 580000000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "STARRED_MIDAS_STAFF",
					Type:  "STARRED_MIDAS_STAFF_500M",
					Price: 580000000,
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
	}

	runHandlerTests(t, &handlers.MidasWeaponHandler{}, testCases)
}
