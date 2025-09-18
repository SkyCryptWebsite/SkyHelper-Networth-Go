package handlers

import (
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/constants"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type ShensAuctionHandler struct{}

func (h ShensAuctionHandler) IsCosmetic() bool {
	return false
}

func (h ShensAuctionHandler) Applies(item *models.NetworthItem) bool {
	return item.ExtraAttributes.Price > 0 && item.ExtraAttributes.Auction > 0 && item.ExtraAttributes.Bid > 0
}

func (h ShensAuctionHandler) Calculate(item *models.NetworthItem, prices map[string]float64) {
	pricePaid := float64(item.ExtraAttributes.Price) * constants.APPLICATION_WORTH["shensAuctionPrice"]
	if pricePaid > item.BasePrice {
		calculationData := models.CalculationData{
			Id:    item.ItemId,
			Type:  "SHENS_AUCTION",
			Price: pricePaid,
			Count: 1,
		}

		item.BasePrice = pricePaid
		item.Calculation = append(item.Calculation, calculationData)

	}
}
