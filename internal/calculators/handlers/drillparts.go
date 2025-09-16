package handlers

import (
	"strings"

	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type DrillPartsHandler struct{}

func (h DrillPartsHandler) IsCosmetic() bool {
	return false
}

func (h DrillPartsHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.DrillPartUpgradeModule != "" ||
		item.ExtraAttributes.DrillPartFuelTank != "" ||
		item.ExtraAttributes.DrillPartEngine != ""
}

func (h DrillPartsHandler) Calculate(item *models.NetworthItem, prices models.Prices) {
	parts := []struct {
		FieldValue string
		FieldName  string
	}{
		{item.ExtraAttributes.DrillPartUpgradeModule, "DRILL_PART_UPGRADE_MODULE"},
		{item.ExtraAttributes.DrillPartFuelTank, "DRILL_PART_FUEL_TANK"},
		{item.ExtraAttributes.DrillPartEngine, "DRILL_PART_ENGINE"},
	}

	for _, part := range parts {
		if part.FieldValue != "" {
			id := strings.ToUpper(part.FieldValue)
			calculationData := models.CalculationData{
				Id:    id,
				Type:  "DRILL_PART",
				Price: prices[id] * constants.APPLICATION_WORTH["drillPart"],
				Count: 1,
			}

			item.Price += calculationData.Price
			item.Calculation = append(item.Calculation, calculationData)
		}
	}
}
