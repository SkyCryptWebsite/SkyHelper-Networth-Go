package handlers

import (
	"duckysolucky/skyhelper-networth-go/internal/constants"
	"duckysolucky/skyhelper-networth-go/internal/models"
)

type JalapenoBookHandler struct{}

func (h JalapenoBookHandler) IsCosmetic() bool {
	return false
}

func (h JalapenoBookHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.JalapenoCount > 0
}

func (h JalapenoBookHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	jalapenoCount := item.ExtraAttributes.JalapenoCount
	calculationData := models.CalculationData{
		Id:    "JALAPENO_BOOK",
		Type:  "JALAPENO_BOOK",
		Price: prices["JALAPENO_BOOK"] * float64(jalapenoCount) * constants.APPLICATION_WORTH["jalapenoBook"],
		Count: jalapenoCount,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
