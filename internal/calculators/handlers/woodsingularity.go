package handlers

import (
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

type WoodSingularityHandler struct{}

func (h WoodSingularityHandler) IsCosmetic() bool {
	return false
}

func (h WoodSingularityHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.WoodSingularityCount > 0
}

func (h WoodSingularityHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	woodSingularityCount := item.ExtraAttributes.WoodSingularityCount
	calculationData := models.CalculationData{
		Id:    "WOOD_SINGULARITY",
		Type:  "WOOD_SINGULARITY",
		Price: prices["WOOD_SINGULARITY"] * float64(woodSingularityCount) * constants.APPLICATION_WORTH["woodSingularity"],
		Count: woodSingularityCount,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
