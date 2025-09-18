package handlers

import (
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type FarmingForDummiesHandler struct{}

func (h FarmingForDummiesHandler) IsCosmetic() bool {
	return false
}

func (h FarmingForDummiesHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.FarmingForDummies > 0
}

func (h FarmingForDummiesHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
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
