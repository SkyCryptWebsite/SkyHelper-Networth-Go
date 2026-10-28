package tests

import (
	"testing"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
)

// Entries SkyHelper-networth added or fixed since the last sync.
func TestUpstreamTableEntries(t *testing.T) {
	for reforge, stone := range map[string]string{
		"warped": "ENDSTONE_GEODE", "jerry_stone": "JERRY_STONE", "blood_shot": "SHRIVELED_CORNEA",
		"lunar": "MOONSTONE", "sunny": "SUNSTONE", "erudite": "DAEDALUS_NOTES", "mantid": "MANTID_CLAW",
		"overpriced": "OVERPRICED_DRINK", "deep_fried": "HASHBROWN", "geometric": "GEOMETRIC_ODDITY",
		"sticky": "EXTREMELY_MILD_ADHESIVE", "majestic": "MORNING_DEW", "marshy": "MARSHROOM", "thorny": "BLOOMING_THORNS",
	} {
		if got := constants.REFORGES[reforge]; got != stone {
			t.Errorf("reforge %s -> %q, want %s", reforge, got, stone)
		}
	}
	for _, stale := range []string{"hyper", "jerry", "bloodshot"} {
		if _, ok := constants.REFORGES[stale]; ok {
			t.Errorf("reforge %s is a wrong id upstream removed", stale)
		}
	}
	for enchantment, item := range map[string]string{"FOREST_PLEDGE": "WATER_HYACINTH", "KARMA": "FORESTS_FAVOR", "STEALTH": "DISTANT_ECHO"} {
		upgrades := constants.ENCHANTMENT_UPGRADES[enchantment]
		if len(upgrades) != 1 || upgrades[0].UpgradeItem != item || upgrades[0].Tier != 6 {
			t.Errorf("%s upgrade = %+v, want %s at tier 6", enchantment, upgrades, item)
		}
	}
	if constants.PICKONIMBUS_DURABILITY != 2000 {
		t.Errorf("Pickonimbus durability %d, want 2000", constants.PICKONIMBUS_DURABILITY)
	}
	if constants.SPECIAL_ENCHANTMENT_NAMES["magmarizer"] != "Pyroclasm" || constants.SPECIAL_ENCHANTMENT_NAMES["triple_strike"] != "Triple-Strike" {
		t.Error("missing special enchantment names")
	}
}
