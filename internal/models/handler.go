package models

type Handler interface {
	Applies(item *NetworthItem) bool
	Calculate(item *NetworthItem, prices Prices)
	IsCosmetic() bool
}

type PetHandler interface {
	Applies(pet *NetworthPet) bool
	Calculate(pet *NetworthPet, prices Prices)
	IsCosmetic() bool
}
