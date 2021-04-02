package datastore

import "github.com/ToteEmmanuel/academy-go-q12021/domain/model"

//mockgen -source=infrastructure/datastore/pokestorage.go -destination=infrastructure/datastore/mock/pokestorage_mock.go mock PokeStorage
type PokeStorage interface {
	FindById(id int) *model.Pokemon
	FindAll() []*model.Pokemon
	Save(pokemon *model.Pokemon) (*model.Pokemon, error)
	FindAllWorkers(typeStr string, items int, itemsPerWorker int) ([]*model.Pokemon, error)
}
