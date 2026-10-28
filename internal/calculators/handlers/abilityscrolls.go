package handlers

import (
	"strings"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type AbilityScrollsHandler struct{}

func (h AbilityScrollsHandler) IsCosmetic() bool {
	return false
}

func (h AbilityScrollsHandler) Applies(item *models.NetworthItem) bool {
	return len(item.ExtraAttributes.AbilityScroll) > 0
}

func (h AbilityScrollsHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	for _, id := range item.ExtraAttributes.AbilityScroll {
		calculationData := models.CalculationData{
			Id:    id,
			Type:  "ABILITY_SCROLL",
			Price: prices[strings.ToUpper(id)] * constants.APPLICATION_WORTH["abilityScroll"],
			Count: 1,
		}

		item.Price += calculationData.Price
		item.Calculation = append(item.Calculation, calculationData)
	}
}
