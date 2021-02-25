package model

type Pokemon struct {
	
	Id int						`json:"id"`
	PokedexNumber int	`json:"pokedexNumber"`
	Name string				`json:"name"`
	Types string			`json:"types"`
	Region string			`json:"region"`
}

type PokemonList []Pokemon
