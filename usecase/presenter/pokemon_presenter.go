package presenter

import "github.com/Topi99/academy-go-q12021/domain/model"

// PokemonPresenter interface
type PokemonPresenter interface {
	ResponsePokemon(p *model.Pokemon) *model.Pokemon
}
