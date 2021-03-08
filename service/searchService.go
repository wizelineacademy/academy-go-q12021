package service

import (
	"github.com/oscarSantoyo/academy-go-q12021/model"

	"github.com/golobby/container"
)

// Search interface of sevice
type Search interface {
	Search(string) ([]model.Doc, error)
}

// SearchImpl implementation of Search service
type SearchImpl struct{}

// Search return the data filtered by
func (s SearchImpl) Search(term string) ([]model.Doc, error) {
	var csvService CsvService
	container.Make(&csvService)
	return csvService.FilterByID(term)
}
