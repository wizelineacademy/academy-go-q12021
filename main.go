package main

import (
	"fmt"
	"github.com/ToteEmmanuel/academy-go-q12021/controller"
	router2 "github.com/ToteEmmanuel/academy-go-q12021/infrastructure/router"
	cvsfilereader "github.com/ToteEmmanuel/academy-go-q12021/tools/reader"
	usecaseinteractor "github.com/ToteEmmanuel/academy-go-q12021/usecase/interactor"
	usecasepresenter "github.com/ToteEmmanuel/academy-go-q12021/usecase/presenter"
	usecaserepository "github.com/ToteEmmanuel/academy-go-q12021/usecase/repository"
	"github.com/go-resty/resty/v2"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	csvPokeStorage, err := cvsfilereader.NewCsvStorage("./poke.csv")
	restClient := resty.New()
	if err != nil {
		fmt.Println("No Source for Pokemons!")
		panic(err)
	}
	negi := negroni.Classic()
	router := router2.NewRouter(initPokeDependencies(csvPokeStorage, restClient))
	negi.UseHandler(router)
	if err := http.ListenAndServe("localhost:3000", negi); err != nil {
		fmt.Printf("Sudden death %v\n", err)
	}
}

func initPokeDependencies(storage *cvsfilereader.CsvPokeStorage, client *resty.Client) controller.PokeController {
	pokePresenter := usecasepresenter.NewPokePresenter()
	pokeRepository := usecaserepository.NewPokeRepository(storage)
	pokeInteractor := usecaseinteractor.NewPokeInteractor(pokeRepository, pokePresenter, client)
	return controller.NewPokeController(pokeInteractor)
}
