package handlers

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type GemstonesHandler struct{}

func (h GemstonesHandler) IsCosmetic() bool {
	return false
}

func (h GemstonesHandler) Applies(item *models.NetworthItem) bool {
	return len(item.ExtraAttributes.Gems) > 0 && len(item.SkyblockItem.GemstoneSlots) > 0
}

func (h GemstonesHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	unlockedSlots := []string{}
	gems := []gemstone{}

	extraAttributesGems := item.ExtraAttributes.Gems

	var unlockedSlotsArr []string
	if unlockedSlotsInterface, ok := item.ExtraAttributes.Gems["unlocked_slots"].([]interface{}); ok {
		unlockedSlotsArr = make([]string, len(unlockedSlotsInterface))
		for i, v := range unlockedSlotsInterface {
			if str, ok := v.(string); ok {
				unlockedSlotsArr[i] = str
			}
		}
	} else if unlockedSlotsStringSlice, ok := item.ExtraAttributes.Gems["unlocked_slots"].([]string); ok {
		unlockedSlotsArr = unlockedSlotsStringSlice
	}

	// Create a copy of unlocked slots to avoid modifying while iterating
	remainingUnlockedSlots := make([]string, len(unlockedSlotsArr))
	copy(remainingUnlockedSlots, unlockedSlotsArr)

	// https://HypixelDev/PublicAPI/discussions/549
	for _, slot := range item.SkyblockItem.GemstoneSlots {
		// Check if this slot is unlocked (moved outside gem processing)
		if len(slot.Costs) > 0 && len(remainingUnlockedSlots) > 0 {
			// Find and remove matching unlocked slot
			for i, gemstoneId := range remainingUnlockedSlots {
				if strings.HasPrefix(gemstoneId, slot.SlotType) {
					unlockedSlots = append(unlockedSlots, slot.SlotType)
					// Remove the unlocked slot from remaining list
					remainingUnlockedSlots = append(remainingUnlockedSlots[:i], remainingUnlockedSlots[i+1:]...)
					break
				}
			}
		} else if len(slot.Costs) == 0 {
			unlockedSlots = append(unlockedSlots, slot.SlotType)
		} else if len(slot.Costs) > 0 && item.ExtraAttributes.Gems["unlocked_slots"] == nil {
			// If there are no costs and no unlocked_slots data, assume the slot is unlocked
			unlockedSlots = append(unlockedSlots, slot.SlotType)
		}

		// Process gems for this slot
		var key string
		for k := range extraAttributesGems {
			if strings.HasPrefix(k, slot.SlotType) && !strings.HasSuffix(k, "_gem") {
				key = k
				break
			}
		}

		if key != "" {
			var gemType string
			if slices.Contains(constants.GEMSTONE_SLOTS, slot.SlotType) {
				gemType, _ = extraAttributesGems[fmt.Sprintf("%s_gem", key)].(string)
			} else {
				gemType = slot.SlotType
			}

			var tier any
			if qualityObj, ok := extraAttributesGems[key].(map[string]any); ok {
				tier = qualityObj["quality"]
			} else {
				tier = extraAttributesGems[key]
			}

			gems = append(gems, gemstone{
				Type:     gemType,
				Tier:     tier,
				SlotType: slot.SlotType,
			})

			delete(extraAttributesGems, key)
		}
	}

	isDivansArmor, _ := regexp.MatchString(`^DIVAN_(HELMET|CHESTPLATE|LEGGINGS|BOOTS)$`, item.ItemId)
	isCrimsonArmor, _ := regexp.MatchString(`^(|HOT_|FIERY_|BURNING_|INFERNAL_)(AURORA|CRIMSON|TERROR|HOLLOW|FERVOR)(_HELMET|_CHESTPLATE|_LEGGINGS|_BOOTS)$`, item.ItemId)
	if isDivansArmor || isCrimsonArmor {
		application := constants.APPLICATION_WORTH["gemstoneSlots"]
		if isDivansArmor {
			application = constants.APPLICATION_WORTH["gemstoneChambers"]
		}

		GEMSTONE_SLOTS := make([]models.GemstoneSlot, len(item.SkyblockItem.GemstoneSlots))
		copy(GEMSTONE_SLOTS, item.SkyblockItem.GemstoneSlots)

		for _, unlockedSlot := range unlockedSlots {
			slotIndex := -1
			for i, s := range GEMSTONE_SLOTS {
				if s.SlotType == unlockedSlot {
					slotIndex = i
					break
				}
			}

			if slotIndex > -1 {
				slot := GEMSTONE_SLOTS[slotIndex]
				total := 0.0
				for _, cost := range slot.Costs {
					if cost.Type == "COINS" {
						total += float64(cost.Coins)
					} else if cost.Type == "ITEM" {
						price := prices[strings.ToUpper(cost.ItemId)]
						total += price * float64(cost.Amount)
					}
				}

				calculationData := models.CalculationData{
					Id:    unlockedSlot,
					Type:  "GEMSTONE_SLOT",
					Price: total * application,
					Count: 1,
				}

				item.Price += calculationData.Price
				item.Calculation = append(item.Calculation, calculationData)

				GEMSTONE_SLOTS = append(GEMSTONE_SLOTS[:slotIndex], GEMSTONE_SLOTS[slotIndex+1:]...)
			}
		}
	}

	for _, gemstone := range gems {
		gemstoneId := strings.ToUpper(fmt.Sprintf("%s_%s_GEM", gemstone.Tier, gemstone.Type))
		calculationData := models.CalculationData{
			Id:    gemstoneId,
			Type:  "GEMSTONE",
			Price: prices[gemstoneId] * constants.APPLICATION_WORTH["gemstone"],
			Count: 1,
		}

		item.Price += calculationData.Price
		item.Calculation = append(item.Calculation, calculationData)

	}
}

type gemstone struct {
	Type     string
	Tier     any
	SlotType string
}
