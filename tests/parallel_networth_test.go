package tests

import (
	"testing"

	skycrypttypes "github.com/SkyCryptWebsite/SkyCrypt-Types"
	skyhelpernetworthgo "github.com/SkyCryptWebsite/SkyHelper-Networth-Go"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/lib"
)

func TestNetworthAndNonCosmeticMatchSequential(t *testing.T) {
	member := loadTestMember(t)
	items, err := lib.ParseItems(&member, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Fixed prices for every item and gem in the fixture, so the run is offline and repeatable.
	prices := map[string]float64{}
	for _, group := range [][]*skycrypttypes.Item{items.Inventory, items.Enderchest, items.Armor, items.Equipment, items.Storage, items.Accessories} {
		for _, item := range group {
			if item == nil || item.Tag == nil || item.Tag.ExtraAttributes == nil {
				continue
			}
			prices[item.Tag.ExtraAttributes.Id] = 1_000_000
		}
	}
	for _, tier := range []string{"ROUGH", "FLAWED", "FINE", "FLAWLESS", "PERFECT"} {
		for _, gem := range []string{"RUBY", "AMBER", "SAPPHIRE", "JADE", "AMETHYST", "TOPAZ", "JASPER", "OPAL", "AQUAMARINE", "CITRINE", "ONYX", "PERIDOT"} {
			prices[tier+"_"+gem+"_GEM"] = 50_000
		}
	}

	calculator := &skyhelpernetworthgo.ProfileNetworthCalculator{Items: items, Prices: prices}
	options := skyhelpernetworthgo.NetworthOptions{OnlyNetworth: true}
	wantNormal := calculator.GetNetworth(options)
	wantNonCosmetic := calculator.GetNonCosmeticNetworth(options)
	if wantNormal.Networth == 0 {
		t.Fatal("fixture priced at zero; the test would prove nothing")
	}

	for range 20 {
		normal, nonCosmetic := calculator.GetNetworthAndNonCosmetic(options)
		if !closeEnough(normal.Networth, wantNormal.Networth) || !closeEnough(nonCosmetic.Networth, wantNonCosmetic.Networth) {
			t.Fatalf("parallel %v/%v, sequential %v/%v", normal.Networth, nonCosmetic.Networth, wantNormal.Networth, wantNonCosmetic.Networth)
		}
	}
}

// closeEnough allows for summation order, which follows map iteration.
func closeEnough(a, b float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff < 0.01
}
