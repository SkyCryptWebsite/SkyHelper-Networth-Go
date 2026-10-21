package tests

import (
	"encoding/json"
	"os"
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/lib"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func loadTestMember(t *testing.T) skycrypttypes.Member {
	t.Helper()
	file, err := os.Open("./test_profile.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = file.Close() }()

	// Only the inventory is decoded: the rest of the fixture no longer matches SkyCrypt-Types.
	var profile struct {
		Members map[string]struct {
			Inventory *skycrypttypes.Inventory `json:"inventory"`
		} `json:"members"`
	}
	if err := json.NewDecoder(file).Decode(&profile); err != nil {
		t.Fatal(err)
	}

	return skycrypttypes.Member{Inventory: profile.Members["fb3d96498a5b4d5b91b763db14b195ad"].Inventory}
}

// Equipment loadout items used to be dropped whenever the "equipment" key was visited last in map order.
func TestParseItemsKeepsEquipmentLoadoutsInAnyOrder(t *testing.T) {
	member := loadTestMember(t)
	base, err := lib.ParseItems(&member, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(base.Equipment) == 0 {
		t.Fatal("test profile has no equipment items")
	}

	member.Loadout.Equipment.Sets = map[int]skycrypttypes.EquipmentLoadout{
		0: {EquipmentSlot1: member.Inventory.Equipment},
	}

	for range 100 {
		parsed, err := lib.ParseItems(&member, nil)
		if err != nil {
			t.Fatal(err)
		}
		if got, want := len(parsed.Equipment), 2*len(base.Equipment); got != want {
			t.Fatalf("expected %d equipment items, got %d", want, got)
		}
	}
}

func TestParseItemsDoesNotModifyMuseum(t *testing.T) {
	member := loadTestMember(t)
	museum := &skycrypttypes.Museum{}
	if _, err := lib.ParseItems(&member, museum); err != nil {
		t.Fatal(err)
	}
	if museum.Items != nil || museum.Special != nil {
		t.Fatal("ParseItems wrote defaults onto the shared museum")
	}
}

func TestParseItemsSkipsUndecodableMuseumItems(t *testing.T) {
	member := loadTestMember(t)
	museum := &skycrypttypes.Museum{
		Items: &map[string]skycrypttypes.MuseumItem{
			"BROKEN": {Items: skycrypttypes.EncodedItems{Data: "not base64 gzip"}},
		},
	}
	if _, err := lib.ParseItems(&member, museum); err != nil {
		t.Fatal(err)
	}
}

type stubItemProvider struct{ item *models.HypixelItem }

func (p stubItemProvider) GetItem(string) *models.HypixelItem { return p.item }

func TestItemCalculatorHandlesMissingCountAndLore(t *testing.T) {
	item := &skycrypttypes.Item{
		Tag: &skycrypttypes.Tag{
			ExtraAttributes: &skycrypttypes.ExtraAttributes{Id: "HYPERION"},
		},
	}
	prices := map[string]float64{"HYPERION": 100}

	calculator := models.NewSkyBlockItemCaclulator(item, prices, stubItemProvider{&models.HypixelItem{Category: "SWORD"}}, models.NetworthOptions{})
	if calculator.Count != 1 {
		t.Fatalf("expected a missing Count to mean 1, got %d", calculator.Count)
	}
	if calculator.IsCosmetic() {
		t.Fatal("item without lore should not be cosmetic")
	}
	if calculator.BasePrice != 100 {
		t.Fatalf("expected base price 100, got %v", calculator.BasePrice)
	}
}

func TestPetCalculatorSkipsUnknownRarity(t *testing.T) {
	service := calculators.NewCalculatorService()
	prices := map[string]float64{
		"LVL_100_WEIRD_BEE":   1000,
		"PET_ITEM_TIER_BOOST": 5000,
	}

	for _, rarity := range []string{"", "WEIRD"} {
		pet := &skycrypttypes.Pet{Type: "BEE", Rarity: rarity, HeldItem: "PET_ITEM_TIER_BOOST", Experience: 1e9}
		calculator := service.NewSkyBlockPetCalculator(pet, prices, models.NetworthOptions{})
		service.CalculatePet(calculator)
		if got := calculator.GetPrice(); got != 0 {
			t.Fatalf("rarity %q: expected price 0, got %v", rarity, got)
		}
	}
}

func TestPetCalculatorTierBoostOnHighestRarity(t *testing.T) {
	service := calculators.NewCalculatorService()
	pet := &skycrypttypes.Pet{Type: "BEE", Rarity: "ULTIMATE", HeldItem: "PET_ITEM_TIER_BOOST", Experience: 1e9}
	calculator := service.NewSkyBlockPetCalculator(pet, map[string]float64{}, models.NetworthOptions{})
	service.CalculatePet(calculator)
}
