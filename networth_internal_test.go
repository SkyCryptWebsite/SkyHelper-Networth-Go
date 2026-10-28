package skyhelpernetworthgo

import (
	"testing"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

// Combined Safari Belts each list the upgrades applied to them; across a profile, each upgrade counts once.
func TestSafariBeltModifiersCountOnce(t *testing.T) {
	counted := map[string]map[string]bool{}
	belt := func() models.NetworthItemResult {
		return models.NetworthItemResult{
			Id:    "SAFARI_BELT",
			Price: 1000 + 300 + 200,
			Calculation: []models.CalculationData{
				{Id: "UPGRADE_A", Type: "UPGRADE", Price: 300},
				{Id: "UPGRADE_B", Type: "UPGRADE", Price: 200, Soulbound: true},
			},
			SoulboundPortion: 200,
		}
	}

	first, second := belt(), belt()
	dropCountedModifiers(&first, counted)
	dropCountedModifiers(&second, counted)

	if first.Price != 1500 || len(first.Calculation) != 2 {
		t.Fatalf("first belt %.0f with %d modifiers, want 1500 with 2", first.Price, len(first.Calculation))
	}
	if second.Price != 1000 || len(second.Calculation) != 0 || second.SoulboundPortion != 0 {
		t.Fatalf("second belt %.0f (soulbound %.0f) with %d modifiers, want only its base price", second.Price, second.SoulboundPortion, len(second.Calculation))
	}
}
