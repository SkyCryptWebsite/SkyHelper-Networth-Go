package handlers

import (
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

type GemstonePowerScrollHandler struct{}

func (h GemstonePowerScrollHandler) IsCosmetic() bool {
	return false
}

func (h GemstonePowerScrollHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.PowerAbilityScroll != ""
}

func (h GemstonePowerScrollHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	powerAbilityScroll := item.ExtraAttributes.PowerAbilityScroll
	calculationData := models.CalculationData{
		Id:    powerAbilityScroll,
		Type:  "GEMSTONE_POWER_SCROLL",
		Price: prices[powerAbilityScroll] * constants.APPLICATION_WORTH["gemstonePowerScroll"],
		Count: 1,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
