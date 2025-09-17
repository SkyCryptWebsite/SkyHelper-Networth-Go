package skyhelpernetworthgo

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/lib"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
	options "github.com/SkyCryptWebsite/SkyHelper-Networth-Go/options"
)

type ProfileNetworthCalculator struct {
	ProfileData         *skycrypttypes.Member
	MuseumData          *skycrypttypes.Museum
	Bank                float64
	Purse               float64
	PersonalBankBalance float64
	Items               *models.ParsedItems
	Prices              models.Prices
}

func NewProfileNetworthCalculator(userProfile *skycrypttypes.Member, museumData *skycrypttypes.Museum, bankBalance float64) (*ProfileNetworthCalculator, error) {
	items, err := lib.ParseItems(userProfile, museumData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse items: %w", err)
	}

	prices, err := lib.GetPrices(true, 5*60, 3)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch prices: %w", err)
	}

	_, err = lib.GetItems(true, 12*60*60, 3)
	if err != nil {
		return nil, fmt.Errorf("failed to get Hypixel items: %w", err)
	}

	return &ProfileNetworthCalculator{
		ProfileData:         userProfile,
		MuseumData:          museumData,
		Bank:                bankBalance,
		Purse:               float64(userProfile.Currencies.CoinPurse),
		PersonalBankBalance: userProfile.Profile.BankAccount,
		Items:               items,
		Prices:              *prices,
	}, nil
}

func (p *ProfileNetworthCalculator) GetNetworth(opts ...options.NetworthOptions) *models.NetworthResult {
	var opt options.NetworthOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	return p.calculate(opt.ToInternal())
}

func (p *ProfileNetworthCalculator) GetNonCosmeticNetworth(opts ...options.NetworthOptions) *models.NetworthResult {
	var opt options.NetworthOptions
	if len(opts) > 0 {
		opt = opts[0]
	}

	opt.NonCosmetic = true
	return p.calculate(opt.ToInternal())
}

