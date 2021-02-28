package repository

import "github.com/Topi99/academy-go-q12021/domain/model"

// PokemonRepository interface
type PokemonRepository interface {
	GetOne(id uint) (*model.Pokemon, error)
}
