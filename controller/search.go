package controller

import (
	"fmt"

	"github.com/golobby/container"
	"github.com/oscarSantoyo/academy-go-q12021/service"
)

var search service.Search

func SearchById(id string) []service.Doc{
	result, err := getSearchService().Search(id)
	if err != nil {
		fmt.Println("There was no records present")
	}
	return result
}

func getSearchService() service.Search {
	if (search == nil) {
		fmt.Println("wiring search service")
		container.Make(&search)
	}
	return search
}
