package handlers

import (
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type PolarvoidBookHandler struct{}

func (h PolarvoidBookHandler) IsCosmetic() bool {
	return false
}

func (h PolarvoidBookHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.Polarvoid > 0
}

func (h PolarvoidBookHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	polarvoidBookCount := item.ExtraAttributes.Polarvoid
	calculationData := models.CalculationData{
		Id:    "POLARVOID_BOOK",
		Type:  "POLARVOID_BOOK",
		Price: prices["POLARVOID_BOOK"] * float64(polarvoidBookCount) * constants.APPLICATION_WORTH["polarvoidBook"],
		Count: polarvoidBookCount,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
