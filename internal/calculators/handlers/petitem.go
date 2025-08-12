package handlers

import (
	"strings"

	"duckysolucky/skyhelper-networth-go/internal/constants"
	"duckysolucky/skyhelper-networth-go/internal/models"
)

type PetItemHandler struct{}

func (h PetItemHandler) IsCosmetic() bool {
	return false
}

func (h PetItemHandler) Applies(item *models.NetworthPet) bool {
	return item.PetData.HeldItem != ""

}

func (h PetItemHandler) Calculate(item *models.NetworthPet, prices models.Prices) {
	petItem := item.PetData.HeldItem
	calculationData := models.CalculationData{
		Id:    strings.ToUpper(petItem),
		Type:  "PET_ITEM",
		Price: prices[petItem] * constants.APPLICATION_WORTH["petItem"],
		Count: 1,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
