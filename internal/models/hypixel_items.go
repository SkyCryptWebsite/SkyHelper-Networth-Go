package models

type HypixelItemsResponse struct {
	Success     bool          `json:"success"`
	Cause       string        `json:"cause,omitempty"`
	LastUpdated int64         `json:"lastUpdated"`
	Items       []HypixelItem `json:"items"`
}

type HypixelItem struct {
	Material          string          `json:"material"`
	Skin              skin            `json:"skin,omitempty"`
	Name              string          `json:"name"`
	Category          string          `json:"category"`
	Rarity            string          `json:"tier"`
	SkyBlockID        string          `json:"id,omitempty"`
	Damage            int             `json:"damage,omitempty"`
	Origin            string          `json:"origin,omitempty"`
	RiftTransferrable bool            `json:"rift_transferrable,omitempty"`
	UpgradeCosts      [][]UpgradeCost `json:"upgrade_costs,omitempty"`
	GemstoneSlots    []GemstoneSlot  `json:"gemstone_slots,omitempty"`
	Prestige          PrestigeCost    `json:"prestige,omitempty"`
}

type skin struct {
	Value     string `json:"value"`
	Signature string `json:"signature,omitempty"`
}

type UpgradeCost struct {
	Type        string `json:"type"`
	EssenceType string `json:"essence_type,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	ItemId      string `json:"item_id,omitempty"`
}

type GemstoneSlot struct {
	SlotType string         `json:"slot_type"`
	Costs    []GemstoneCost `json:"costs"`
}

type GemstoneCost struct {
	Type   string `json:"type"`
	ItemId string `json:"item_id"`
	Coins  int    `json:"coins"`
	Amount int    `json:"amount"`
}
type PrestigeCost struct {
	ItemId string        `json:"item_id"`
	Costs  []UpgradeCost `json:"costs"`
}
