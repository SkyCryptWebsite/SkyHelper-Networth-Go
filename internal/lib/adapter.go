package lib

import "github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"

type ItemProviderAdapter struct{}

func NewItemProviderAdapter() *ItemProviderAdapter {
	return &ItemProviderAdapter{}
}

func (adapter *ItemProviderAdapter) GetItem(itemId string) *models.HypixelItem {
	return GetItem(itemId)
}
