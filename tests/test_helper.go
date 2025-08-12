package tests

import (
	"reflect"
	"testing"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

type TestCase struct {
	description                    string
	item                           *models.NetworthItem
	prices                         map[string]float64
	shouldApply                    bool
	expectedNewBasePrice           float64
	expectedPriceChange            float64
	expectedSoulboundPortionChange float64
	expectedCalculation            []models.CalculationData
}

type PetTestCase struct {
	description          string
	item                 *models.NetworthPet
	prices               map[string]float64
	shouldApply          bool
	expectedNewBasePrice float64
	expectedPriceChange  float64
	expectedCalculation  []models.CalculationData
}

func runHandlerTests(t *testing.T, handler models.Handler, testCases []TestCase) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Check if handler applies
			applies := handler.Applies(tc.item)
			if applies != tc.shouldApply {
				t.Errorf("Expected applies() to return %v, got %v", tc.shouldApply, applies)
			}

			// Validate calculation
			priceBefore := tc.item.Price
			soulboundPortionBefore := tc.item.SoulboundPortion

			if tc.shouldApply {
				tc.item.Prices = tc.prices

				handler.Calculate(tc.item, tc.prices)
			}

			if !reflect.DeepEqual(tc.item.Calculation, tc.expectedCalculation) {
				t.Errorf("Calculation does not match expected.\nExpected: %+v\nGot: %+v", tc.expectedCalculation, tc.item.Calculation)
			}

			if tc.expectedNewBasePrice != 0 { // NOTE: Potential issues, probably should use *float, and nil check
				if tc.item.BasePrice != tc.expectedNewBasePrice {
					t.Errorf("Expected base price to be set to %f but got %f", tc.expectedNewBasePrice, tc.item.BasePrice)
				}
			}

			if tc.expectedSoulboundPortionChange != 0 { // NOTE: Potential issue, probably should be *float64, so we can nil check
				actualChange := tc.item.SoulboundPortion - soulboundPortionBefore
				if actualChange != tc.expectedSoulboundPortionChange {
					t.Errorf("Expected soulbound portion to increase by %f but got %f", tc.expectedSoulboundPortionChange, actualChange)
				}
			}

			actualPriceChange := tc.item.Price - priceBefore
			if actualPriceChange != tc.expectedPriceChange {
				t.Errorf("Expected price to increase by %f but got %f", tc.expectedPriceChange, actualPriceChange)
			}
		})
	}
}

func runPetHandlerTests(t *testing.T, handler models.PetHandler, testCases []PetTestCase) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			// Check if handler applies
			applies := handler.Applies(tc.item)
			if applies != tc.shouldApply {
				t.Errorf("Expected applies() to return %v, got %v", tc.shouldApply, applies)
			}

			// Validate calculation
			priceBefore := tc.item.Price

			if tc.shouldApply {
				tc.item.Prices = tc.prices

				handler.Calculate(tc.item, tc.prices)
			}

			if !reflect.DeepEqual(tc.item.Calculation, tc.expectedCalculation) {
				t.Errorf("Calculation does not match expected.\nExpected: %+v\nGot: %+v", tc.expectedCalculation, tc.item.Calculation)
			}

			if tc.expectedNewBasePrice != 0 { // NOTE: Potential issues, probably should use *float, and nil check
				if tc.item.BasePrice != tc.expectedNewBasePrice {
					t.Errorf("Expected base price to be set to %f but got %f", tc.expectedNewBasePrice, tc.item.BasePrice)
				}
			}

			actualPriceChange := tc.item.Price - priceBefore
			if actualPriceChange != tc.expectedPriceChange {
				t.Errorf("Expected price to increase by %f but got %f", tc.expectedPriceChange, actualPriceChange)
			}
		})
	}
}
