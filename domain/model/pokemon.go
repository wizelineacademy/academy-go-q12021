package model

// Pokemon is the core type pokemon, structure get from the csv
type Pokemon struct {
	ID   int    `json:ID`
	Name string `json:Name`
	URL  string `json:Url`
}

// AllPokemons is a type build from Pokemon just to have a type for an array of Pokemons
type AllPokemons []Pokemon
