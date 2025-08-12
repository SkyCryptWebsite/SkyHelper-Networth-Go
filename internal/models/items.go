package models

type DecodedItem struct {
	Count  int         `nbt:"Count" json:"count,omitempty"`
	Damage int         `nbt:"Damage" json:"damage,omitempty"`
	Tag    *DecodedTag `nbt:"tag" json:"tag,omitempty"`
}

type DecodedTag struct {
	ExtraAttributes *ExtraAttributes `nbt:"ExtraAttributes" json:"extra_attributes,omitempty"`
	Display         *DecodedDisplay  `nbt:"display" json:"display,omitempty"`
}

type ExtraAttributes struct {
	Id                    string         `nbt:"id" json:"id,omitempty"`
	Uuid                  string         `nbt:"uuid" json:"uuid,omitempty"`
	AbilityScroll         []string       `nbt:"ability_scroll" json:"ability_scroll,omitempty"`
	AdditionalCoins       int64          `nbt:"additional_coins" json:"additional_coins,omitempty"`
	ArtOfPeaceApplied     int            `nbt:"artOfPeaceApplied" json:"artOfPeaceApplied,omitempty"`
	ArtOfWarCount         int            `nbt:"art_of_war_count" json:"art_of_war_count,omitempty"`
	Boosters              []string       `nbt:"boosters" json:"boosters,omitempty"`
	DivanPowderCoating    int            `nbt:"divan_powder_coating" json:"divan_powder_coating,omitempty"`
	DungeonItemLevel      any            `nbt:"dungeon_item_level" json:"dungeon_item_level,omitempty"`
	DyeItem               string         `nbt:"dye_item" json:"dye_item,omitempty"`
	Enchantments          map[string]int `nbt:"enchantments" json:"enchantments,omitempty"`
	Ethermerge            int            `nbt:"ethermerge" json:"ethermerge,omitempty"`
	FarmingForDummies     int            `nbt:"farming_for_dummies_count" json:"farming_for_dummies_count,omitempty"`
	Gems                  map[string]any `nbt:"gems" json:"gems,omitempty"`
	HotPotatoCount        int            `nbt:"hot_potato_count" json:"hot_potato_count,omitempty"`
	JalapenoCount         int            `nbt:"jalapeno_count" json:"jalapeno_count,omitempty"`
	ManaDisintegrator     int            `nbt:"mana_disintegrator_count" json:"mana_disintegrator_count,omitempty"`
	Modifier              string         `nbt:"modifier" json:"modifier,omitempty"`
	NewYearCakeBagData    []byte         `nbt:"new_year_cake_bag_data" json:"new_year_cake_bag_data,omitempty"`
	NewYearCakeBagYears   []int          `nbt:"new_year_cake_bag_years" json:"new_year_cake_bag_years,omitempty"`
	NewYearsCake          int            `nbt:"new_years_cake" json:"new_years_cake,omitempty"`
	PartyHatEmoji         string         `nbt:"party_hat_emoji" json:"party_hat_emoji,omitempty"`
	PartyHatColor         string         `nbt:"party_hat_color" json:"party_hat_color,omitempty"`
	PetInfo               string         `nbt:"petInfo" json:"petInfo,omitempty"`
	PickonimbusDurability int            `nbt:"pickonimbus_durability" json:"pickonimbus_durability,omitempty"`
	Polarvoid             int            `nbt:"polarvoid" json:"polarvoid,omitempty"`
	Price                 int64          `nbt:"price" json:"price,omitempty"`
	PowerAbilityScroll    string         `nbt:"power_ability_scroll" json:"power_ability_scroll,omitempty"`
	Recombobulated        int            `nbt:"rarity_upgrades" json:"rarity_upgrades,omitempty"`
	Runes                 map[string]int `nbt:"runes" json:"runes,omitempty"`
	SackPss               int            `nbt:"sack_pss" json:"sack_pss,omitempty"`
	Skin                  string         `nbt:"skin" json:"skin,omitempty"`
	TalismanEnrichment    string         `nbt:"talisman_enrichment" json:"talisman_enrichment,omitempty"`
	// Timestamp             int64                  `nbt:"timestamp" json:"timestamp,omitempty"`
	ThunderCharge          int            `nbt:"thunder_charge" json:"thunder_charge,omitempty"`
	TunedTransmission      int            `nbt:"tuned_transmission" json:"tuned_transmission,omitempty"`
	UpgradeLevel           any            `nbt:"upgrade_level" json:"upgrade_level,omitempty"`
	WinningBid             int64          `nbt:"winning_bid" json:"winning_bid,omitempty"`
	WoodSingularityCount   int            `nbt:"wood_singularity_count" json:"wood_singularity_count,omitempty"`
	DrillPartUpgradeModule string         `nbt:"drill_part_upgrade_module" json:"drill_part_upgrade_module,omitempty"`
	DrillPartFuelTank      string         `nbt:"drill_part_fuel_tank" json:"drill_part_fuel_tank,omitempty"`
	DrillPartEngine        string         `nbt:"drill_part_engine" json:"drill_part_engine,omitempty"`
	Line                   RodPart        `nbt:"line" json:"line,omitempty"`
	Hook                   RodPart        `nbt:"hook" json:"hook,omitempty"`
	Sinker                 RodPart        `nbt:"sinker" json:"sinker,omitempty"`
	Soulbound              bool           `nbt:"donated_museum" json:"donated_museum,omitempty"`
	Attributes             map[string]int `nbt:"attributes" json:"attributes,omitempty"`
	Edition                int            `nbt:"edition" json:"edition,omitempty"`
	Shiny                  bool           `nbt:"is_shiny" json:"is_shiny,omitempty"`
	ItemTier               int            `nbt:"item_tier" json:"item_tier,omitempty"`

	// TODO: confirm these 2 types
	Auction int64 `nbt:"auction" json:"auction,omitempty"`
	Bid     int64 `nbt:"bid" json:"bid,omitempty"`
}

type RodPart struct {
	Part      string `nbt:"part" json:"part,omitempty"`
	Soulbound bool   `nbt:"donated_museum" json:"donated_museum,omitempty"`
}

type DecodedNewYearCakeBagData struct {
	Items []DecodedItem `nbt:"i" json:"items,omitempty"`
}

type DecodedDisplay struct {
	Name  string   `nbt:"Name" json:"name,omitempty"`
	Lore  []string `nbt:"Lore" json:"lore,omitempty"`
	Color int      `nbt:"color" json:"color,omitempty"`
}

type DecodedInventory struct {
	Items []DecodedItem `nbt:"i" json:"items,omitempty"`
}
