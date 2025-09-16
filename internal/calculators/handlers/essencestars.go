package handlers

import (
	"fmt"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type EssenceStarsHandler struct{}

func (h EssenceStarsHandler) IsCosmetic() bool {
	return false
}

func (h EssenceStarsHandler) Applies(item *models.NetworthItem) bool {
	return item.SkyblockItem != nil && len(item.SkyblockItem.UpgradeCosts) > 0 && getUpgradeLevel(item) > 0
}

func getStarCosts(item *models.NetworthItem, upgrades [][]models.UpgradeCost, prestigeItem string) float64 {
	price, star := 0.0, 0
	allCalculations := []*models.CalculationData{}
	for _, root := range upgrades {
		star++
		for _, upgrade := range root {
			// ? NOTE: MIGHT NEED TO BE CHANGED
			if prestigeItem != "" && len(upgrades) == 1 {
				star = 0
			}

			calculationData := starCost(item, upgrade, star)
			allCalculations = append(allCalculations, calculationData)
			if prestigeItem == "" && calculationData != nil {
				price += calculationData.Price
				item.Calculation = append(item.Calculation, *calculationData)
			}
		}
	}

	if prestigeItem != "" && len(allCalculations) > 0 {
		prestige := allCalculations[0].Type == "PRESTIGE"
		totalCost := 0.0
		for _, calculation := range allCalculations {
			totalCost += calculation.Price
		}

		calculationData := &models.CalculationData{
			Id:    prestigeItem,
			Type:  "STARS",
			Price: totalCost,
			Count: star,
		}

		if prestige {
			calculationData.Type = "PRESTIGE"
			calculationData.Count = 1
		}

		item.Calculation = append(item.Calculation, *calculationData)
		price += calculationData.Price
	}

	return price
}

func starCost(item *models.NetworthItem, upgrade models.UpgradeCost, star int) *models.CalculationData {
	upgradeType := upgrade.ItemId
	if upgradeType == "" || upgrade.EssenceType != "" {
		upgradeType = fmt.Sprintf("ESSENCE_%s", upgrade.EssenceType)
	}

	upgradePrice := item.Prices[upgradeType]
	if upgradePrice == 0 {
		return &models.CalculationData{}
	}

	calculationType := "PRESTIGE"
	if star > 0 {
		calculationType = "STAR"
	}

	APPLICATION_WORTH := 1.0
	if upgrade.EssenceType != "" {
		APPLICATION_WORTH = constants.APPLICATION_WORTH["essence"]
	}

	upgradeId := upgrade.ItemId
	if upgrade.EssenceType != "" {
		upgradeId = fmt.Sprintf("%s_ESSENCE", upgrade.EssenceType)
	}

	return &models.CalculationData{
		Id:    upgradeId,
		Type:  calculationType,
		Price: float64(upgrade.Amount) * upgradePrice * APPLICATION_WORTH,
		Count: upgrade.Amount,
		Star:  star,
	}
}

func (h EssenceStarsHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	level := min(getUpgradeLevel(item), len(item.SkyblockItem.UpgradeCosts))
	item.Price += getStarCosts(item, item.SkyblockItem.UpgradeCosts[:level], "")
}
