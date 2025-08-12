package handlers

import (
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

type ReforgeHandler struct{}

func (h ReforgeHandler) IsCosmetic() bool {
	return false
}

func (h ReforgeHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.Modifier != "" && item.SkyblockItem.Category != "ACCESSORY"
}

func (h ReforgeHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	reforge := item.ExtraAttributes.Modifier
	if len(constants.REFORGES[reforge]) == 0 {
		return
	}

	reforgeId := constants.REFORGES[reforge]
	calculationData := models.CalculationData{
		Id:    reforgeId,
		Type:  "REFORGE",
		Price: prices[reforgeId] * constants.APPLICATION_WORTH["reforge"],
		Count: 1,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
