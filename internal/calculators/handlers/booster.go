package handlers

import (
	"fmt"
	"strings"

	"github.com/duckysolucky/skyhelper-networth-go/internal/constants"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

type BoosterHandler struct{}

func (h BoosterHandler) IsCosmetic() bool {
	return false
}

func (h BoosterHandler) Applies(item *models.NetworthItem) bool {
	return len(item.ExtraAttributes.Boosters) > 0
}

func (h BoosterHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	for _, booster := range item.ExtraAttributes.Boosters {
		boosterId := fmt.Sprintf("%s_BOOSTER", strings.ToUpper(booster))
		calculationData := models.CalculationData{
			Id:    boosterId,
			Type:  "BOOSTER",
			Price: prices[boosterId] * constants.APPLICATION_WORTH["booster"],
			Count: 1,
		}

		item.Price += calculationData.Price
		item.Calculation = append(item.Calculation, calculationData)
	}
}
