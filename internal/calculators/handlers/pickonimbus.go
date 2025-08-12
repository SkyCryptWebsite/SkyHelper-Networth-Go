package handlers

import (
	"github.com/duckysolucky/skyhelper-networth-go/internal/constants"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

type PickonimbusHandler struct{}

func (h PickonimbusHandler) IsCosmetic() bool {
	return false
}

func (h PickonimbusHandler) Applies(item *models.NetworthItem) bool {
	return item.ItemId == "PICKONIMBUS" && item.ExtraAttributes.PickonimbusDurability > 0
}

func (h PickonimbusHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	reduction := float64(item.ExtraAttributes.PickonimbusDurability) / float64(constants.PICKONIMBUS_DURABILITY)

	calculationData := models.CalculationData{
		Id:    "PICKONIMBUS_DURABLITY",
		Type:  "PICKONIMBUS",
		Price: item.BasePrice * float64(reduction-1),
		Count: int(constants.PICKONIMBUS_DURABILITY - item.ExtraAttributes.PickonimbusDurability),
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
