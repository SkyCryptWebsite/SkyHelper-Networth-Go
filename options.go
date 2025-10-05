package skyhelpernetworthgo

import "github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"

type NetworthOptions struct {
	Prices           map[string]float64 `json:"prices"`
	NonCosmetic      bool               `json:"nonCosmetic"`
	CachePrices      bool               `json:"cachePrices"`
	OnlyNetworth     bool               `json:"onlyNetworth"`
	IncludeItemData  bool               `json:"includeItemData"`
	SortItems        bool               `json:"sortItems"`
	StackItems       bool               `json:"stackItems"`
	KeepInvalidItems bool               `json:"keepInvalidItems"`
}

func (opts NetworthOptions) ToInternal() models.NetworthOptions {
	return models.NetworthOptions{
		Prices:           opts.Prices,
		NonCosmetic:      opts.NonCosmetic,
		CachePrices:      opts.CachePrices,
		OnlyNetworth:     opts.OnlyNetworth,
		IncludeItemData:  opts.IncludeItemData,
		SortItems:        opts.SortItems,
		StackItems:       opts.StackItems,
		KeepInvalidItems: opts.KeepInvalidItems,
	}
}
