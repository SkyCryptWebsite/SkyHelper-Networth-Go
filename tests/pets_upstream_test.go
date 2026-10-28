package tests

import (
	"math"
	"strings"
	"testing"

	skycrypttypes "github.com/SkyCryptWebsite/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func newPet(pet skycrypttypes.Pet, prices map[string]float64) *models.NetworthPet {
	return models.NewSkyBlockPetCalculator(&pet, prices, models.NetworthOptions{})
}

// Ported from SkyHelper-networth's "Phoenix rarity conversion" tests.
func TestSpecialPhoenixLevelsAndPrice(t *testing.T) {
	for _, c := range []struct {
		tier, formerTier string
		maxXP            float64
		name             string
	}{
		{"SPECIAL", "EPIC", 18608500, "Special Phoenix"},
		{"VERY_SPECIAL", "LEGENDARY", 25353230, "Very Special Phoenix"},
	} {
		for _, exp := range []float64{0, c.maxXP / 2, c.maxXP, c.maxXP + 1000} {
			prices := map[string]float64{"LVL_1_" + c.tier + "_PHOENIX": 1000, "LVL_100_" + c.tier + "_PHOENIX": 100000}
			pet := newPet(skycrypttypes.Pet{Type: "PHOENIX", Rarity: c.tier, Experience: exp}, prices)
			former := newPet(skycrypttypes.Pet{Type: "PHOENIX", Rarity: c.formerTier, Experience: exp}, nil)

			if pet.Level != former.Level {
				t.Errorf("%s at %.0f XP: level %+v, want the %s curve %+v", c.tier, exp, pet.Level, c.formerTier, former.Level)
			}
			if float64(pet.Level.ExperienceToMax) != c.maxXP || float64(pet.Level.ExperienceTo100) != c.maxXP {
				t.Errorf("%s: xp to max %d, to 100 %d, want %.0f", c.tier, pet.Level.ExperienceToMax, pet.Level.ExperienceTo100, c.maxXP)
			}
			if !strings.Contains(pet.PetName, c.name) {
				t.Errorf("name %q, want it to contain %q", pet.PetName, c.name)
			}
			want := 100000.0
			if exp < c.maxXP {
				want = 1000 + 99000*exp/c.maxXP
			}
			if math.Abs(pet.BasePrice-want) > 0.01 {
				t.Errorf("%s at %.0f XP: base price %.2f, want %.2f", c.tier, exp, pet.BasePrice, want)
			}
			if id := pet.GetPetId(); id != "LVL_100_"+c.tier+"_PHOENIX" {
				t.Errorf("pet id %s", id)
			}
		}
	}
}

func TestPhoenixIgnoresTierBoost(t *testing.T) {
	for _, tier := range []string{"SPECIAL", "VERY_SPECIAL", "LEGENDARY"} {
		original := newPet(skycrypttypes.Pet{Type: "PHOENIX", Rarity: tier, Experience: 100000}, nil)
		boosted := newPet(skycrypttypes.Pet{Type: "PHOENIX", Rarity: tier, Experience: 100000, HeldItem: "PET_ITEM_TIER_BOOST"}, nil)
		if boosted.Level != original.Level || boosted.PetName != original.PetName {
			t.Errorf("%s Phoenix with a Tier Boost: %q %+v, want %q %+v", tier, boosted.PetName, boosted.Level, original.PetName, original.Level)
		}
	}
	// Other pets still get boosted.
	boosted := newPet(skycrypttypes.Pet{Type: "ENDERMAN", Rarity: "EPIC", Experience: 100000, HeldItem: "PET_ITEM_TIER_BOOST"}, nil)
	if !strings.Contains(boosted.PetName, "Legendary") {
		t.Errorf("a boosted epic Enderman should read Legendary, got %q", boosted.PetName)
	}
}

func TestLevel100PetUsesLevel100Price(t *testing.T) {
	prices := map[string]float64{
		"LVL_1_LEGENDARY_GOLDEN_DRAGON":   1_000_000,
		"LVL_100_LEGENDARY_GOLDEN_DRAGON": 700_000_000,
		"LVL_200_LEGENDARY_GOLDEN_DRAGON": 1_200_000_000,
	}
	level100 := newPet(skycrypttypes.Pet{Type: "GOLDEN_DRAGON", Rarity: "LEGENDARY", Experience: 25353230}, prices)
	if level100.Level.Level != 100 || level100.BasePrice != 700_000_000 {
		t.Errorf("level %d priced %.0f, want level 100 at the level 100 price", level100.Level.Level, level100.BasePrice)
	}

	// Below 100 the price scales by the XP towards level 100, not towards 200.
	half := newPet(skycrypttypes.Pet{Type: "GOLDEN_DRAGON", Rarity: "LEGENDARY", Experience: 25353230 / 2}, prices)
	want := 1_000_000 + (700_000_000-1_000_000)*0.5
	if math.Abs(half.BasePrice-want) > 1 {
		t.Errorf("half way to 100: %.0f, want %.0f", half.BasePrice, want)
	}

	// Between 100 and 200 it scales by level.
	level150 := newPet(skycrypttypes.Pet{Type: "GOLDEN_DRAGON", Rarity: "LEGENDARY", Experience: 25353230 + 50*1886700}, prices)
	if level150.Level.Level != 150 || math.Abs(level150.BasePrice-950_000_000) > 1 {
		t.Errorf("level %d priced %.0f, want level 150 at 950000000", level150.Level.Level, level150.BasePrice)
	}
}
