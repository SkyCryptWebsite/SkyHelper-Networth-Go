package handlers

import (
	"github.com/duckysolucky/skyhelper-networth-go/internal/constants"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

type PotatoBookHandler struct{}

func (h PotatoBookHandler) IsCosmetic() bool {
	return false
}

func (h PotatoBookHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.HotPotatoCount > 0
}

func (h PotatoBookHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	potatoBookCount := item.ExtraAttributes.HotPotatoCount
	hotPotatoBookCount := min(potatoBookCount, 10.0)
	calculationData := models.CalculationData{
		Id:    "HOT_POTATO_BOOK",
		Type:  "HOT_POTATO_BOOK",
		Price: prices["HOT_POTATO_BOOK"] * float64(hotPotatoBookCount) * constants.APPLICATION_WORTH["hotPotatoBook"],
		Count: hotPotatoBookCount,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)

	if potatoBookCount > 10 {
		fumingPotatoBookCount := potatoBookCount - 10
		calculationData := models.CalculationData{
			Id:    "FUMING_POTATO_BOOK",
			Type:  "FUMING_POTATO_BOOK",
			Price: prices["FUMING_POTATO_BOOK"] * float64(fumingPotatoBookCount) * constants.APPLICATION_WORTH["fumingPotatoBook"],
			Count: fumingPotatoBookCount,
		}

		item.Price += calculationData.Price
		item.Calculation = append(item.Calculation, calculationData)
	}
}
