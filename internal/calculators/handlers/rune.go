package handlers

import (
	"fmt"
	"strings"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type RuneHandler struct{}

func (h RuneHandler) IsCosmetic() bool {
	return true
}

func (h RuneHandler) Applies(item *models.NetworthItem) bool {
	if strings.HasPrefix(item.ItemId, "RUNE") {
		return false
	}

	for _, runeCount := range item.ExtraAttributes.Runes {
		if runeCount > 0 {
			return true
		}
	}

	return false
}

func (h RuneHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	runes := item.ExtraAttributes.Runes
	for runeType, runeTier := range runes {
		runeId := fmt.Sprintf("RUNE_%s_%d", strings.ToUpper(runeType), runeTier)
		calculationData := models.CalculationData{
			Id:    runeId,
			Type:  "RUNE",
			Price: prices[runeId] * constants.APPLICATION_WORTH["runes"],
			Count: 1,
		}

		item.Price += calculationData.Price
		item.Calculation = append(item.Calculation, calculationData)
	}
}
