package repository

import (
	"errors"
	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
	"github.com/ToteEmmanuel/academy-go-q12021/infrastructure/datastore"
)

type PokeRepository interface {
	FindById(id int32) (*model.Pokemon, error)
	FindAll() ([]*model.Pokemon, error)
	Save(*model.Pokemon) (*model.Pokemon, error)
	FindAllWorkers(query string, items int, worker int) ([]*model.Pokemon, error)
}

type pokeRepository struct {
	storage *datastore.CsvPokeStorage
}

func NewPokeRepository(storage *datastore.CsvPokeStorage) PokeRepository {
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

func (pR *pokeRepository) FindAllWorkers(typeQuery string, items, itemsPerWorker int) ([]*model.Pokemon, error) {
	pokemon, err:= pR.storage.FindAllWorkers(typeQuery, items, itemsPerWorker)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

func (pR *pokeRepository) Save(pokemon *model.Pokemon) (*model.Pokemon, error) {
	pokemon, err := pR.storage.Save(pokemon)
	if(err != nil) {
		return nil, err
	}
	return pokemon, nil
}
