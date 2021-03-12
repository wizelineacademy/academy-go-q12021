package presenter

import "github.com/AlejandroSeguraWIZ/academy-go-q12021/domain/model"

type PokemonPresenter interface {
	ResponsePokemons([]model.Pokemon) []model.PokemonResponse
}
