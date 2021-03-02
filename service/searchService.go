package service

import (
	"fmt"
	"github.com/golobby/container"
)

var csvService CsvService

type Search interface {
	Search(string) ([]Doc, error)
}

type SearchImpl struct{}

func init() {
	fmt.Println("ignited")
}

func (s SearchImpl) Search(term string) ([]Doc, error) {
	return getCsvService().FilterById(term)
}

func getCsvService() CsvService {
	if csvService == nil {
		container.Make(&csvService)
	}
	return csvService
}
