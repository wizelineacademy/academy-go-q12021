package repository

import (
	"errors"
	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	cvsfilereader "github.com/ToteEmmanuel/academy-go-q12021/tools/reader"
)

type PokeRepository interface {
	FindById(id int32) (*model.Pokemon, error)
	FindAll() ([]*model.Pokemon, error)
}

type pokeRepository struct {
	storage *cvsfilereader.CsvPokeStorage
}

func NewPokeRepository(storage *cvsfilereader.CsvPokeStorage) PokeRepository {
	return &pokeRepository{storage}
}

func (pR *pokeRepository) FindById(id int32) (*model.Pokemon, error) {
	pokemon := pR.storage.FindById(id)
	if pokemon == nil {
		return nil, errors.New("not found")
	}
	return pokemon, nil
}

func (pR *pokeRepository) FindAll() ([]*model.Pokemon, error) {
	pokemon := pR.storage.FindAll()
	return pokemon, nil
}
