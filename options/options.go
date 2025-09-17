package options

import "github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"

type NetworthOptions struct {
	Prices           map[string]float64 `json:"prices"`
	NonCosmetic      bool               `json:"nonCosmetic"`
	CachePrices      bool               `json:"cachePrices"`
	OnlyNetworth     bool               `json:"onlyNetworth"`
	IncludeItemData  bool               `json:"includeItemData"`
	SortItems        bool               `json:"sortItems"`
	StackItems       bool               `json:"stackItems"`
	KeepInvalidItems bool               `json:"KeepInvalidItems"`
}

func (opts NetworthOptions) ToInternal() models.NetworthOptions {
	var prices models.Prices
	if opts.Prices != nil {
		prices = models.Prices(opts.Prices)
	}

	return models.NetworthOptions{
		Prices:           prices,
		NonCosmetic:      opts.NonCosmetic,
		CachePrices:      opts.CachePrices,
		OnlyNetworth:     opts.OnlyNetworth,
		IncludeItemData:  opts.IncludeItemData,
		SortItems:        opts.SortItems,
		StackItems:       opts.StackItems,
		KeepInvalidItems: opts.KeepInvalidItems,
	}
}
