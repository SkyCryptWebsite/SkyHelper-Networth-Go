package handlers

import (
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/lib"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

type PrestigeHandler struct{}

func (h PrestigeHandler) IsCosmetic() bool {
	return false
}

func (h PrestigeHandler) Applies(item *models.NetworthItem) bool {
	for id := range constants.PRESTIGES {
		if id == item.ItemId {
			return true
		}
	}

	return false
}

func (h PrestigeHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	if prices[item.ItemId] > 0 {
		return
	}

	for _, prestigeItem := range constants.PRESTIGES[item.ItemId] {
		foundItem := lib.GetItem(prestigeItem)
		if len(foundItem.UpgradeCosts) > 0 {
			item.Price += getStarCosts(item, foundItem.UpgradeCosts, prestigeItem)
		}

		if len(foundItem.Prestige.Costs) > 0 {
			upgradeCostsv2 := [][]models.UpgradeCost{foundItem.Prestige.Costs} // Wrapped to match types (less work for fn and types)
			item.Price += getStarCosts(item, upgradeCostsv2, prestigeItem)

		}

		if prices[prestigeItem] > 0 {
			calculationData := models.CalculationData{
				Id:    prestigeItem,
				Type:  "BASE_PRESTIGE_ITEM",
				Price: prices[prestigeItem],
				Count: 1,
			}
			item.Price += calculationData.Price
			item.Calculation = append(item.Calculation, calculationData)
			break
		}
	}
}
