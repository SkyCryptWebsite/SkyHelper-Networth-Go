package handlers

import (
	"slices"

	"github.com/duckysolucky/skyhelper-networth-go/internal/constants"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

type PetCandyHandler struct{}

func (h PetCandyHandler) IsCosmetic() bool {
	return false
}

func (h PetCandyHandler) Applies(item *models.NetworthPet) bool {
	matPetCandyXp := item.PetData.CandyUsed * 1_000_000
	xpLessPetCandy := item.PetData.Experience - float64(matPetCandyXp)
	return item.PetData.CandyUsed > 0 && !slices.Contains(constants.BLOCKED_CANDY_REDUCE_PETS, item.PetData.Type) && xpLessPetCandy < float64(item.Level.ExperienceToMax)
}

func (h PetCandyHandler) Calculate(item *models.NetworthPet, prices models.Prices) {
	reduceValue := item.BasePrice * (1 - constants.APPLICATION_WORTH["petCandy"])
	maxReduction := 2_500_000.0
	if item.Level.Level == 100 {
		maxReduction = 5_000_000.0
	}

	reduceValue = min(reduceValue, maxReduction)
	calculationData := models.CalculationData{
		Id:    "CANDY",
		Type:  "PET_CANDY",
		Price: -reduceValue,
		Count: item.PetData.CandyUsed,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
