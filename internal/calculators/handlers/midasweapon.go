package handlers

import (
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type MidasWeaponHandler struct{}

func (h MidasWeaponHandler) IsCosmetic() bool {
	return false
}

func (h MidasWeaponHandler) Applies(item *models.NetworthItem) bool {
	for key := range constants.MIDAS_SWORDS {
		if key == item.ItemId {
			return true
		}
	}

	return false
}

func (h MidasWeaponHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	maxBid := constants.MIDAS_SWORDS[item.ItemId].MaxBid
	swordType := constants.MIDAS_SWORDS[item.ItemId].Type
	winningBid := item.ExtraAttributes.WinningBid
	additionalCoins := item.ExtraAttributes.AdditionalCoins

	if (winningBid+additionalCoins) >= maxBid && prices[swordType] > 0 {
		calculationData := models.CalculationData{
			Id:    item.ItemId,
			Type:  swordType,
			Price: prices[swordType],
			Count: 1,
		}

		item.BasePrice = calculationData.Price
		item.Calculation = append(item.Calculation, calculationData)
	}
}
