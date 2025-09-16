package calculators

import (
	skycrypttypes "github.com/DuckySoLucky/SkyCrypt-Types"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/calculators/handlers"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/lib"
	"github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"
)

type CalculatorService struct {
	itemHandlers []models.Handler
	petHandlers  []models.PetHandler
	itemProvider models.ItemProvider
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{
		itemHandlers: handlers.ItemHandlers,
		petHandlers:  handlers.PetHandlers,
		itemProvider: lib.NewItemProviderAdapter(),
	}
}

func (cs *CalculatorService) NewSkyBlockItemCalculator(item *skycrypttypes.Item, prices models.Prices, options models.NetworthOptions) *models.NetworthItem {
	return models.NewSkyBlockItemCaclulator(item, prices, cs.itemProvider, options)
}

func (cs *CalculatorService) NewSkyBlockPetCalculator(pet *skycrypttypes.Pet, prices models.Prices, options models.NetworthOptions) *models.NetworthPet {
	return models.NewSkyBlockPetCalculator(pet, prices, options)
}

func (cs *CalculatorService) CalculateItem(item *models.NetworthItem) {
	item.Calculate(cs.itemHandlers)
}

func (cs *CalculatorService) CalculatePet(pet *models.NetworthPet) {
	pet.Calculate(cs.petHandlers)
}

func (cs *CalculatorService) NewBasicItemCalculator(item *models.BasicItem, prices models.Prices) *models.BasicNetworthItem {
	return models.NewBasicItemNetworthCalculator(item, prices, cs.itemProvider)
}

func (cs *CalculatorService) CalculateBasicItem(item *models.BasicNetworthItem) {
	item.Calculate()
}
