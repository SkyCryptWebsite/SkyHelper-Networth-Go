package models

import (
	"fmt"
	"slices"
	"sort"

	skycrypttypes "github.com/SkyCryptWebsite/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
)

type NetworthPet struct {
	PetData skycrypttypes.Pet `json:"petData"`

	NonCosmetic bool   `json:"nonCosmetic"`
	Skin        string `json:"string"`
	BasePetId   string `json:"basePetId"`
	PetId       string `json:"petId"`
	Level       Level  `json:"level"`
	PetName     string `json:"petName"`

	Prices map[string]float64 `json:"prices"`

	Calculation []CalculationData `json:"calculation"`
	BasePrice   float64           `json:"basePrice"`
	Price       float64           `json:"price"`
}

type Level struct {
	Level           int `json:"level"`
	ExperienceToMax int `json:"xpMax"`
	// ExperienceTo100 is the XP the pet's rarity needs for level 100, which prices below 100 scale against.
	ExperienceTo100 int     `json:"xpMaxTo100"`
	Experience      float64 `json:"xp"`
}

func (item *NetworthPet) getTier() int {
	return slices.Index(constants.RARITIES, item.PetData.Rarity)
}

func (item *NetworthPet) getTierName() string {
	return item.PetData.Rarity
}

// hasKnownRarity is false for rarities missing from constants.RARITIES; such pets are left unpriced.
func (item *NetworthPet) hasKnownRarity() bool {
	return item.getTier() >= 0
}

func (item *NetworthPet) getTierBoostedTier() int {
	tier := item.getTier()
	// A Tier Boost doesn't raise a Phoenix's rarity.
	if tier >= 0 && tier+1 < len(constants.RARITIES) && item.PetData.HeldItem == "PET_ITEM_TIER_BOOST" && item.PetData.Type != "PHOENIX" {
		return tier + 1
	}

	return tier
}

func (item *NetworthPet) getTierBoostedTierName() string {
	tier := item.getTierBoostedTier()
	if tier < 0 {
		return item.getTierName()
	}

	return constants.RARITIES[tier]
}

func (item *NetworthPet) IsSoulbound() bool {
	return slices.Contains(constants.SOULBOUND_PETS, item.PetData.Type)
}

func (item *NetworthPet) IsCosmetic() bool {
	return item.Skin != ""
}

func (item *NetworthPet) getPetLevelPrices() map[string]float64 {
	basePrices := map[string]float64{
		"LVL_1":   item.Prices[fmt.Sprintf("LVL_1_%s", item.BasePetId)],
		"LVL_100": item.Prices[fmt.Sprintf("LVL_100_%s", item.BasePetId)],
		"LVL_200": item.Prices[fmt.Sprintf("LVL_200_%s", item.BasePetId)],
	}

	// If the pet has a skin, use the max between skinnedPrice and basePrice
	if item.Skin != "" && !item.NonCosmetic {
		return map[string]float64{
			"LVL_1":   max(item.Prices[fmt.Sprintf("LVL_1_%s", item.PetId)], basePrices["LVL_1"]),
			"LVL_100": max(item.Prices[fmt.Sprintf("LVL_100_%s", item.PetId)], basePrices["LVL_100"]),
			"LVL_200": max(item.Prices[fmt.Sprintf("LVL_200_%s", item.PetId)], basePrices["LVL_200"]),
		}
	}

	return basePrices
}

func (item *NetworthPet) getBasePrice() {
	if !item.hasKnownRarity() {
		return
	}

	petLevelPrices := item.getPetLevelPrices()

	basePrice := petLevelPrices["LVL_200"]
	if basePrice == 0 {
		basePrice = petLevelPrices["LVL_100"]
	}

	// Up to level 100 a pet is worth its level 100 price, even one that can reach 200.
	if item.Level.Level <= 100 {
		basePrice = petLevelPrices["LVL_100"]
	}
	item.BasePrice = basePrice

	// Below level 100, scale between the level 1 and level 100 prices by the XP towards level 100.
	if item.Level.Level < 100 && item.Level.ExperienceToMax > 0 && item.Level.ExperienceTo100 > 0 {
		baseFormula := (petLevelPrices["LVL_100"] - petLevelPrices["LVL_1"]) / float64(item.Level.ExperienceTo100)
		if baseFormula != 0 {
			item.BasePrice = baseFormula*item.Level.Experience + petLevelPrices["LVL_1"]
		}
	}

	// Between 100 and 200, scale between the level 100 and level 200 prices by level.
	if item.Level.Level > 100 && item.Level.Level < 200 {
		level := item.Level.Level - 100
		baseFormula := (petLevelPrices["LVL_200"] - petLevelPrices["LVL_100"]) / 100
		if baseFormula != 0 {
			item.BasePrice = baseFormula*float64(level) + petLevelPrices["LVL_100"]
		}
	}
}

