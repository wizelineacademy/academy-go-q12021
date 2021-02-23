package models

type Pokemon struct {
	
	Id int						`json:id`
	PokedexNumber int	`json:pokedex_number`
	Name string				`json:name`
	Types string			`json:types`
	Region string			`json:region`
}

type PokemonList []Pokemon