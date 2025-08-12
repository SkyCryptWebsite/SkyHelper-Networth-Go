package handlers

import (
	"github.com/duckysolucky/skyhelper-networth-go/internal/constants"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

type FarmingForDummiesHandler struct{}

func (h FarmingForDummiesHandler) IsCosmetic() bool {
	return false
}

func (h FarmingForDummiesHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.FarmingForDummies > 0
}

func (h FarmingForDummiesHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	farmingForDummiesCount := item.ExtraAttributes.FarmingForDummies
	calculationData := models.CalculationData{
		Id:    "FARMING_FOR_DUMMIES",
		Type:  "FARMING_FOR_DUMMIES",
		Price: prices["FARMING_FOR_DUMMIES"] * float64(farmingForDummiesCount) * constants.APPLICATION_WORTH["farmingForDummies"],
		Count: farmingForDummiesCount,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
