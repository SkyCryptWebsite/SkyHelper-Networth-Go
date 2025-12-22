package handlers

import (
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type Overclocker3000Handler struct{}

func (h Overclocker3000Handler) IsCosmetic() bool {
	return false
}

func (h Overclocker3000Handler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.Overclocker3000 > 0
}

func (h Overclocker3000Handler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	calculationData := models.CalculationData{
		Id:    "OVERCLOCKER_3000",
		Type:  "OVERCLOCKER_3000",
		Price: prices["OVERCLOCKER_3000"] * float64(item.ExtraAttributes.Overclocker3000) * constants.APPLICATION_WORTH["overclocker3000"],
		Count: item.ExtraAttributes.Overclocker3000,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
