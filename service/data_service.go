package service

import (
	"github.com/grethelBello/academy-go-q12021/model"
)

// DataService is an interface for the modules to implement the specigic logic of the entities
type DataService interface {
	Init() error
	Get(id int) model.Response
	List(count, page int) model.Response
	Filter(typeFilter model.TypeFilter, items, itemsPerWorker int) model.Response
	Sync() error
}
