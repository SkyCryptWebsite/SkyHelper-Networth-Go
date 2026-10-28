# SkyHelper-Networth-Go

[![discord](https://img.shields.io/discord/720018827433345138?logo=discord)](https://discord.com/invite/fd4Be4W)
[![license](https://img.shields.io/badge/license-MIT-green)](LICENSE.MD)
[![Go Reference](https://pkg.go.dev/badge/github.com/SkyCryptWebsite/SkyHelper-Networth-Go.svg)](https://pkg.go.dev/github.com/SkyCryptWebsite/SkyHelper-Networth-Go)

A Go port of [SkyHelper](https://skyhelper.altpapier.dev/)'s networth calculation. Give it a SkyBlock profile member
from the [Hypixel API](https://api.hypixel.net/) and it decodes every inventory, prices each item (with its
enchantments, gemstones, reforges, upgrades, pets and so on) and returns the networth per category.

> [!NOTE]
> This module is built for [SkyCrypt](https://github.com/SkyCryptWebsite/SkyCrypt-Backend). It covers the networth
> calculation itself, normal and non-cosmetic, but not the Node module's `UpdateManager` or `NetworthManager`.

## Installation

```bash
go get github.com/SkyCryptWebsite/SkyHelper-Networth-Go
```

Profile data uses the types from [SkyCrypt-Types](https://github.com/SkyCryptWebsite/SkyCrypt-Types), so decode the
Hypixel responses with those:

```bash
go get github.com/SkyCryptWebsite/SkyCrypt-Types
```

## Quick start

Fetch the player's profiles (`/v2/skyblock/profiles?uuid=...`), pick a profile and the member, and calculate:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	skycrypttypes "github.com/SkyCryptWebsite/SkyCrypt-Types"
	networth "github.com/SkyCryptWebsite/SkyHelper-Networth-Go"
)

func main() {
	// The body of https://api.hypixel.net/v2/skyblock/profiles?uuid=<uuid>
	body, err := os.ReadFile("profiles.json")
	if err != nil {
		log.Fatal(err)
	}
	var response struct {
		Profiles []skycrypttypes.Profile `json:"profiles"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}

	const uuid = "4855c53ee4fb4100997600a92fc50984" // no dashes, as Hypixel keys members
	for _, profile := range response.Profiles {
		if !profile.Selected {
			continue
		}
		member := profile.Members[uuid]

		var bank float64
		if profile.Banking != nil && profile.Banking.Balance != nil {
			bank = *profile.Banking.Balance
		}

		// Pass the member's museum (from /v2/skyblock/museum) instead of nil to count donated items too.
		calculator, err := networth.NewProfileNetworthCalculator(&member, nil, bank)
		if err != nil {
			log.Fatal(err)
		}

		result := calculator.GetNetworth()
		fmt.Printf("%s: %.0f coins (%.0f without soulbound items)\n",
			profile.CuteName, result.Networth, result.UnsoulboundNetworth)
		for category, total := range result.Types {
			fmt.Printf("  %-24s %.0f\n", category, total.Total)
		}
	}
}
```

`NewProfileNetworthCalculator` decodes the inventories once; you can then calculate as often as you like with
different options. The first call downloads item prices and the Hypixel item list (see [Prices](#prices)).

## Normal and non-cosmetic networth

Non-cosmetic networth leaves out cosmetic items (skins, dyes, runes and the like):

```go
normal := calculator.GetNetworth()
nonCosmetic := calculator.GetNonCosmeticNetworth()

// Or both at once. On more than one CPU the two run side by side.
normal, nonCosmetic = calculator.GetNetworthAndNonCosmetic(networth.NetworthOptions{})
fmt.Printf("%.0f (%.0f without cosmetics)\n", normal.Networth, nonCosmetic.Networth)
```

## Options

Every calculation takes an optional `networth.NetworthOptions`:

| Field | Default | What it does |
|---|---|---|
| `Prices` | fetched | Use this price list (item ID → coins) instead of the cached SkyHelper prices. |
| `NonCosmetic` | `false` | Leave cosmetic items out. `GetNonCosmeticNetworth` sets it for you. |
| `OnlyNetworth` | `false` | Only compute the totals; `Types[...].Items` stay empty. Much less allocation when you only show the number. |
| `IncludeItemData` | `false` | Attach the decoded item (`ItemData`) or pet (`PetData`) to every result item. |
| `KeepInvalidItems` | `false` | Keep items that couldn't be priced (no SkyBlock ID) in the item lists, with a price of 0. |
| `SortItems` | `false` | Sort each category's items by price, highest first. |
| `StackItems` | `false` | Merge identical items into one entry with a `Count`. |
| `CachePrices` | `false` | Currently unused: prices are always cached in process (see [Prices](#prices)). |

```go
result := calculator.GetNetworth(networth.NetworthOptions{
	OnlyNetworth: true, // just the totals
})
fmt.Println(result.Networth)

detailed := calculator.GetNetworth(networth.NetworthOptions{
	SortItems:  true,
	StackItems: true,
})
for _, item := range detailed.Types["inventory"].Items {
	fmt.Printf("%dx %s: %.0f\n", item.Count, item.Name, item.Price)
}
```

## The result

`GetNetworth` returns a `*NetworthResult`:

```go
type NetworthResult struct {
	Networth            float64 // everything, coins included
	UnsoulboundNetworth float64 // without soulbound items
	NoInventory         bool    // no items could be parsed for the member
	IsNonCosmetic       bool
	Purse               float64
	Bank                float64 // the profile's (coop) bank
	PersonalBank        float64
	Types               map[string]*NetworthType
}

type NetworthType struct {
	Total            float64
	UnsoulboundTotal float64
	Items            []NetworthItemResult // empty with OnlyNetworth
}
```

The categories in `Types` are `armor`, `equipment`, `wardrobe`, `inventory`, `enderchest`, `accessories`,
`personal_vault`, `storage`, `fishing_bag`, `potion_bag`, `sacks_bag`, `candy_inventory`,
`carnival_mask_inventory`, `quiver`, `sacks`, `essence`, `pets`, `museum`, `farming_toolkit` and `hunting_toolkit`
(tools stored in the toolkits; ones taken out are counted in the inventory they're in).

Each `NetworthItemResult` has the item's `Name`, `Id`, `Price`, `BasePrice`, `Count`, whether it is `Soulbound` or
`Cosmetic`, and `Calculation`: the list of what added to the price (each enchantment, gemstone, reforge, and so on),
which is handy for a tooltip that explains a price.

## Pricing chosen inventories

To price a few inventories without a whole profile (a single backpack, a trade window, items from somewhere else),
pass the encoded inventories by name. The names become the keys of `Types`:

```go
inventories := networth.SpecifiedInventory{
	"inventory":  member.Inventory.Inventory,
	"enderchest": member.Inventory.Enderchest,
}
result, err := networth.CalculateFromSpecifiedInventories(inventories,
	networth.NetworthOptions{IncludeItemData: true}.ToInternal())
if err != nil {
	log.Fatal(err)
}
for _, item := range result.Types["enderchest"].Items {
	fmt.Println(item.Name, item.Price)
}
```

This is also a quick way to decode inventories for display: with `IncludeItemData` and `KeepInvalidItems`, every
item comes back with its decoded NBT (`ItemData`) and its price.

## Prices

Prices come from [SkyHelperBot/Prices](https://github.com/SkyHelperBot/Prices) and the item list from Hypixel's
`/v2/resources/skyblock/items`. Both are cached in process:

- Prices are reused for 5 minutes. After that the cached copy is still served while a fresh one downloads in the
  background, so a calculation never waits on GitHub once prices have loaded.
- The item list is kept for 12 hours.

Load them at startup so the first request doesn't wait:

```go
// Arguments: use the cache, cache time in seconds (0 = default), retries (0 = default).
if _, err := networth.GetPrices(true, 0, 0); err != nil {
	log.Printf("prices not loaded yet: %v", err)
}
if _, err := networth.GetItems(true, 0, 0); err != nil {
	log.Printf("items not loaded yet: %v", err)
}
```

If prices have never loaded (GitHub unreachable), the calculator returns an error that matches
`networth.ErrPricesUnavailable`, so you can answer with "try again shortly" instead of a generic failure:

```go
if errors.Is(err, networth.ErrPricesUnavailable) {
	// 503: prices are still loading
}
```

To use your own prices, pass them in `NetworthOptions.Prices`. Keys are item IDs as SkyHelper names them
(`HYPERION`, `LVL_100_GOLDEN_DRAGON`, `ENCHANTMENT_SHARPNESS_7`, ...).

## Performance

- Decoding inventories is most of the work. Create one calculator per profile member and reuse it for every
  calculation you need.
- With more than one CPU, inventories are decoded and priced in parallel, and normal and non-cosmetic networth run
  side by side. With `GOMAXPROCS=1` (a prefork child, for example) everything runs on the calling goroutine, which
  avoids goroutine overhead that couldn't run in parallel anyway.
- Use `OnlyNetworth` when you only need the totals.

## Testing

```bash
go test ./tests/...

# One area at a time
go test ./tests/ -run TestEnchantedBook
go test ./tests/ -run TestPets
go test ./tests/ -run TestGemstones
```

Results must match the Node SkyHelper module. When you change a calculator, add or update a test case built from
real item data.

## License

MIT, see [LICENSE.MD](LICENSE.MD).
