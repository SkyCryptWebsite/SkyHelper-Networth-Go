package models

import (
	"fmt"
	"slices"
	"sort"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
)

type NetworthPet struct {
	PetData skycrypttypes.Pet `json:"petData"`

	NonCosmetic bool   `json:"nonCosmetic"`
	Skin        string `json:"string"`
	BasePetId   string `json:"basePetId"`
	PetId       string `json:"petId"`
	Level       Level  `json:"level"`
	PetName     string `json:"petName"`

	Prices Prices `json:"prices"`

	Calculation []CalculationData `json:"calculation"`
	BasePrice   float64           `json:"basePrice"`
	Price       float64           `json:"price"`
}

type Level struct {
	Level           int     `json:"level"`
	ExperienceToMax int     `json:"xpMax"`
	Experience      float64 `json:"xp"`
}

func (item *NetworthPet) getTier() int {
	return slices.Index(constants.RARITIES, item.PetData.Rarity)
}

func (item *NetworthPet) getTierName() string {
	return item.PetData.Rarity
}

func (item *NetworthPet) getTierBoostedTier() int {
	if item.PetData.HeldItem == "PET_ITEM_TIER_BOOST" {
		return item.getTier() + 1
	}

	return item.getTier()
}

func (item *NetworthPet) getTierBoostedTierName() string {
	return constants.RARITIES[item.getTierBoostedTier()]
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
	petLevelPrices := item.getPetLevelPrices()

	basePrice := petLevelPrices["LVL_200"]
	if basePrice == 0 {
		basePrice = petLevelPrices["LVL_100"]
	}

	item.BasePrice = basePrice

	if item.Level.Level < 100 && item.Level.ExperienceToMax > 0 {
		baseFormula := (petLevelPrices["LVL_100"] - petLevelPrices["LVL_1"]) / float64(item.Level.ExperienceToMax)

		if baseFormula > 0 {
			item.BasePrice = baseFormula*item.Level.Experience + petLevelPrices["LVL_1"]

		}
	}

	if item.Level.Level > 100 && item.Level.Level < 200 {
		level := item.Level.Level - 100
		if level > 1 {
			baseFormula := (petLevelPrices["LVL_200"] - petLevelPrices["LVL_100"]) / 100
			if baseFormula > 0 {
				item.BasePrice = baseFormula*float64(level) + petLevelPrices["LVL_100"]
			}
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

	petOffset := constants.RARITY_OFFSET_PETS[item.getTierBoostedTierName()]
	petLevels := constants.PET_LEVELS[petOffset : petOffset+maxPetLevel-1]

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
		Level:           level,
		ExperienceToMax: int(sumOfExp),
		Experience:      item.PetData.Experience,
	}
}

func NewSkyBlockPetCalculator(item *skycrypttypes.Pet, prices Prices, options NetworthOptions) *NetworthPet {
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
