package tests

import (
	"testing"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

func TestSoulboundPetSkinHandler(t *testing.T) {
	testCases := []PetTestCase{
		{
			description: "Applies correctly",
			item: &models.NetworthPet{
				PetData:     models.SkyblockPet{Type: "GRANDMA_WOLF", Rarity: "LEGENDARY", Experience: 0, Skin: "GRANDMA_WOLF_REAL"},
				Price:       100,
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{"PET_SKIN_GRANDMA_WOLF_REAL": 65000000},
			expectedPriceChange: 65000000 * constants.APPLICATION_WORTH["soulboundPetSkins"],
			shouldApply:         true,
			expectedCalculation: []models.CalculationData{
				{
					Id:    "GRANDMA_WOLF_REAL",
					Type:  "SOULBOUND_PET_SKIN",
					Price: 65000000 * constants.APPLICATION_WORTH["soulboundPetSkins"],
					Count: 1,
				},
			},
		},
		{
			description: "Does not apply",
			item: &models.NetworthPet{
				Level:       models.Level{},
				PetData:     models.SkyblockPet{Type: "BLACK_CAT", Rarity: "MYTHIC", Skin: "BLACK_CAT_PURRANORMAL"},
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
		{
			description: "Does not apply",
			item: &models.NetworthPet{
				Level:       models.Level{},
				PetData:     models.SkyblockPet{Type: "BLACK_CAT", Rarity: "MYTHIC"},
				Calculation: []models.CalculationData{},
			},
			prices:              map[string]float64{},
			shouldApply:         false,
			expectedCalculation: []models.CalculationData{},
		},
	}

	runPetHandlerTests(t, &handlers.SoulboundPetSkinHandler{}, testCases)
}
