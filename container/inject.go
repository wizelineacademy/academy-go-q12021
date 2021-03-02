package container

import (
	"fmt"

	"github.com/golobby/container"
	"github.com/oscarSantoyo/academy-go-q12021/service"
)

func Connect ( ){
	fmt.Println("Connecting container")
	// var instance = container.NewContainer();
	container.Singleton(func () service.Search {
		return &service.SearchImpl{}
	})
	container.Singleton(func () service.CsvService {
		return &service.CsvServiceImpl{}
	})
}
