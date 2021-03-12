package repository

import "github.com/AlejandroSeguraWIZ/academy-go-q12021/domain/model"

type PokemonRepository interface {
	FetchAll() ([]model.Pokemon, error)
}
