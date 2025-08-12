package handlers

import (
	"github.com/duckysolucky/skyhelper-networth-go/internal/constants"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

type EtherwarpConduitHandler struct{}

func (h EtherwarpConduitHandler) IsCosmetic() bool {
	return false
}

func (h EtherwarpConduitHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.Ethermerge > 0
}

func (h EtherwarpConduitHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	calculationData := models.CalculationData{
		Id:    "ETHERWARP_CONDUIT",
		Type:  "ETHERWARP_CONDUIT",
		Price: prices["ETHERWARP_CONDUIT"] * constants.APPLICATION_WORTH["etherwarp"],
		Count: 1,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
