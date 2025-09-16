package handlers

import "github.com/SkyCryptWebsite/SkyHelper-Networth-Go/internal/models"

var ItemHandlers []models.Handler = []models.Handler{
	ArtOfPeaceHandler{},
	ArtOfWarHandler{},
	BoosterHandler{},
	DivanPowderCoatingHandler{},
	DrillPartsHandler{},
	DyeHandler{},
	EnchantedBookHandler{},
	EnrichmentHandler{},
	EssenceStarsHandler{},
	EtherwarpConduitHandler{},
	FarmingForDummiesHandler{},
	GemstonePowerScrollHandler{},
	GemstonesHandler{},
	ItemEnchantments{},
	JalapenoBookHandler{},
	ManaDisintegratorHandler{},
	MasterStarsHandler{},
	MidasWeaponHandler{},
	NecronBladeScrollHandler{},
	NewYearCakeBagHandler{},

	PickonimbusHandler{},
	PocketSackInASackHandler{},
	PolarvoidBookHandler{},
	PotatoBookHandler{},
	PrestigeHandler{},
	PulseRingThunderHandler{},
	RecombobulatorHandler{},
	ReforgeHandler{},
	RodPartsHandler{},
	RuneHandler{},
	ShensAuctionHandler{},

	SoulboundSkinHandler{},
	TransmissionTunerHandler{},
	WoodSingularityHandler{},
}

// PetHandlers will be implemented properly later
var PetHandlers []models.PetHandler = []models.PetHandler{
	PetCandyHandler{},
	SoulboundPetSkinHandler{},
	PetItemHandler{}, // MUST BE LAST
}
