package lib

import (
	"fmt"
	"strings"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

func ParseItems(profileData *skycrypttypes.Member, museumData *skycrypttypes.Museum) (*models.ParsedItems, error) {
	sharedInventory := profileData.SharedInventory
	inventory := profileData.Inventory
	if inventory == nil {
		inventory = &skycrypttypes.Inventory{}
	}

	itemsToDecode := map[string]string{
		"armor":                   inventory.Armor.Data,
		"equipment":               inventory.Equipment.Data,
		"wardrobe":                inventory.Wardrobe.Data,
		"inventory":               inventory.Inventory.Data,
		"enderchest":              inventory.Enderchest.Data,
		"accessories":             inventory.BagContents.TalismanBag.Data,
		"personal_vault":          inventory.PersonalVault.Data,
		"fishing_bag":             inventory.BagContents.FishingBag.Data,
		"potion_bag":              inventory.BagContents.PotionBag.Data,
		"sacks_bag":               inventory.BagContents.SacksBag.Data,
		"candy_inventory":         sharedInventory.CandyInventory.Data,
		"carnival_mask_inventory": sharedInventory.CarnivalMaskInventory.Data,
		"quiver":                  inventory.BagContents.Quiver.Data,
	}

	for backpack, contents := range inventory.Backpack {
		itemsToDecode["storage_"+backpack] = contents.Data
	}

	for backpack, icon := range inventory.BackpackIcons {
		itemsToDecode["storage_icon_"+backpack] = icon.Data
	}

	var items models.ParsedItems
	for key, data := range itemsToDecode {
		decodedItems, err := DecodeInventory(data)
		if err != nil {
			return nil, fmt.Errorf("failed to decode inventory for %s: %w", key, err)
		}

		categoryItems := make([]*skycrypttypes.Item, 0)
		for idx := range decodedItems.Items {
			item := &decodedItems.Items[idx]
			if item.Tag == nil || item.Tag.ExtraAttributes == nil {
				continue
			}

			categoryItems = append(categoryItems, item)
		}
		if strings.HasPrefix(key, "storage") {
			items.Storage = append(items.Storage, categoryItems...)
		} else {
			switch key {
			case "armor":
				items.Armor = categoryItems
			case "equipment":
				items.Equipment = categoryItems
			case "wardrobe":
				items.Wardrobe = categoryItems
			case "inventory":
				items.Inventory = categoryItems
			case "enderchest":
				items.Enderchest = categoryItems
			case "accessories":
				items.Accessories = categoryItems
			case "personal_vault":
				items.PersonalVault = categoryItems
			case "fishing_bag":
				items.FishingBag = categoryItems
			case "potion_bag":
				items.PotionBag = categoryItems
			case "sacks_bag":
				items.SacksBag = categoryItems
			case "candy_inventory":
				items.CandyInventory = categoryItems
			case "carnival_mask_inventory":
				items.CarnivalMaskInventory = categoryItems
			case "quiver":
				items.Quiver = categoryItems
			}
		}
	}

	if museumData != nil {
		if museumData.Items == nil {
			museumData.Items = &map[string]skycrypttypes.MuseumItem{}
		}

		if museumData.Special == nil {
			museumData.Special = &[]skycrypttypes.MuseumItem{}
		}

		for _, item := range *museumData.Items {
			if item.Borrowing {
				continue
			}

			decodedItems, _ := DecodeInventory(item.Items.Data)
			for _, item := range decodedItems.Items {
				items.Museum = append(items.Museum, &item)
			}
		}

		for _, item := range *museumData.Special {
			decodedItems, _ := DecodeInventory(item.Items.Data)
			for _, item := range decodedItems.Items {
				items.Museum = append(items.Museum, &item)
			}
		}
	}

	postParseItems(*profileData, &items)
	return &items, nil
}

func postParseItems(profileData skycrypttypes.Member, items *models.ParsedItems) error {
	categories := [][]*skycrypttypes.Item{
		items.Armor,
		items.Equipment,
		items.Wardrobe,
		items.Inventory,
		items.Enderchest,
		items.Accessories,
		items.PersonalVault,
		items.FishingBag,
		items.PotionBag,
		items.SacksBag,
		items.CandyInventory,
		items.CarnivalMaskInventory,
		items.Quiver,
		items.Storage,
	}

	for _, categoryItems := range categories {
		for _, item := range categoryItems {
			if item.Tag == nil || item.Tag.ExtraAttributes == nil || item.Tag.ExtraAttributes.NewYearCakeBagData == nil {
				continue
			}

			if item.Tag.ExtraAttributes.NewYearCakeBagData != nil {
				cakeBagData, err := DecodeFromBytes(item.Tag.ExtraAttributes.NewYearCakeBagData)
				if err != nil {
					return fmt.Errorf("failed to decode New Year Cake Bag data: %w", err)
				}

				cakeBagYears := make([]int, 0)
				for _, cake := range cakeBagData.Items {
					if cake.Tag == nil || cake.Tag.ExtraAttributes == nil {
						continue
					}

					cakeBagYears = append(cakeBagYears, cake.Tag.ExtraAttributes.NewYearsCake)
				}

				item.Tag.ExtraAttributes.NewYearCakeBagYears = cakeBagYears
			}
		}
	}

	for essence, data := range profileData.Currencies.Essence {
		items.Essence = append(items.Essence, &models.BasicItem{
			Id:     "ESSENCE_" + strings.ToUpper(essence),
			Amount: data.Current,
		})
	}
	var sackCounts map[string]int
	if profileData.SackCounts != nil {
		sackCounts = profileData.SackCounts
	} else if profileData.Inventory != nil {
		sackCounts = profileData.Inventory.Sacks
	}
	for item, amount := range sackCounts {
		if amount <= 0 {
			continue
		}
		items.Sacks = append(items.Sacks, &models.BasicItem{
			Id:     item,
			Amount: amount,
		})
	}

	if profileData.Pets != nil {
		for _, pet := range profileData.Pets.Pets {
			items.Pets = append(items.Pets, &pet)
		}
	}
	return nil
}
