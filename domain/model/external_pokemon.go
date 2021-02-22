package model

// ExternalPokemon is a type to handle pokemons from the API from the results object
type ExternalPokemon struct {
	Name string `json:name`
	URL  string `json:url`
}

// PokemonFromAPI is a type to handle the main object result from the call to the API
type PokemonFromAPI struct {
	Count    int               `json:count`
	Next     string            `json:next`
	Previous string            `json:previous`
	Results  []ExternalPokemon `json:results`
}
