package handlers

import (
	"duckysolucky/skyhelper-networth-go/internal/constants"
	"duckysolucky/skyhelper-networth-go/internal/models"
)

type PulseRingThunderHandler struct{}

func (h PulseRingThunderHandler) IsCosmetic() bool {
	return false
}

func (h PulseRingThunderHandler) Applies(item *models.NetworthItem) bool {
	return item.ItemId == "PULSE_RING" && item.ExtraAttributes.ThunderCharge > 0
}

func (h PulseRingThunderHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	thunderUpgrades := min(item.ExtraAttributes.ThunderCharge, constants.MAX_THUNDER_CHARGE) / 50_000
	calculationData := models.CalculationData{
		Id:    "THUNDER_IN_A_BOTTLE",
		Type:  "THUNDER_CHARGE",
		Price: prices["THUNDER_IN_A_BOTTLE"] * float64(thunderUpgrades) * constants.APPLICATION_WORTH["thunderInABottle"],
		Count: thunderUpgrades,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
