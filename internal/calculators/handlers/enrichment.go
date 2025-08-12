package handlers

import (
	"math"
	"strings"

	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/constants"
	"github.com/DuckySoLucky/SkyHelper-Networth-Go/internal/models"
)

type EnrichmentHandler struct{}

func (h EnrichmentHandler) IsCosmetic() bool {
	return false
}

func (h EnrichmentHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.TalismanEnrichment != ""
}

func (h EnrichmentHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	// Find the minimum enrichment priceww
	enrichmentPrice := math.Inf(1) // Start with positive infinity
	for _, enrichment := range constants.ENRICHMENTS {
		if price, exists := prices[enrichment]; exists && price < enrichmentPrice {
			enrichmentPrice = price
		}
	}

	if enrichmentPrice != math.Inf(1) {
		calculationData := models.CalculationData{
			Id:    strings.ToUpper(item.ExtraAttributes.TalismanEnrichment),
			Type:  "TALISMAN_ENRICHMENT",
			Price: enrichmentPrice * constants.APPLICATION_WORTH["enrichment"],
			Count: 1,
		}

		item.Price += calculationData.Price
		item.Calculation = append(item.Calculation, calculationData)
	}
}