func (p *ProfileNetworthCalculator) calculate(options models.NetworthOptions) *models.NetworthResult {
	if options.Prices != nil {
		p.Prices = options.Prices
	}

	calculatorService := calculators.NewCalculatorService()

	categories := map[string]models.CategoryInfo{
		"armor":                   {Items: p.Items.Armor, Type: "decoded"},
		"equipment":               {Items: p.Items.Equipment, Type: "decoded"},
		"wardrobe":                {Items: p.Items.Wardrobe, Type: "decoded"},
		"inventory":               {Items: p.Items.Inventory, Type: "decoded"},
		"enderchest":              {Items: p.Items.Enderchest, Type: "decoded"},
		"accessories":             {Items: p.Items.Accessories, Type: "decoded"},
		"personal_vault":          {Items: p.Items.PersonalVault, Type: "decoded"},
		"fishing_bag":             {Items: p.Items.FishingBag, Type: "decoded"},
		"potion_bag":              {Items: p.Items.PotionBag, Type: "decoded"},
		"sacks_bag":               {Items: p.Items.SacksBag, Type: "decoded"},
		"candy_inventory":         {Items: p.Items.CandyInventory, Type: "decoded"},
		"carnival_mask_inventory": {Items: p.Items.CarnivalMaskInventory, Type: "decoded"},
		"quiver":                  {Items: p.Items.Quiver, Type: "decoded"},
		"storage":                 {Items: p.Items.Storage, Type: "decoded"},
		"sacks":                   {Items: p.Items.Sacks, Type: "basic"},
		"essence":                 {Items: p.Items.Essence, Type: "basic"},
		"pets":                    {Items: p.Items.Pets, Type: "pets"},
		"museum":                  {Items: p.Items.Museum, Type: "decoded"},
	}

	totalNetworth, totalUnsoulboundNetworth := 0.0, 0.0
	output := make(map[string]*models.NetworthType, len(categories))
	for categoryId, categoryInfo := range categories {
		output[categoryId] = &models.NetworthType{
			Total:            0,
			UnsoulboundTotal: 0,
			Items:            []models.NetworthItemResult{},
		}

		switch categoryInfo.Type {
		case "decoded":
			decodedItems := categoryInfo.Items.([]*skycrypttypes.Item)
			for _, item := range decodedItems {
				if item.Tag == nil || item.Tag.ExtraAttributes == nil {
					if options.KeepInvalidItems {
						output[categoryId].Items = append(output[categoryId].Items, models.NetworthItemResult{})
					}

					continue
				}

				var result models.NetworthItemResult

				if item.Tag.ExtraAttributes.PetInfo != "" {
					var petData *skycrypttypes.Pet
					err := json.Unmarshal([]byte(item.Tag.ExtraAttributes.PetInfo), &petData)
					if err != nil {
						if options.KeepInvalidItems {
							output[categoryId].Items = append(output[categoryId].Items, models.NetworthItemResult{})
						}

						continue
					}

					petCalculator := calculatorService.NewSkyBlockPetCalculator(petData, p.Prices, options)
					calculatorService.CalculatePet(petCalculator)

					result = models.NetworthItemResult{
						Name:             petCalculator.PetName,
						LoreName:         petCalculator.PetName,
						Id:               petCalculator.PetData.Type,
						CustomId:         petCalculator.GetPetId(),
						Price:            petCalculator.Price + petCalculator.BasePrice,
						SoulboundPortion: 0, // Pets don't have soulbound portions in the same way
						BasePrice:        petCalculator.BasePrice,
						Calculation:      petCalculator.GetCalculation(),
						Count:            1,
						Soulbound:        petCalculator.IsSoulbound(),
						Cosmetic:         petCalculator.IsCosmetic(),
					}

					if options.IncludeItemData {
						result.ItemData = petData
					}
				} else {
					itemCalculator := calculatorService.NewSkyBlockItemCalculator(item, p.Prices, options)
					if itemCalculator.IsCosmetic() && options.NonCosmetic {
						continue
					}

					calculatorService.CalculateItem(itemCalculator)

					result = models.NetworthItemResult{
						Name:             itemCalculator.ItemName,
						LoreName:         item.Tag.Display.Name,
						Id:               itemCalculator.ExtraAttributes.Id,
						CustomId:         itemCalculator.ItemId,
						Price:            itemCalculator.GetPrice(),
						SoulboundPortion: itemCalculator.SoulboundPortion,
						BasePrice:        itemCalculator.BasePrice,
						Calculation:      itemCalculator.GetCalculation(),
						Count:            itemCalculator.Count,
						Soulbound:        itemCalculator.IsSoulbound(),
						Cosmetic:         itemCalculator.IsCosmetic(),
					}

					if options.IncludeItemData {
						result.ItemData = item
					}
				}

				if result.Price == 0 {
					if options.KeepInvalidItems {
						output[categoryId].Items = append(output[categoryId].Items, models.NetworthItemResult{})
					}

					continue
				}

				sort.Slice(result.Calculation, func(i, j int) bool {
					return result.Calculation[i].Price > result.Calculation[j].Price
				})

				totalNetworth += result.Price
				output[categoryId].Total += result.Price
				if !result.Soulbound {
					output[categoryId].UnsoulboundTotal += result.Price - result.SoulboundPortion
					totalUnsoulboundNetworth += result.Price - result.SoulboundPortion
				}

				if !options.OnlyNetworth {
					output[categoryId].Items = append(output[categoryId].Items, result)
				}
			}

		case "basic":
			// Handle sacks and essence items
			basicItems := categoryInfo.Items.([]*models.BasicItem)
			for _, item := range basicItems {
				if strings.HasPrefix(item.Id, "RUNE") && options.NonCosmetic {
					continue
				}

				itemCalculator := calculatorService.NewBasicItemCalculator(item, p.Prices)
				calculatorService.CalculateBasicItem(itemCalculator)

				result := models.NetworthItemResult{
					Name:             itemCalculator.ItemName,
					LoreName:         itemCalculator.ItemName, // Basic items don't have display names
					Id:               itemCalculator.ItemId,
					CustomId:         itemCalculator.ItemId,
					Price:            itemCalculator.Price + itemCalculator.BasePrice,
					SoulboundPortion: itemCalculator.SoulboundPortion,
					BasePrice:        itemCalculator.BasePrice,
					Calculation:      itemCalculator.Calculation,
					Count:            itemCalculator.Amount,
					Soulbound:        false, // Basic items are not soulbound
					Cosmetic:         false, // Basic items are not cosmetic
				}

				if options.IncludeItemData {
					result.ItemData = item
				}

				if result.Price == 0 {
					continue
				}

				sort.Slice(result.Calculation, func(i, j int) bool {
					return result.Calculation[i].Price > result.Calculation[j].Price
				})

				totalNetworth += result.Price
				output[categoryId].Total += result.Price
				if !result.Soulbound {
					output[categoryId].UnsoulboundTotal += result.Price - result.SoulboundPortion
					totalUnsoulboundNetworth += result.Price - result.SoulboundPortion
				}

				if !options.OnlyNetworth {
					output[categoryId].Items = append(output[categoryId].Items, result)
				}
			}

		case "pets":
			pets := categoryInfo.Items.([]*skycrypttypes.Pet)
			for _, pet := range pets {
				petCalculator := calculatorService.NewSkyBlockPetCalculator(pet, p.Prices, options)
				calculatorService.CalculatePet(petCalculator)

				result := models.NetworthItemResult{
					Name:             petCalculator.PetName,
					LoreName:         petCalculator.PetName,
					Id:               petCalculator.PetData.Type,
					CustomId:         petCalculator.GetPetId(),
					Price:            petCalculator.Price + petCalculator.BasePrice,
					SoulboundPortion: 0, // Pets don't have soulbound portions in the same way
					BasePrice:        petCalculator.BasePrice,
					Calculation:      petCalculator.GetCalculation(),
					Count:            1,
					Soulbound:        petCalculator.IsSoulbound(),
					Cosmetic:         petCalculator.IsCosmetic(),
				}

				if options.IncludeItemData {
					result.ItemData = pet
				}

				if result.Price == 0 {
					continue
				}

				sort.Slice(result.Calculation, func(i, j int) bool {
					return result.Calculation[i].Price > result.Calculation[j].Price
				})

				totalNetworth += result.Price
				output[categoryId].Total += result.Price
				if !result.Soulbound {
					output[categoryId].UnsoulboundTotal += result.Price - result.SoulboundPortion
					totalUnsoulboundNetworth += result.Price - result.SoulboundPortion
				}

				if !options.OnlyNetworth {
					output[categoryId].Items = append(output[categoryId].Items, result)
				}
			}
		}

		if options.SortItems {
			sort.Slice(output[categoryId].Items, func(i, j int) bool {
				return output[categoryId].Items[i].Price > output[categoryId].Items[j].Price
			})
		}

		if options.StackItems {
			var stackedItems []models.NetworthItemResult

			for _, item := range output[categoryId].Items {
				found := false
				for i := range stackedItems {
					existing := &stackedItems[i]

					if (existing.CustomId == item.CustomId || existing.Id == item.Id) &&
						existing.Price/float64(existing.Count) == item.Price/float64(item.Count) &&
						existing.Soulbound == item.Soulbound {

						existing.Price += item.Price
						existing.Count += item.Count

						if existing.BasePrice == 0 {
							existing.BasePrice = item.BasePrice
						}
						if len(existing.Calculation) == 0 {
							existing.Calculation = item.Calculation
						}

						found = true
						break
					}
				}

				if !found {
					stackedItems = append(stackedItems, item)
				}

				output[categoryId].Items = stackedItems
			}

		}
	}

	return &models.NetworthResult{
		Networth:            totalNetworth + p.Purse + p.Bank + p.PersonalBankBalance,
		UnsoulboundNetworth: totalUnsoulboundNetworth + p.Purse + p.Bank + p.PersonalBankBalance,
		NoInventory:         p.Items == nil,
		IsNonCosmetic:       options.NonCosmetic,
		Purse:               p.Purse,
		Bank:                p.Bank,
		PersonalBank:        p.PersonalBankBalance,
		Types:               output,
	}
}
