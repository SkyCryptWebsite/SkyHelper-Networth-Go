package handlers

import (
	"fmt"
	"strings"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type BoosterHandler struct{}

func (h BoosterHandler) IsCosmetic() bool {
	return false
}

func (h BoosterHandler) Applies(item *models.NetworthItem) bool {
	return len(item.ExtraAttributes.Boosters) > 0
}

func (h BoosterHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
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
