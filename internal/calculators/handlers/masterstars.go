package handlers

import (
	"fmt"
	"strconv"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type MasterStarsHandler struct{}

func (h MasterStarsHandler) IsCosmetic() bool {
	return false
}

func getUpgradeLevel(item *models.NetworthItem) int {
	// Get dungeon_item_level, extract digits only, convert to int
	dungeonItemLevel := 0
	if item.ExtraAttributes.DungeonItemLevel != nil {
		dungeonItemLevelAsString := fmt.Sprintf("%v", item.ExtraAttributes.DungeonItemLevel)
		if parsed, err := strconv.Atoi(dungeonItemLevelAsString); err == nil {
			dungeonItemLevel = parsed
		} else {
			if parsed, err := strconv.Atoi(dungeonItemLevelAsString[:len(dungeonItemLevelAsString)-1]); err == nil {
				dungeonItemLevel = parsed
			}
		}
	}

	// Get upgrade_level, extract digits only, convert to int
	upgradeLevel := 0
	if item.ExtraAttributes.UpgradeLevel != nil {
		upgradeLevelAsString := fmt.Sprintf("%v", item.ExtraAttributes.UpgradeLevel)
		if parsed, err := strconv.Atoi(upgradeLevelAsString); err == nil {
			upgradeLevel = parsed
		} else {
			// If parsing fails, fall back to the original method
			if parsed, err := strconv.Atoi(upgradeLevelAsString[:len(upgradeLevelAsString)-1]); err == nil {
				upgradeLevel = parsed
			}
		}
	}

	return max(dungeonItemLevel, upgradeLevel)
}

func (h MasterStarsHandler) Applies(item *models.NetworthItem) bool {
	return item.SkyblockItem != nil && len(item.SkyblockItem.UpgradeCosts) > 0 && getUpgradeLevel(item) > 5
}

func (h MasterStarsHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	upgradeLevel := getUpgradeLevel(item) - 5
	starsUsed := min(upgradeLevel, 5)

	if len(item.SkyblockItem.UpgradeCosts) <= 5 {
		for i := range starsUsed {
			starId := constants.MASTER_STARS[i]
			calculationData := models.CalculationData{
				Id:    starId,
				Type:  "MASTER_STAR",
				Price: prices[starId] * constants.APPLICATION_WORTH["masterStar"],
				Count: 1,
			}

			item.Price += calculationData.Price
			item.Calculation = append(item.Calculation, calculationData)
		}
	}

}
