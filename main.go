package main

import (
	"fmt"
	"log"

	"github.com/ramrodo/golang-bootcamp-2020/config"
	"github.com/ramrodo/golang-bootcamp-2020/usecase/repository"
	// "github.com/ramrodo/golang-bootcamp-2020/usecase/repository"
)

func main() {
	config.ReadConfig()
	// fmt.Printf("Server listen at %s:%s\n", config.C.Server.URL, config.C.Server.Port)

	// This returns error: "panic: runtime error: invalid memory address or nil pointer dereference"
	// var interactor interactor.FilmInteractor
	// repository := interactor.FilmRepository
	// controller := FilmController
	// films, err := repository.FindAll()
	// 	films, err := interactor.Get()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// fmt.Println("films:", films)

	all, err := repository.FindAll()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Title:", all[0].Title)
	fmt.Println("Description:", all[0].Description)
}
