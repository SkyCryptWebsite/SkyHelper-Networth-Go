package handlers

import (
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type TransmissionTunerHandler struct{}

func (h TransmissionTunerHandler) IsCosmetic() bool {
	return false
}

func (h TransmissionTunerHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.TunedTransmission > 0
}

func (h TransmissionTunerHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	tunedTransmissionCount := item.ExtraAttributes.TunedTransmission
	calculationData := models.CalculationData{
		Id:    "TRANSMISSION_TUNER",
		Type:  "TRANSMISSION_TUNER",
		Price: prices["TRANSMISSION_TUNER"] * float64(tunedTransmissionCount) * constants.APPLICATION_WORTH["tunedTransmission"],
		Count: tunedTransmissionCount,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
