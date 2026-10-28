package models

import (
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strings"

	skycrypttypes "github.com/SkyCryptWebsite/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
)

// titleCase converts a string to title case
func titleCase(str string) string {
	if str == "" {
		return ""
	}
	words := strings.Split(strings.ToLower(strings.ReplaceAll(str, "_", " ")), " ")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

// ItemProvider is an interface for getting item data
type ItemProvider interface {
	GetItem(itemId string) *HypixelItem
}

type NetworthItem struct {
	ItemName        string                         `json:"itemName"`
	ExtraAttributes *skycrypttypes.ExtraAttributes `json:"extraAttributes"`
	ItemId          string                         `json:"itemId"`
	SkyblockItem    *HypixelItem                   `json:"skyblockItem"`
	ItemLore        []string                       `json:"itemLore"`
	Count           int                            `json:"count"`
	BaseItemId      string                         `json:"baseItemId"`

	Prices map[string]float64 `json:"prices"`

	NonCosmetic      bool              `json:"nonCosmetic"`
	Calculation      []CalculationData `json:"calculation"`
	BasePrice        float64           `json:"basePrice"`
	Price            float64           `json:"price"`
	SoulboundPortion float64           `json:"soulboundPortion"`
}

func NewSkyBlockItemCaclulator(item *skycrypttypes.Item, prices map[string]float64, itemProvider ItemProvider, options NetworthOptions) *NetworthItem {
	// Minecraft treats a missing Count as a single item.
	count := 1
	if item.Count != nil {
		count = *item.Count
	}

	networthItem := &NetworthItem{
		// Cleaned here, as upstream does, so every name check (cosmetic detection included) sees the plain name.
		ItemName:        REMOVE_PERCENTAGE_COLOR_CODES_REGEX.ReplaceAllString(REMOVE_COLOR_CODES_REGEX.ReplaceAllString(item.Tag.Display.Name, ""), ""),
		ExtraAttributes: item.Tag.ExtraAttributes,
		ItemId:          item.Tag.ExtraAttributes.Id,
		SkyblockItem:    itemProvider.GetItem(item.Tag.ExtraAttributes.Id),
		ItemLore:        item.Tag.Display.Lore,
		Count:           count,
		BaseItemId:      item.Tag.ExtraAttributes.Id,

		Prices: prices,

		NonCosmetic:      options.NonCosmetic,
		Calculation:      []CalculationData{},
		BasePrice:        0.0,
		Price:            0.0,
		SoulboundPortion: 0.0,
	}

	networthItem.GetBasePrice()

	return networthItem
}

func (item *NetworthItem) isRune() bool {
	return item.ItemId == "RUNE" && len(item.ExtraAttributes.Runes) > 0
}

func (item *NetworthItem) isUniqueRune() bool {
	return item.ItemId == "UNIQUE_RUNE" && len(item.ExtraAttributes.Runes) > 0
}

func (item *NetworthItem) getItemId() string {
	if item.ExtraAttributes.Skin != "" && !item.NonCosmetic {
		itemId := fmt.Sprintf("%s_SKINNED_%s", item.ItemId, item.ExtraAttributes.Skin)
		if item.Prices[itemId] > item.Prices[item.ItemId] {
			return itemId
		}
	}

	if item.ItemId == "PARTY_HAT_SLOTH" && item.ExtraAttributes.PartyHatEmoji != "" {
		itemId := fmt.Sprintf("%s_%s", item.ItemId, strings.ToUpper(item.ExtraAttributes.PartyHatEmoji))
		if item.Prices[itemId] > 0 {
			return itemId
		}
	}

	if (item.isRune() || item.isUniqueRune()) && !item.NonCosmetic {
		for runeType, runeTier := range item.ExtraAttributes.Runes {
			if runeTier > 0 {
				return fmt.Sprintf("RUNE_%s_%d", strings.ToUpper(runeType), runeTier)
			}
		}
	}

	if item.ItemId == "NEW_YEAR_CAKE" {
		return fmt.Sprintf("NEW_YEAR_CAKE_%d", item.ExtraAttributes.NewYearsCake)
	}

	partyHatAccessories := []string{"PARTY_HAT_CRAB", "PARTY_HAT_CRAB_ANIMATED", "BALLOON_HAT_2024", "BALLOON_HAT_2025"}
	if slices.Contains(partyHatAccessories, item.ItemId) && item.ExtraAttributes.PartyHatColor != "" {
		return fmt.Sprintf("%s_%s", item.ItemId, strings.ToUpper(item.ExtraAttributes.PartyHatColor))
	}

	if item.ItemId == "DCTR_SPACE_HELM" && item.ExtraAttributes.Edition != 0 {
		return "DCTR_SPACE_HELM_EDITIONED"
	}

	if item.ItemId == "CREATIVE_MIND" && item.ExtraAttributes.Edition == 0 {
		return "CREATIVE_MIND_UNEDITIONED"
	}

	if item.ItemId == "ANCIENT_ELEVATOR" && item.ExtraAttributes.Edition != 0 {
		return "ANCIENT_ELEVATOR_EDITIONED"
	}

	if strings.HasPrefix(item.ItemId, "STARRED_") && item.Prices[item.ItemId] == 0 && item.Prices[strings.ReplaceAll(item.ItemId, "STARRED_", "")] > 0 {
		return strings.ReplaceAll(item.ItemId, "STARRED_", "")
	}

	itemId := fmt.Sprintf("%s_SHINY", item.ItemId)
	if item.ExtraAttributes.Shiny && item.Prices[itemId] > 0 {
		return itemId
	}

	return item.ItemId
}

var REMOVE_COLOR_CODES_REGEX = regexp.MustCompile(`§[0-9a-fk-orA-FK-OR]`)
var REMOVE_PERCENTAGE_COLOR_CODES_REGEX = regexp.MustCompile(`%%[^%]+%%`)

func (item *NetworthItem) getItemName() string {
	name := REMOVE_COLOR_CODES_REGEX.ReplaceAllString(item.ItemName, "")
	name = REMOVE_PERCENTAGE_COLOR_CODES_REGEX.ReplaceAllString(name, "")

	itemsWithRarites := []string{"Beastmaster Crest", "Griffin Upgrade Stone", "Wisp Upgrade Stone"}
	if slices.Contains(itemsWithRarites, name) && item.SkyblockItem != nil {
		rarity := item.SkyblockItem.Rarity
		if rarity == "" {
			rarity = "common"
		}

		return fmt.Sprintf("%s (%s)", name, titleCase(rarity))
	}

	if strings.HasSuffix(name, "Exp Boost") && item.SkyblockItem != nil {
		var itemId = "Unknown"
		if item.SkyblockItem.SkyBlockID != "" {
			parts := strings.Split(item.SkyblockItem.SkyBlockID, "_")
			itemId = titleCase(parts[len(parts)-1])
		}

		return fmt.Sprintf("%s (%s)", name, itemId)

	}

	return name
}

func (item *NetworthItem) IsCosmetic() bool {
	// Upstream checks twice, before pricing (with the base item ID) and after (with the priced one, such as
	// X_SKINNED_Y). It glues the ID straight onto the name, so this matches its results exactly, quirks included:
	// "PARTY_HAT_CRAB_ANIMATED" + "Yellow Crab Hat" contains "DYE".
	isSkinOrDye := false
	for _, id := range []string{item.BaseItemId, item.ItemId} {
		testId := strings.ToUpper(id + item.ItemName)
		isSkinOrDye = isSkinOrDye || strings.Contains(testId, "DYE") || strings.Contains(testId, "SKIN")
	}
	// The lore's last line marks cosmetics even for items missing from the Hypixel item list.
	isCosmetic := len(item.ItemLore) > 0 && strings.Contains(item.ItemLore[len(item.ItemLore)-1], "COSMETIC")
	isMemento := false
	if item.SkyblockItem != nil {
		isCosmetic = isCosmetic || item.SkyblockItem.Category == "COSMETIC"
		isMemento = item.SkyblockItem.Category == "MEMENTO"
	}
	isOnCosmeticBlacklist := slices.Contains(constants.NON_COSMETIC_ITEMS, item.BaseItemId)
	return isCosmetic || isSkinOrDye || isMemento || isOnCosmeticBlacklist || item.isRune() || item.isUniqueRune()
}

func (item *NetworthItem) IsRecombobulated() bool {
	// ? NOTE: iItemTier is rarity obtained in dungeons from a drop (x/50 score)
	return item.ExtraAttributes.Recombobulated > 0 && !(item.ExtraAttributes.ItemTier > 0)
}

func (item *NetworthItem) IsSoulbound() bool {
	return item.ExtraAttributes.Soulbound || slices.Contains(item.ItemLore, "§8§l* §8Co-op Soulbound §8§l*") || slices.Contains(item.ItemLore, "§8§l* §8Soulbound §8§l*")
}

func (item *NetworthItem) GetBasePrice() {
	item.ItemName = item.getItemName()
	item.ItemId = item.getItemId()

	itemPrice := item.Prices[item.ItemId]
	item.BasePrice = itemPrice * float64(item.Count)
}

func (item *NetworthItem) Calculate(handlers []Handler) {
	for _, handler := range handlers {
		if item.NonCosmetic && handler.IsCosmetic() {
			continue
		}

		if !handler.Applies(item) {
			continue
		}

		handler.Calculate(item, item.Prices)
	}
}

func (item *NetworthItem) GetPrice() float64 {
	return item.Price + item.BasePrice
}

func (item *NetworthItem) GetCalculation() []CalculationData {
	sort.Slice(item.Calculation, func(i, j int) bool {
		return item.Calculation[i].Price > item.Calculation[j].Price
	})

	return item.Calculation
}
