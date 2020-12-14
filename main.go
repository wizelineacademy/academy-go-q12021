// package main, core program
// Author: Rubén Vázquez
package main

import (
	"log"

	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/infrastructure/controller"
	"golang-bootcamp-2020/infrastructure/router"
	"golang-bootcamp-2020/infrastructure/services"
	"golang-bootcamp-2020/interface/usecase"
)

// main function inject dependencies
func main() {
	// load environment settings for environment
	err := config.ReadConfig("config")
	if err != nil {
		log.Fatal(err)
	}

	// Load client *service.Client
	s := services.NewClient()

	// load Usecase with client and services
	u := usecase.NewUsecase(s)

	// load controller using usecases
	c := controller.NewController(u)

	// router using controller
	router.NewRouter(c)
}
