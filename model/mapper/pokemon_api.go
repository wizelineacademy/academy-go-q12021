package mapper

import (
	"github.com/wizelineacademy/academy-go-q12021/model"
)

func PokemonAPItoPokemon(pokemonAPI model.PokemonAPI) model.Pokemon {
	var type1 string
	var type2 string
	if len(pokemonAPI.Types) < 2 {
		type1 = pokemonAPI.Types[0].Type.Name
		type2 = ""
	} else {
		type1 = pokemonAPI.Types[0].Type.Name
		type2 = pokemonAPI.Types[1].Type.Name
	}
	return model.Pokemon{
		Id:             pokemonAPI.Id,
		Name:           pokemonAPI.Name,
		Height:         pokemonAPI.Height,
		Weight:         pokemonAPI.Weight,
		BaseExperience: pokemonAPI.BaseExperience,
		PrimaryType:    type1,
		SecondaryType:  type2,
	}
}
