package models

type Handler interface {
	Applies(item *NetworthItem) bool
	Calculate(item *NetworthItem, prices map[string]float64)
	IsCosmetic() bool
}

type PetHandler interface {
	Applies(pet *NetworthPet) bool
	Calculate(pet *NetworthPet, prices map[string]float64)
	IsCosmetic() bool
}
