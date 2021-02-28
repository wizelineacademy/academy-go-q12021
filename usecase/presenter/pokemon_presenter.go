package presenter

import "github.com/Topi99/academy-go-q12021/domain/model"

// PokemonPresenter interface
type PokemonPresenter interface {
	GetOne(p *model.Pokemon) *model.Pokemon
}
