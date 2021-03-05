package repository

import (
	"errors"

	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	csv_file_reader "github.com/ToteEmmanuel/academy-go-q12021/tools/reader"
)

type pokeRepository struct {
	storage *csv_file_reader.CsvPokeStorage
}

type PokeRepository interface {
	FindById(id int32) (*model.Pokemon, error)
	FindAll() []*model.Pokemon
}

func NewPokeRepository(storage *csv_file_reader.CsvPokeStorage) PokeRepository {
	return &pokeRepository{storage}
}

func (pR *pokeRepository) FindById(id int32) (*model.Pokemon, error) {
	pokemon := pR.storage.FindById(id)
	if pokemon == nil {
		return nil, errors.New("not found")
	}
	return pokemon, nil
}

func (pR *pokeRepository) FindAll() []*model.Pokemon {
	pokemon := pR.storage.FindAll()
	return pokemon
}
