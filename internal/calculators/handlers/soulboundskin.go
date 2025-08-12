package handlers

import (
	"strings"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

type SoulboundSkinHandler struct{}

func (h SoulboundSkinHandler) IsCosmetic() bool {
	return true
}

func (h SoulboundSkinHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.Skin != "" && !strings.Contains(item.ItemId, item.ExtraAttributes.Skin) && item.IsSoulbound() && !item.NonCosmetic
}

func (h SoulboundSkinHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	skinId := item.ExtraAttributes.Skin
	if prices[skinId] == 0 {
		return
	}

	calculationData := models.CalculationData{
		Id:    skinId,
		Type:  "SOULBOUND_SKIN",
		Price: prices[skinId] * constants.APPLICATION_WORTH["soulboundSkins"],
		Count: 1,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
