package handlers

import (
	"fmt"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

type NewYearCakeBagHandler struct{}

func (h NewYearCakeBagHandler) IsCosmetic() bool {
	return false
}

func (h NewYearCakeBagHandler) Applies(item *models.NetworthItem) bool {
	return len(item.ExtraAttributes.NewYearCakeBagYears) > 0
}

func (h NewYearCakeBagHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	cakePrices := 0.0
	for _, year := range item.ExtraAttributes.NewYearCakeBagYears {
		cakeId := fmt.Sprintf("NEW_YEAR_CAKE_%d", year)
		cakePrices += prices[cakeId]
	}

	calculationData := models.CalculationData{
		Id:    "NEW_YEAR_CAKES",
		Type:  "NEW_YEAR_CAKES",
		Price: cakePrices,
		Count: 1,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
