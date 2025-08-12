package handlers

import (
	"github.com/duckysolucky/skyhelper-networth-go/internal/constants"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

type ArtOfPeaceHandler struct{}

func (h ArtOfPeaceHandler) IsCosmetic() bool {
	return false
}

func (h ArtOfPeaceHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.ArtOfPeaceApplied > 0
}

func (h ArtOfPeaceHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	artOfPieceAmount := item.ExtraAttributes.ArtOfPeaceApplied
	calculationData := models.CalculationData{
		Id:    "THE_ART_OF_PEACE",
		Type:  "THE_ART_OF_PEACE",
		Price: prices["THE_ART_OF_PEACE"] * float64(artOfPieceAmount) * constants.APPLICATION_WORTH["artOfPeace"],
		Count: artOfPieceAmount,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
