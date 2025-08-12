package models

type SkyblockProfile struct {
	Members map[string]SkyblockProfileMember `json:"members"`
	Banking skyblockBanking                  `json:"banking"`
}

type SkyblockProfileMember struct {
	Inventory       skyblockInventory       `json:"inventory"`
	SharedInventory skyblockSharedInventory `json:"shared_inventory"`
	Pets            skyblockPets            `json:"pets_data"`
	Currencies      skyblockCurrencies      `json:"currencies"`
	SackCounts      map[string]int          `json:"sack_counts"`
	Profile         memberProfileData       `json:"profile"`
}

type skyblockInventory struct {
	Armor            nbtData                      `json:"inv_armor"`
	BackpackContents map[string]nbtData           `json:"backpack_contents"`
	BackpackIcons    map[string]nbtData           `json:"backpack_icons"`
	BagContents      skyblockInventoryBagContents `json:"bag_contents"`
	Enderchest       nbtData                      `json:"ender_chest_contents"`
	Equipment        nbtData                      `json:"equipment_contents"`
	Inventory        nbtData                      `json:"inv_contents"`
	PersonalVault    nbtData                      `json:"personal_vault_contents"`
	SackCounts       map[string]int               `json:"sacks_counts"`
	Wardrobe         nbtData                      `json:"wardrobe_contents"`
}

type skyblockSharedInventory struct {
	CandyInventory        nbtData `json:"candy_inventory_contents"`
	CarnivalMaskInventory nbtData `json:"carnival_mask_inventory_contents"`
}

type skyblockInventoryBagContents struct {
	Accessories nbtData `json:"talisman_bag"`
	FishingBag  nbtData `json:"fishing_bag"`
	PotionBag   nbtData `json:"potion_bag"`
	SacksBag    nbtData `json:"sacks_bag"`
	Quiver      nbtData `json:"quiver"`
}

type nbtData struct {
	Data string `json:"data"`
	Type int    `json:"type"`
}

type skyblockBanking struct {
	Balance float64 `json:"balance"`
}

type skyblockPets struct {
	Pets []SkyblockPet `json:"pets"`
}

type skyblockCurrencies struct {
	Essence map[string]essence `json:"essence"`
	Coins   float64            `json:"coin_purse"`
}

type essence struct {
	Current int `json:"current"`
}

type memberProfileData struct {
	PersonalBank float64 `json:"bank_account"`
}
