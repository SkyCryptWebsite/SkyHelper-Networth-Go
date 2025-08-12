package constants

// BlockedEnchantments maps item IDs to lists of banned enchantments.
var BLOCKED_ENCHANTMENTS = map[string][]string{
	"BONE_BOOMERANG":         {"OVERLOAD", "POWER", "ULTIMATE_SOUL_EATER"},
	"DEATH_BOW":              {"OVERLOAD", "POWER", "ULTIMATE_SOUL_EATER"},
	"GARDENING_AXE":          {"REPLENISH"},
	"GARDENING_HOE":          {"REPLENISH"},
	"ADVANCED_GARDENING_AXE": {"REPLENISH"},
	"ADVANCED_GARDENING_HOE": {"REPLENISH"},
}

// IgnoredEnchantments maps enchantment names to ignored PET_LEVELS.
var IGNORED_ENCHANTMENTS = map[string]int{
	"SCAVENGER": 5,
}

// StackingEnchantments is a list of enchantments that can stack.
var STACKING_ENCHANTMENTS = []string{
	"EXPERTISE", "COMPACT", "ABSORB", "CULTIVATING", "CHAMPION", "HECATOMB", "TOXOPHILITE",
}

// IgnoreSilex is a list of items where Silex should be ignored.
var IGNORE_SILEX = []string{
	"PROMISING_SPADE", "PROMISING_AXE",
}

// MasterStars is a list of star upgrade item IDs.
var MASTER_STARS = []string{
	"FIRST_MASTER_STAR", "SECOND_MASTER_STAR", "THIRD_MASTER_STAR", "FOURTH_MASTER_STAR", "FIFTH_MASTER_STAR",
}

// AllowedRecombobulatedCategories is a list of item categories allowed to be recombobulated.
var AllowedRecombobulatedCategories = []string{
	"ACCESSORY", "NECKLACE", "GLOVES", "BRACELET", "BELT", "CLOAK", "VACUUM",
}

// AllowedRecombobulatedIDs is a list of specific item IDs that are allowed to be recombobulated.
var ALLOWED_RECOMBOBULATED_IDS = []string{
	"DIVAN_HELMET", "DIVAN_CHESTPLATE", "DIVAN_LEGGINGS", "DIVAN_BOOTS",
	"FERMENTO_HELMET", "FERMENTO_CHESTPLATE", "FERMENTO_LEGGINGS", "FERMENTO_BOOTS",
	"SHADOW_ASSASSIN_CLOAK", "STARRED_SHADOW_ASSASSIN_CLOAK",
}

// Enrichments is a list of enrichment item IDs.
var ENRICHMENTS = []string{
	"TALISMAN_ENRICHMENT_CRITICAL_CHANCE", "TALISMAN_ENRICHMENT_CRITICAL_DAMAGE",
	"TALISMAN_ENRICHMENT_DEFENSE", "TALISMAN_ENRICHMENT_HEALTH", "TALISMAN_ENRICHMENT_INTELLIGENCE",
	"TALISMAN_ENRICHMENT_MAGIC_FIND", "TALISMAN_ENRICHMENT_WALK_SPEED", "TALISMAN_ENRICHMENT_STRENGTH",
	"TALISMAN_ENRICHMENT_ATTACK_SPEED", "TALISMAN_ENRICHMENT_FEROCITY", "TALISMAN_ENRICHMENT_SEA_CREATURE_CHANCE",
}

// SpecialEnchantmentNames maps internal enchantment names to user-friendly names.
var SPECIAL_ENCHANTMENT_NAMES = map[string]string{
	"aiming":               "Dragon Tracer",
	"counter_strike":       "Counter-Strike",
	"pristine":             "Prismatic",
	"turbo_cacti":          "Turbo-Cacti",
	"turbo_cane":           "Turbo-Cane",
	"turbo_carrot":         "Turbo-Carrot",
	"turbo_cocoa":          "Turbo-Cocoa",
	"turbo_melon":          "Turbo-Melon",
	"turbo_mushrooms":      "Turbo-Mushrooms",
	"turbo_potato":         "Turbo-Potato",
	"turbo_pumpkin":        "Turbo-Pumpkin",
	"turbo_warts":          "Turbo-Warts",
	"turbo_wheat":          "Turbo-Wheat",
	"ultimate_reiterate":   "Ultimate Duplex",
	"ultimate_bobbin_time": "Ultimate Bobbin' Time",
}

// GemstoneSlots lists the types of available gemstone slots.
var GEMSTONE_SLOTS = []string{
	"COMBAT", "OFFENSIVE", "DEFENSIVE", "MINING", "UNIVERSAL", "CHISEL",
}

// NonCosmeticItems is a set of item IDs that are not considered cosmetic.
var NON_COSMETIC_ITEMS = []string{
	"ANCIENT_ELEVATOR",
	"BEDROCK",
	"CREATIVE_MIND",
	"DCTR_SPACE_HELM",
	"DEAD_BUSH_OF_LOVE",
	"DUECES_BUILDER_CLAY",
	"GAME_BREAKER",
	"POTATO_BASKET",
}

var PICKONIMBUS_DURABILITY = 5000

var MAX_THUNDER_CHARGE = 5000000

var ROD_PART_TYPES = []string{"line", "hook", "sinker"}

var ENCHANTMENT_UPGRADES = map[string]struct {
	UpgradeItem string
	Tier        int
}{
	"SCAVENGER":       {"GOLDEN_BOUNTY", 6},
	"PESTERMINATOR":   {"PESTHUNTING_GUIDE", 6},
	"LUCK_OF_THE_SEA": {"GOLD_BOTTLE_CAP", 7},
	"PISCARY":         {"TROUBLED_BUBBLE", 7},
	"FRAIL":           {"SEVERED_PINCER", 7},
	"SPIKED_HOOK":     {"OCTOPUS_TENDRIL", 7},
	"CHARM":           {"CHAIN_END_TIMES", 6},
}

type MidasWeapon struct {
	MaxBid int64  `json:"maxBid"`
	Type   string `json:"type"`
}

var MIDAS_SWORDS = map[string]MidasWeapon{
	"MIDAS_SWORD": {
		MaxBid: 50_000_000,
		Type:   "MIDAS_SWORD_50M",
	},
	"STARRED_MIDAS_SWORD": {
		MaxBid: 250_000_000,
		Type:   "STARRED_MIDAS_SWORD_250M",
	},
	"MIDAS_STAFF": {
		MaxBid: 100_000_000,
		Type:   "MIDAS_STAFF_100M",
	},
	"STARRED_MIDAS_STAFF": {
		MaxBid: 500_000_000,
		Type:   "STARRED_MIDAS_STAFF_500M",
	},
}
