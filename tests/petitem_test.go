package tests

import (
	"testing"

	"github.com/duckysolucky/skyhelper-networth-go/internal/calculators/handlers"
	"github.com/duckysolucky/skyhelper-networth-go/internal/constants"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

func TestPetItemHandler(t *testing.T) {
	testCases := []PetTestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthPet{
				PetData:     models.SkyblockPet{HeldItem: "PET_ITEM_MINING_SKILL_BOOST_UNCOMMON"},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"PET_ITEM_MINING_SKILL_BOOST_UNCOMMON": 200000},
			expectedPriceChange: 200000 * constants.APPLICATION_WORTH["petItem"],
			shouldApply:         true,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "PET_ITEM_MINING_SKILL_BOOST_UNCOMMON",
					Type:  "PET_ITEM",
					Price: 200000 * constants.APPLICATION_WORTH["petItem"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthPet{
				Level:       models.Level{},
				PetData:     models.SkyblockPet{},
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runPetHandlerTests(t, &handlers.PetItemHandler{}, testCases)
}
