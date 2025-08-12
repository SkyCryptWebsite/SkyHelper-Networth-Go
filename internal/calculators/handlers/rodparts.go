package handlers

import (
	"strings"

	"github.com/duckysolucky/skyhelper-networth-go/internal/constants"
	"github.com/duckysolucky/skyhelper-networth-go/internal/models"
)

type RodPartsHandler struct{}

func (h RodPartsHandler) IsCosmetic() bool {
	return false
}

func (h RodPartsHandler) Applies(item *models.NetworthItem) bool {
	hasRodPart := false
	ea := item.ExtraAttributes
	if ea.Line.Part != "" || ea.Hook.Part != "" || ea.Sinker.Part != "" {
		hasRodPart = true
	}
	return hasRodPart
}

func (h RodPartsHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	ea := item.ExtraAttributes
	rodParts := []struct {
		key       string
		value     string
		Soulbound bool
	}{
		{"line", ea.Line.Part, ea.Line.Soulbound},
		{"hook", ea.Hook.Part, ea.Hook.Soulbound},
		{"sinker", ea.Sinker.Part, ea.Sinker.Soulbound},
	}
	for _, part := range rodParts {
		if part.value != "" {
			calculationData := models.CalculationData{
				Id:        strings.ToUpper(part.value),
				Type:      "ROD_PART",
				Price:     prices[strings.ToUpper(part.value)] * constants.APPLICATION_WORTH["rodPart"],
				Count:     1,
				Soulbound: part.Soulbound,
			}

			item.Price += calculationData.Price
			item.Calculation = append(item.Calculation, calculationData)
			if part.Soulbound {
				item.SoulboundPortion += calculationData.Price
			}
		}
	}
}
