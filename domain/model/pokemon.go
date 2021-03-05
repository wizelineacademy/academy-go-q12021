package model

//Representation of a Pokemon
type Pokemon struct {
	Id         int32
	Species    string
	Sprite     string
	FlavorText string
	Types      []string
}

//Creates an struct with pokemagic
func NewPokemon(id int32, species, sprite, flavorText string, types ...string) *Pokemon {
	p := Pokemon{
		Id:         id,
		Species:    species,
		Sprite:     sprite,
		FlavorText: flavorText,
		Types:      types,
	}
	return &p
}
