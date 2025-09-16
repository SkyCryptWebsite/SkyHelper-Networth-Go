package lib

import "github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"

type ItemProviderAdapter struct{}

func NewItemProviderAdapter() *ItemProviderAdapter {
	return &ItemProviderAdapter{}
}

func (adapter *ItemProviderAdapter) GetItem(itemId string) *models.HypixelItem {
	return GetItem(itemId)
}
