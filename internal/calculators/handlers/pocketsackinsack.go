package handlers

import (
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type PocketSackInASackHandler struct{}

func (h PocketSackInASackHandler) IsCosmetic() bool {
	return false
}

func (h PocketSackInASackHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.SackPss > 0
}

func (h PocketSackInASackHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	sackCount := item.ExtraAttributes.SackPss
	calculationData := models.CalculationData{
		Id:    "POCKET_SACK_IN_A_SACK",
		Type:  "POCKET_SACK_IN_A_SACK",
		Price: prices["POCKET_SACK_IN_A_SACK"] * float64(sackCount) * constants.APPLICATION_WORTH["pocketSackInASack"],
		Count: sackCount,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
