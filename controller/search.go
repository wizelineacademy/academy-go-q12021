package controller

import (
	"github.com/oscarSantoyo/academy-go-q12021/model"
	"github.com/oscarSantoyo/academy-go-q12021/service"

	"github.com/golobby/container"
	"github.com/labstack/gommon/log"
)

var search service.Search

// SearchByID returns data read from file filtered by ID
func SearchByID(id string) []model.Doc {
	result, err := getSearchService().Search(id)
	if err != nil {
		log.Info("There was no records present")
	}
	return result
}

func getSearchService() service.Search {
	if search == nil {
		log.Info("wiring search service")
		container.Make(&search)
	}
	return search
}
