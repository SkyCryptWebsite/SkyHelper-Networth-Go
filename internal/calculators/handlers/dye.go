package handlers

import (
	"strings"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type DyeHandler struct{}

func (h DyeHandler) IsCosmetic() bool {
	return true
}

func (h DyeHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.DyeItem != ""
}

func (h DyeHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	dyeItem := strings.ToUpper(item.ExtraAttributes.DyeItem)
	calculationData := models.CalculationData{
		Id:    dyeItem,
		Type:  "DYE",
		Price: prices[dyeItem] * constants.APPLICATION_WORTH["dye"],
		Count: 1,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
