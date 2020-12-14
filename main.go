package main

/// This software is a HTTP API to retrieve song lyrics. It connects to external Happi API
/// to get the information
import (
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/alexis-aguirre/golang-bootcamp-2020/config"
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/datastore"
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/router"
	r "github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/routes"
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/services"
)

func main() {
	Config := &config.C
	Config.ReadConfig()
	startServices()
	router := router.NewRouter()
	r.AddRoutes(router)

	server := &http.Server{
		Addr:           ":" + Config.PORT,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening in port: " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

//startServices starts the system services and prepares it to dependency injection
func startServices() {
	registry := services.NewServiceRegistry()

	db := datastore.InitializeDB()
	registry.RegisterService(services.DATABASE, db)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot get current directory")
	}
	logger := datastore.InitializeLogger(path.Join(currentDir, "logfile.csv"))
	registry.RegisterService(services.LOGGER, logger)

}
