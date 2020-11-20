package main

import (
	"log"
	"net/http"
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
	startRegistries()
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

func startRegistries() {
	registry := services.NewServiceRegistry()

	db := datastore.InitializeDB()
	registry.RegisterService(db)

	registry.StartAll()
}
