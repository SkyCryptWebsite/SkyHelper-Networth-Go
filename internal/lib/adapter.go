package lib

import "duckysolucky/skyhelper-networth-go/internal/models"

type ItemProviderAdapter struct{}

func NewItemProviderAdapter() *ItemProviderAdapter {
	return &ItemProviderAdapter{}
}

func (adapter *ItemProviderAdapter) GetItem(itemId string) *models.HypixelItem {
	return GetItem(itemId)
}
