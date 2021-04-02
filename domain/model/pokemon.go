package model

//Representation of a Pokemon
type Pokemon struct {
	Id         int      `json:"id"`
	Species    string   `json:"species"`
	Sprite     string   `json:"sprite"`
	FlavorText string   `json:"flavorText"`
	Types      []string `json:"types"`
}

//Creates an struct with pokemagic
func NewPokemon(id int, species, sprite, flavorText string, types ...string) *Pokemon {
	p := Pokemon{
		Id:         id,
		Species:    species,
		Sprite:     sprite,
		FlavorText: flavorText,
		Types:      types,
	}
	return &p
}
