package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/lib"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type EnchantedBookHandler struct{}

func (h EnchantedBookHandler) IsCosmetic() bool {
	return false
}

func (h EnchantedBookHandler) Applies(item *models.NetworthItem) bool {
	return item.ItemId == "ENCHANTED_BOOK" && len(item.ExtraAttributes.Enchantments) > 0
}

func (h EnchantedBookHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	enchantmentPrice := 0.0
	isSingleBook := len(item.ExtraAttributes.Enchantments) == 1
	for enchantment, level := range item.ExtraAttributes.Enchantments {
		enchantmentId := fmt.Sprintf("ENCHANTMENT_%s_%d", strings.ToUpper(enchantment), level)
		price := prices[enchantmentId]
		if price == 0 {
			continue
		}

		multiplier := 1.0
		if !isSingleBook {
			multiplier = constants.APPLICATION_WORTH["enchantments"]
		}

		calculationData := models.CalculationData{
			Id:    strings.ToUpper(enchantment + "_" + strconv.Itoa(level)),
			Type:  "ENCHANT",
			Price: price * multiplier,
			Count: 1,
		}

		enchantmentPrice += calculationData.Price
		item.Calculation = append(item.Calculation, calculationData)
		if isSingleBook {
			if val, ok := constants.SPECIAL_ENCHANTMENT_NAMES[enchantment]; ok {
				item.ItemName = val
			} else {
				item.ItemName = lib.TitleCase(strings.ReplaceAll(enchantment, "_", " "))
			}
		}
	}

	if enchantmentPrice > 0 {
		item.BasePrice = enchantmentPrice
	}
}
