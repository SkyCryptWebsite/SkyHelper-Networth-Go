package handlers

import (
	"strings"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

type NecronBladeScrollHandler struct{}

func (h NecronBladeScrollHandler) IsCosmetic() bool {
	return false
}

func (h NecronBladeScrollHandler) Applies(item *models.NetworthItem) bool {
	return len(item.ExtraAttributes.AbilityScroll) > 0
}

func (h NecronBladeScrollHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	for _, id := range item.ExtraAttributes.AbilityScroll {
		calculationData := models.CalculationData{
			Id:    id,
			Type:  "NECRON_SCROLL",
			Price: prices[strings.ToUpper(id)] * constants.APPLICATION_WORTH["necronBladeScroll"],
			Count: 1,
		}

		item.Price += calculationData.Price
		item.Calculation = append(item.Calculation, calculationData)
	}
}