func (item *NetworthPet) GetPetId() string {
	petLevelPrices := item.getPetLevelPrices()
	if petLevelPrices["LVL_200"] > 0 {
		return fmt.Sprintf("LVL_200_%s", item.PetId)
	} else if petLevelPrices["LVL_100"] > 0 {
		return fmt.Sprintf("LVL_100_%s", item.PetId)
	}

	return fmt.Sprintf("LVL_1_%s", item.PetId)
}

func (item *NetworthPet) getPetLevel() Level {
	maxPetLevel := constants.PETS_SPECIAL_LEVELS[item.PetData.Type]
	if maxPetLevel == 0 {
		maxPetLevel = 100
	}

	rarity := item.getTierBoostedTierName()
	if item.PetData.Type == "BINGO" {
		// The Bingo pet levels like a common pet, whatever its rarity.
		rarity = "COMMON"
	}
	petOffset := constants.RARITY_OFFSET_PETS[rarity]
	petLevels := constants.PET_LEVELS[petOffset:min(petOffset+maxPetLevel-1, len(constants.PET_LEVELS))]

	level := 1
	totalExp := 0.0
	for i := range min(len(petLevels), maxPetLevel) {
		totalExp += float64(petLevels[i])
		if totalExp > item.PetData.Experience {
			totalExp -= float64(petLevels[i])
			break
		}

		level++
	}

	sumOfExp := 0.0
	for i := range min(len(petLevels), maxPetLevel) {
		sumOfExp += float64(petLevels[i])
	}

	return Level{
		Level:           min(level, maxPetLevel),
		ExperienceToMax: int(sumOfExp),
		ExperienceTo100: constants.XP_TO_LEVEL_100[item.getTierBoostedTierName()],
		Experience:      item.PetData.Experience,
	}
}

func NewSkyBlockPetCalculator(item *skycrypttypes.Pet, prices map[string]float64, options NetworthOptions) *NetworthPet {
	networthItem := &NetworthPet{
		PetData: *item,

		NonCosmetic: options.NonCosmetic,
		Skin:        item.Skin,
		// BasePetId:   "",
		// PetId:       "",
		// Level:       Level{},
		// PetName:     "",

		Prices: prices,

		Calculation: []CalculationData{},
		BasePrice:   0.0,
		Price:       0.0,
	}

	networthItem.BasePetId = fmt.Sprintf("%s_%s", networthItem.getTierName(), networthItem.PetData.Type)
	petId := networthItem.BasePetId
	if networthItem.Skin != "" {
		petId = fmt.Sprintf("%s_SKINNED_%s", petId, networthItem.Skin)
	}

	networthItem.PetId = petId
	networthItem.Level = networthItem.getPetLevel()

	rarity := titleCase(networthItem.getTierBoostedTierName())
	customPetName := constants.CUSTOM_PET_NAMES[networthItem.PetData.Type]
	if customPetName == "" {
		customPetName = titleCase(networthItem.PetData.Type)
	}

	networthItem.PetName = fmt.Sprintf("[Lvl %d] %s %s", networthItem.Level.Level, rarity, customPetName)

	networthItem.getBasePrice()

	return networthItem
}

func (item *NetworthPet) Calculate(handlers []PetHandler) {
	if !item.hasKnownRarity() {
		return
	}

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

func (item *NetworthPet) GetPrice() float64 {
	return item.Price + item.BasePrice
}

func (item *NetworthPet) GetCalculation() []CalculationData {
	sort.Slice(item.Calculation, func(i, j int) bool {
		return item.Calculation[i].Price > item.Calculation[j].Price
	})

	return item.Calculation
}
