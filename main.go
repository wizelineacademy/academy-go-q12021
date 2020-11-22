package main

import (
	"fmt"
	"log"

	"github.com/ramrodo/golang-bootcamp-2020/config"
	"github.com/ramrodo/golang-bootcamp-2020/usecase/repository"
)

func main() {
	config.ReadConfig()
	// fmt.Printf("Server listen at %s:%s\n", config.C.Server.URL, config.C.Server.Port)

	all, err := repository.FindAll()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Title:", all[0].Title)
	fmt.Println("Description:", all[0].Description)
}
