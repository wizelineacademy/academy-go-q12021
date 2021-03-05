package registry

import (
	"github.com/ToteEmmanuel/academy-go-q12021/interface/controller"
	csv_file_reader "github.com/ToteEmmanuel/academy-go-q12021/tools/reader"
)

type registry struct {
	storage *csv_file_reader.CsvPokeStorage
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(csvPokeStorage *csv_file_reader.CsvPokeStorage) Registry {
	return &registry{csvPokeStorage}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokeController()
}
