package handlers

import (
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type ArtOfWarHandler struct{}

func (h ArtOfWarHandler) IsCosmetic() bool {
	return false
}

func (h ArtOfWarHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.ArtOfWarCount > 0
}

func (h ArtOfWarHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	artOfWarCount := item.ExtraAttributes.ArtOfWarCount
	calculationData := models.CalculationData{
		Id:    "THE_ART_OF_WAR",
		Type:  "THE_ART_OF_WAR",
		Price: prices["THE_ART_OF_WAR"] * float64(artOfWarCount) * constants.APPLICATION_WORTH["artOfWar"],
		Count: artOfWarCount,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
