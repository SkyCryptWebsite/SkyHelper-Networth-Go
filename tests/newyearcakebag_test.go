package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestNewYearCakeBagHandler(t *testing.T) {
	testCases := []TestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthItem{
				ItemId: "NEW_YEAR_CAKE_BAG",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					NewYearCakeBagYears: []int{0, 1, 2, 3, 4, 5},
				},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"NEW_YEAR_CAKE_1": 1000000, "NEW_YEAR_CAKE_2": 2000000, "NEW_YEAR_CAKE_3": 3000000, "NEW_YEAR_CAKE_4": 4000000, "NEW_YEAR_CAKE_5": 5000000},
			shouldApply:         true,
			expectedPriceChange: 1000000 + 2000000 + 3000000 + 4000000 + 5000000,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "NEW_YEAR_CAKES",
					Type:  "NEW_YEAR_CAKES",
					Price: 1000000 + 2000000 + 3000000 + 4000000 + 5000000,
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthItem{
				ItemId: "NEW_YEAR_CAKE_BAG",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{
					NewYearCakeBagYears: []int{},
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
				ItemId:          "NEW_YEAR_CAKE_BAG",
				ExtraAttributes: &skycrypttypes.ExtraAttributes{},
				Price:           100,
				Calculation:     []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runHandlerTests(t, &handlers.NewYearCakeBagHandler{}, testCases)
}
