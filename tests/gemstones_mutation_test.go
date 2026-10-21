package tests

import (
	"testing"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

// Normal and non-cosmetic networth price the same decoded items one after the other, and the backend reads the
// gems again afterwards to show them, so pricing must leave the item's gems alone.
func TestGemstonesHandlerDoesNotConsumeGems(t *testing.T) {
	gems := map[string]any{
		"JADE_0":         "FINE",
		"COMBAT_0":       "PERFECT",
		"COMBAT_0_gem":   "SAPPHIRE",
		"unlocked_slots": []any{"JADE_0", "COMBAT_0"},
	}
	prices := map[string]float64{"FINE_JADE_GEM": 100, "PERFECT_SAPPHIRE_GEM": 1000}
	newItem := func() *models.NetworthItem {
		return &models.NetworthItem{
			ItemId:          "DIVAN_DRILL",
			ExtraAttributes: &skycrypttypes.ExtraAttributes{Id: "DIVAN_DRILL", Gems: gems},
			SkyblockItem: &models.HypixelItem{GemstoneSlots: []models.GemstoneSlot{
				{SlotType: "JADE"},
				{SlotType: "COMBAT"},
			}},
			Prices: prices,
		}
	}

	first := newItem()
	handlers.GemstonesHandler{}.Calculate(first, prices)
	second := newItem()
	handlers.GemstonesHandler{}.Calculate(second, prices)

	if first.Price == 0 || first.Price != second.Price {
		t.Fatalf("pricing the same gems twice gave %v then %v", first.Price, second.Price)
	}
	if len(gems) != 4 {
		t.Fatalf("pricing removed gems from the item: %v", gems)
	}
}
