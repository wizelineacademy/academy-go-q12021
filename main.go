package main

import (
	"flag"
	"fmt"
	configz "github.com/ToteEmmanuel/academy-go-q12021/config"
	"github.com/ToteEmmanuel/academy-go-q12021/controller"
	"github.com/ToteEmmanuel/academy-go-q12021/infrastructure/datastore"
	router2 "github.com/ToteEmmanuel/academy-go-q12021/infrastructure/router"
	usecaseinteractor "github.com/ToteEmmanuel/academy-go-q12021/usecase/interactor"
	usecaserepository "github.com/ToteEmmanuel/academy-go-q12021/usecase/repository"
	"github.com/go-resty/resty/v2"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {

	filePath := flag.String("file", "./config.yml", "The config files path.")
	flag.Parse()

	config, err := configz.Load(filePath)
	if err != nil {
		log.Fatalf("Could not find config file %s %v", *filePath, err)
	}
	csvPokeStorage, err := datastore.NewCsvStorage(config.CsvName)
	if err != nil {
		log.Fatalf("No Source for Pokemons!(file not found %s) %v\n", config.CsvName, err)
		panic(err)
	}
	restClient := resty.New()
	negi := negroni.Classic()
	router := router2.NewRouter(initPokeDependencies(&csvPokeStorage, restClient))
	negi.UseHandler(router)
	if err := http.ListenAndServe(config.AppHost, negi); err != nil {
		fmt.Printf("Sudden death %v\n", err)
	}
}

func initPokeDependencies(storage *datastore.PokeStorage, client *resty.Client) controller.PokeController {
	pokeRepository := usecaserepository.NewPokeRepository(*storage)
	pokeInteractor := usecaseinteractor.NewPokeInteractor(pokeRepository, client, usecaseinteractor.NewInfoClient())
	return controller.NewPokeController(pokeInteractor)
}
