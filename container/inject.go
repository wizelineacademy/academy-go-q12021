package container

import (
	"github.com/oscarSantoyo/academy-go-q12021/service"

	"github.com/golobby/container"
	"github.com/labstack/gommon/log"
)

// Connect wires container with services references
func Connect() {
	log.Info("Connecting container")
	container.Singleton(func() service.Search {
		return &service.SearchImpl{}
	})
	container.Singleton(func() service.CsvService {
		return &service.CsvServiceImpl{}
	})
	container.Singleton(func() service.ConfigService {
		return &service.ConfigImpl{}
	})
}
