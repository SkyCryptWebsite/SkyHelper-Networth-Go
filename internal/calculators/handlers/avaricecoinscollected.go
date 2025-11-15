package handlers

import (
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type AvariceCoinsCollectedHandler struct{}

func (h AvariceCoinsCollectedHandler) IsCosmetic() bool {
	return false
}

func (h AvariceCoinsCollectedHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.CollectedCoins > 0
}

func (h AvariceCoinsCollectedHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	zeroPrice := prices["CROWN_OF_AVARICE"]
	billionPrice := prices["CROWN_OF_AVARICE_1B"]
	coinsCollected := item.ExtraAttributes.CollectedCoins
	if coinsCollected > 1_000_000_000 {
		coinsCollected = 1_000_000_000
	}
	newPrice := zeroPrice + (billionPrice-zeroPrice)*float64(coinsCollected)/1_000_000_000

	calculationData := models.CalculationData{
		Id:    "CROWN_OF_AVARICE",
		Type:  "CROWN_OF_AVARICE",
		Price: newPrice,
		Count: int(coinsCollected),
	}
	item.BasePrice = calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
