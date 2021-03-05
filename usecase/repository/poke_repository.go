package repository

import "github.com/ToteEmmanuel/academy-go-q12021/domain/model"

type PokeRepository interface {
	FindById(id int32) (*model.Pokemon, error)
	FindAll() []*model.Pokemon
}
