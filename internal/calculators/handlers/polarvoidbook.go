package handlers

import (
	"github.com/duckysolucky/skyhelper-networth-go/internal/constants"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

type PolarvoidBookHandler struct{}

func (h PolarvoidBookHandler) IsCosmetic() bool {
	return false
}

func (h PolarvoidBookHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.Polarvoid > 0
}

func (h PolarvoidBookHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
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
