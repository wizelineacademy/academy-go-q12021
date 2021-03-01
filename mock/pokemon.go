package mock

import "pokeapi/model"

func MockPokemon() *[]model.Pokemon {
	Pokemons := []model.Pokemon{
		{ID: 1, Name: "greninja", URL: "https://pokeapi.co/api/v2/pokemon/658/"},
		{ID: 2, Name: "ursaring", URL: "https://pokeapi.co/api/v2/pokemon/217/"},
		{ID: 3, Name: "arcanine", URL: "https://pokeapi.co/api/v2/pokemon/59/"},
		{ID: 4, Name: "gengar", URL: "https://pokeapi.co/api/v2/pokemon/94/"},
		{ID: 5, Name: "porygon", URL: "https://pokeapi.co/api/v2/pokemon/137/"},
	}
	return &Pokemons
}
