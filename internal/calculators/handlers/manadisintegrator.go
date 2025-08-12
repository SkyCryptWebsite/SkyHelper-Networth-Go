package handlers

import (
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

type ManaDisintegratorHandler struct{}

func (h ManaDisintegratorHandler) IsCosmetic() bool {
	return false
}

func (h ManaDisintegratorHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.ManaDisintegrator > 0
}

func (h ManaDisintegratorHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	manaDisintegratorCount := item.ExtraAttributes.ManaDisintegrator
	calculationData := models.CalculationData{
		Id:    "MANA_DISINTEGRATOR",
		Type:  "MANA_DISINTEGRATOR",
		Price: prices["MANA_DISINTEGRATOR"] * float64(manaDisintegratorCount) * constants.APPLICATION_WORTH["manaDisintegrator"],
		Count: manaDisintegratorCount,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
