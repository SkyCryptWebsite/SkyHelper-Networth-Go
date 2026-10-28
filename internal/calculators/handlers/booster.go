package handlers

import (
	"fmt"
	"strings"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type BoosterHandler struct{}

func (h BoosterHandler) IsCosmetic() bool {
	return false
}

func (h BoosterHandler) Applies(item *models.NetworthItem) bool {
	return len(item.ExtraAttributes.Boosters) > 0 || len(item.ExtraAttributes.BoosterTiers) > 0
}

func (h BoosterHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	// Newer items list each booster with the tier it was upgraded to: the base booster plus one upgrade per tier
	// above common (FORAGING_WISDOM_BOOSTER, then FORAGING_WISDOM_BOOSTER_UNCOMMON, ...).
	if len(item.ExtraAttributes.BoosterTiers) > 0 {
		for booster, tier := range item.ExtraAttributes.BoosterTiers {
			boosterIds := []string{fmt.Sprintf("%s_BOOSTER", strings.ToUpper(booster))}
			for t := 1; t < tier && t < len(constants.RARITIES); t++ {
				boosterIds = append(boosterIds, fmt.Sprintf("%s_BOOSTER_%s", strings.ToUpper(booster), constants.RARITIES[t]))
			}

			for _, boosterId := range boosterIds {
				price := prices[boosterId]
				// Some base boosters are only sold with their rarity, like HUNTING_WISDOM_BOOSTER_COMMON.
				if price == 0 && strings.HasSuffix(boosterId, "_BOOSTER") {
					boosterId += "_COMMON"
					price = prices[boosterId]
				}
				h.add(item, boosterId, price)
			}
		}
		return
	}

	for _, booster := range item.ExtraAttributes.Boosters {
		boosterId := fmt.Sprintf("%s_BOOSTER", strings.ToUpper(booster))
		h.add(item, boosterId, prices[boosterId])
	}
}

func (h BoosterHandler) add(item *models.NetworthItem, boosterId string, price float64) {
	if price == 0 {
		return
	}
	calculationData := models.CalculationData{
		Id:    boosterId,
		Type:  "BOOSTER",
		Price: price * constants.APPLICATION_WORTH["booster"],
		Count: 1,
	}
	item.Price += calculationData.Price
	item.Calculation = append(item.Calculation, calculationData)
}
