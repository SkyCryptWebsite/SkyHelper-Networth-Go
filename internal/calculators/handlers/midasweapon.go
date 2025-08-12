package handlers

import (
	"duckysolucky/skyhelper-networth-go/internal/constants"
	"duckysolucky/skyhelper-networth-go/internal/models"
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

func (h MidasWeaponHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
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
