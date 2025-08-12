package handlers

import (
	"duckysolucky/skyhelper-networth-go/internal/constants"
	"duckysolucky/skyhelper-networth-go/internal/models"
)

type DivanPowderCoatingHandler struct{}

func (h DivanPowderCoatingHandler) IsCosmetic() bool {
	return false
}

func (h DivanPowderCoatingHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.DivanPowderCoating > 0
}

func (h DivanPowderCoatingHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	calculationData := models.CalculationData{
		Id:    "DIVAN_POWDER_COATING",
		Type:  "DIVAN_POWDER_COATING",
		Price: prices["DIVAN_POWDER_COATING"] * constants.APPLICATION_WORTH["divanPowderCoating"],
		Count: item.ExtraAttributes.DivanPowderCoating,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
