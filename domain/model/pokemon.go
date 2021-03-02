package model

// Pokemon is a struct that contains basic Pokemon information
type Pokemon struct {
	
	Id int						`json:"id"`
	PokedexNumber int	`json:"pokedexNumber"`
	Name string				`json:"name"`
	Types string			`json:"types"`
	Region string			`json:"region"`
}

// PokemonList is an array of Pokemon
type PokemonList []Pokemon
