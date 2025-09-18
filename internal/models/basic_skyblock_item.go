package models

import (
	"strings"
)

type BasicNetworthItem struct {
	ItemId       string       `json:"itemId"`
	Amount       int          `json:"amount"`
	ItemName     string       `json:"itemName"`
	SkyblockItem *HypixelItem `json:"skyblockItem"`

	Prices map[string]float64 `json:"prices"`

	NonCosmetic      bool              `json:"nonCosmetic"`
	Calculation      []CalculationData `json:"calculation"`
	BasePrice        float64           `json:"basePrice"`
	Price            float64           `json:"price"`
	SoulboundPortion float64           `json:"soulboundPortion"`
}

func NewBasicItemNetworthCalculator(item *BasicItem, prices map[string]float64, itemProvider ItemProvider) *BasicNetworthItem {
	networthItem := &BasicNetworthItem{
		ItemId:       item.Id,
		Amount:       item.Amount,
		SkyblockItem: itemProvider.GetItem(item.Id),

		Prices: prices,

		NonCosmetic:      false,
		Calculation:      []CalculationData{},
		BasePrice:        0.0,
		Price:            0.0,
		SoulboundPortion: 0.0,
	}

	networthItem.getItemName()

	return networthItem
}

func (item *BasicNetworthItem) getItemName() {
	if strings.Contains(item.ItemId, "ESSENCE") {
		parts := strings.Split(item.ItemId, "_")
		for i := len(parts) - 1; i >= 0; i-- {
			parts[i] = titleCase(parts[i])
		}

		item.ItemName = strings.Join(parts, " ")
		return
	}

	if item.SkyblockItem != nil && item.SkyblockItem.Name != "" {
		item.ItemName = item.SkyblockItem.Name
		return

	}

	item.ItemName = titleCase(item.ItemId)
}

func (item *BasicNetworthItem) Calculate() {
	itemPrice := item.Prices[item.ItemId]
	if itemPrice == 0 {
		return
	}

	totalPrice := itemPrice * float64(item.Amount)
	if totalPrice == 0 {
		return
	}

	item.Price += totalPrice
}
