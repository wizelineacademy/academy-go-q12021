package models

type Pokemon struct {
	
	Id int				`json:id`
	Name string		`json:name`
	Types string	`json:types`
	Region string	`json:region`
}

type PokemonList []Pokemon