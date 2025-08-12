package models

type SkyblockPet struct {
	Type       string  `json:"type"`
	Experience float64 `json:"exp"`
	Active     bool    `json:"active,omitempty"`
	Rarity     string  `json:"tier"`
	HeldItem   string  `json:"heldItem"`
	CandyUsed  int     `json:"candyUsed"`
	Skin       string  `json:"skin"`
}
