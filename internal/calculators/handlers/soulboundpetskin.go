package handlers

import (
	"fmt"

	"duckysolucky/skyhelper-networth-go/internal/constants"
	"duckysolucky/skyhelper-networth-go/internal/models"
)

type SoulboundPetSkinHandler struct{}

func (h SoulboundPetSkinHandler) IsCosmetic() bool {
	return false
}

func (h SoulboundPetSkinHandler) Applies(item *models.NetworthPet) bool {
	return item.PetData.Skin != "" && item.IsSoulbound() && !item.NonCosmetic
}

func (h SoulboundPetSkinHandler) Calculate(item *models.NetworthPet, prices models.Prices) {
	petSkinId := fmt.Sprintf("PET_SKIN_%s", item.PetData.Skin)
	if prices[petSkinId] == 0 {
		return
	}

	calculationData := models.CalculationData{
		Id:    item.PetData.Skin,
		Type:  "SOULBOUND_PET_SKIN",
		Price: prices[petSkinId] * constants.APPLICATION_WORTH["soulboundPetSkins"],
		Count: 1,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
