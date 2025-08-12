package models

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
	Accessories           []*DecodedItem `json:"accessories"`
	Armor                 []*DecodedItem `json:"armor"`
	CandyInventory        []*DecodedItem `json:"candy_inventory"`
	CarnivalMaskInventory []*DecodedItem `json:"carnival_mask_inventory"`
	Equipment             []*DecodedItem `json:"equipment"`
	Enderchest            []*DecodedItem `json:"enderchest"`
	Essence               []*BasicItem   `json:"essence"`
	FishingBag            []*DecodedItem `json:"fishing_bag"`
	Inventory             []*DecodedItem `json:"inventory"`
	PersonalVault         []*DecodedItem `json:"personal_vault"`
	Pets                  []*SkyblockPet `json:"pets"`
	PotionBag             []*DecodedItem `json:"potion_bag"`
	Quiver                []*DecodedItem `json:"quiver"`
	Sacks                 []*BasicItem   `json:"sacks"`
	SacksBag              []*DecodedItem `json:"sacks_bag"`
	Storage               []*DecodedItem `json:"storage"`
	Wardrobe              []*DecodedItem `json:"wardrobe"`
	Museum                []*DecodedItem `json:"museum"`
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
