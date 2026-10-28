package handlers

import (
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type PickonimbusHandler struct{}

func (h PickonimbusHandler) IsCosmetic() bool {
	return false
}

func (h PickonimbusHandler) Applies(item *models.NetworthItem) bool {
	// Only a used one loses value. 0 means the field is missing: Go can't tell it from a fully used one.
	durability := item.ExtraAttributes.PickonimbusDurability
	return item.ItemId == "PICKONIMBUS" && durability > 0 && durability < constants.PICKONIMBUS_DURABILITY
}

func (h PickonimbusHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
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
