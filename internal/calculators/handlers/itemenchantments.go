package handlers

import (
	"fmt"
	"slices"
	"strings"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

type ItemEnchantments struct{}

func (h ItemEnchantments) IsCosmetic() bool {
	return false
}

func (h ItemEnchantments) Applies(item *models.NetworthItem) bool {
	return item.ItemId != "ENCHANTED_BOOK" && len(item.ExtraAttributes.Enchantments) > 0
}

func (h ItemEnchantments) Calculate(item *models.NetworthItem, prices models.Prices) {
	for id, level := range item.ExtraAttributes.Enchantments {
		upperCasedId := strings.ToUpper(id)
		if slices.Contains(constants.BLOCKED_ENCHANTMENTS[item.ItemId], upperCasedId) {
			continue
		}

		if constants.IGNORED_ENCHANTMENTS[upperCasedId] == level {
			continue
		}

		// Set stacking enchantments to 1 since that is the only value we track
		if slices.Contains(constants.STACKING_ENCHANTMENTS, upperCasedId) {
			level = 1
		}

		if upperCasedId == "EFFICIENCY" && level >= 6 && !slices.Contains(constants.IGNORE_SILEX, item.ItemId) {
			defaultMaxEfficiencyLevel := 5
			if item.ItemId == "STONK_PICKAXE" {
				defaultMaxEfficiencyLevel = 6
			}

			efficiencyLevel := level - defaultMaxEfficiencyLevel
			if efficiencyLevel > 0 {
				calculationData := models.CalculationData{
					Id:    "SIL_EX",
					Type:  "SILEX",
					Price: prices["SIL_EX"] * float64(efficiencyLevel) * constants.APPLICATION_WORTH["silex"],
					Count: efficiencyLevel,
				}

				item.Price += calculationData.Price
				item.Calculation = append(item.Calculation, calculationData)
			}
		}

		for enchantmentId, enchantmentData := range constants.ENCHANTMENT_UPGRADES {
			if upperCasedId == enchantmentId && level >= enchantmentData.Tier {
				calculationData := models.CalculationData{
					Id:    enchantmentData.UpgradeItem,
					Type:  "ENCHANTMENT_UPGRADE",
					Price: prices[enchantmentData.UpgradeItem] * constants.APPLICATION_WORTH["enchantmentUpgrades"],
					Count: 1,
				}

				item.Price += calculationData.Price
				item.Calculation = append(item.Calculation, calculationData)
			}
		}

		formattedId := fmt.Sprintf("%s_%d", upperCasedId, level)
		formattedEnchantmentID := fmt.Sprintf("ENCHANTMENT_%s", formattedId)
		if prices[formattedEnchantmentID] == 0 {
			continue
		}

		multiplier := constants.ENCHANTMENTS_WORTH[upperCasedId]
		if multiplier == 0 {
			multiplier = constants.APPLICATION_WORTH["enchantments"]
		}

		calculationData := models.CalculationData{
			Id:    formattedId,
			Type:  "ENCHANTMENT",
			Price: prices[formattedEnchantmentID] * multiplier,
			Count: 1,
		}

		item.Price += calculationData.Price
		item.Calculation = append(item.Calculation, calculationData)
	}
}
