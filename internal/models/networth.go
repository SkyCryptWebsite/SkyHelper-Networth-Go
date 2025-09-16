package models

import skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"

type Prices map[string]float64

type CalculationData struct {
	Id        string  `json:"id"`
	Type      string  `json:"type"`
	Price     float64 `json:"price"`
	Count     int     `json:"count"`
	Soulbound bool    `json:"soulbound,omitempty"`
	Star      int     `json:"star,omitempty"`
}

type ParsedItems struct {
	Accessories           []*skycrypttypes.Item `json:"accessories"`
	Armor                 []*skycrypttypes.Item `json:"armor"`
	CandyInventory        []*skycrypttypes.Item `json:"candy_inventory"`
	CarnivalMaskInventory []*skycrypttypes.Item `json:"carnival_mask_inventory"`
	Equipment             []*skycrypttypes.Item `json:"equipment"`
	Enderchest            []*skycrypttypes.Item `json:"enderchest"`
	Essence               []*BasicItem          `json:"essence"`
	FishingBag            []*skycrypttypes.Item `json:"fishing_bag"`
	Inventory             []*skycrypttypes.Item `json:"inventory"`
	PersonalVault         []*skycrypttypes.Item `json:"personal_vault"`
	Pets                  []*skycrypttypes.Pet  `json:"pets"`
	PotionBag             []*skycrypttypes.Item `json:"potion_bag"`
	Quiver                []*skycrypttypes.Item `json:"quiver"`
	Sacks                 []*BasicItem          `json:"sacks"`
	SacksBag              []*skycrypttypes.Item `json:"sacks_bag"`
	Storage               []*skycrypttypes.Item `json:"storage"`
	Wardrobe              []*skycrypttypes.Item `json:"wardrobe"`
	Museum                []*skycrypttypes.Item `json:"museum"`
}

type BasicItem struct {
	Id     string `json:"id"`
	Amount int    `json:"amount"`
}

type NetworthOptions struct {
	Prices          Prices
	NonCosmetic     bool
	CachePrices     bool
	OnlyNetworth    bool
	IncludeItemData bool
	SortItems       bool
	StackItems      bool
}

type NetworthResult struct {
	Networth            float64                  `json:"networth"`
	UnsoulboundNetworth float64                  `json:"unsoulboundNetworth"`
	NoInventory         bool                     `json:"noInventory"`
	IsNonCosmetic       bool                     `json:"isNonCosmetic"`
	Purse               float64                  `json:"purse"`
	Bank                float64                  `json:"bank"`
	PersonalBank        float64                  `json:"personalBank"`
	Types               map[string]*NetworthType `json:"types"`
}

type NetworthType struct {
	Total            float64              `json:"total"`
	UnsoulboundTotal float64              `json:"unsoulboundTotal"`
	Items            []NetworthItemResult `json:"items"`
}

type NetworthItemResult struct {
	Name             string            `json:"name"`
	LoreName         string            `json:"loreName"`
	Id               string            `json:"id"`
	CustomId         string            `json:"customId"`
	Price            float64           `json:"price"`
	SoulboundPortion float64           `json:"soulboundPortion"`
	BasePrice        float64           `json:"basePrice"`
	Calculation      []CalculationData `json:"calculation"`
	Count            int               `json:"count"`
	Soulbound        bool              `json:"soulbound"`
	Cosmetic         bool              `json:"cosmetic"`
	ItemData         interface{}       `json:"itemData,omitempty"`
}

type CategoryInfo struct {
	Items interface{}
	Type  string // "decoded", "basic", "pets"
}
