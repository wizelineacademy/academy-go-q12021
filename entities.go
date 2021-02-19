package main

// Pokemon Type
type Pokemon struct {
	ID     int    `json:ID`
	Name   string `json:Name`
	Weight string `json:Weight`
	Height string `json:Height`
}

// AllPokemons Type
type AllPokemons []Pokemon
