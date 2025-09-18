package handlers

import (
	"slices"
	"strings"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type RecombobulatorHandler struct{}

func (h RecombobulatorHandler) IsCosmetic() bool {
	return false
}

func (h RecombobulatorHandler) Applies(item *models.NetworthItem) bool {
	hasEnchantments := len(item.ExtraAttributes.Enchantments) > 0
	allowsRecombs := false
	if item.SkyblockItem != nil {
		category := item.SkyblockItem.Category
		allowsRecombs = slices.Contains(constants.AllowedRecombobulatedCategories, category) || slices.Contains(constants.ALLOWED_RECOMBOBULATED_IDS, item.ItemId)
	}

	lastLoreLine := ""
	if len(item.ItemLore) > 0 {
		lastLoreLine = item.ItemLore[len(item.ItemLore)-1]
	}

	isAccessory := strings.Contains(lastLoreLine, "ACCESSORY") || strings.Contains(lastLoreLine, "HATCESSORY")
	return item.IsRecombobulated() && (hasEnchantments || allowsRecombs || isAccessory)

}

func (h RecombobulatorHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	multiplier := constants.APPLICATION_WORTH["recombobulator"]
	if item.ItemId == "BONE_BOOMERANG" {
		multiplier *= 0.5
	}

	calculationData := models.CalculationData{
		Id:    "RECOMBOBULATOR_3000",
		Type:  "RECOMBOBULATOR_3000",
		Price: prices["RECOMBOBULATOR_3000"] * multiplier,
		Count: 1,
	}

	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
