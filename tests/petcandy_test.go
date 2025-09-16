package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestPetCandyHandler(t *testing.T) {
	testCases := []PetTestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthPet{
				Level:       models.Level{ExperienceToMax: 25000000, Level: 100},
				PetData:     skycrypttypes.Pet{CandyUsed: 10},
				BasePrice:   100000,
				Price:       100000,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			expectedPriceChange: 100000*constants.APPLICATION_WORTH["petCandy"] - 100000,
			shouldApply:         true,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "CANDY",
					Type:  "PET_CANDY",
					Price: 100000*constants.APPLICATION_WORTH["petCandy"] - 100000,
					Count: 10,
				},
			},
		},
		{
			description: "Applies correctly  with cap and level 100",
			item: &models.NetworthPet{
				Level:       models.Level{ExperienceToMax: 25000000, Level: 100},
				PetData:     skycrypttypes.Pet{CandyUsed: 10},
				BasePrice:   100000000,
				Price:       100000000,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			expectedPriceChange: -5000000,
			shouldApply:         true,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "CANDY",
					Type:  "PET_CANDY",
					Price: -5000000,
					Count: 10,
				},
			},
		},
		{
			description: "Applies correctly with cap and not level 100",
			item: &models.NetworthPet{
				Level:       models.Level{ExperienceToMax: 25000000, Level: 90},
				PetData:     skycrypttypes.Pet{CandyUsed: 10},
				BasePrice:   100000000,
				Price:       100000000,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			expectedPriceChange: -2500000,
			shouldApply:         true,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "CANDY",
					Type:  "PET_CANDY",
					Price: -2500000,
					Count: 10,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthPet{
				Level:       models.Level{ExperienceToMax: 25000000},
				PetData:     skycrypttypes.Pet{},
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply",
			item: &models.NetworthPet{
				Level: models.Level{ExperienceToMax: 25000000},
				PetData: skycrypttypes.Pet{
					Experience: 35000000,
					CandyUsed:  10,
				},
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runPetHandlerTests(t, &handlers.PetCandyHandler{}, testCases)
}
