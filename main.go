package main

import (
	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/infrastructure/controller"
	"golang-bootcamp-2020/infrastructure/router"
	"golang-bootcamp-2020/infrastructure/services"
	"golang-bootcamp-2020/interface/usecase"
)

func main() {
	// load environment settings
	config.ReadConfig()

	// Load client *service.Client
	s := services.NewClient()

	// load Usecase with client and services
	u := usecase.NewUsecase(s)

	// load controller using usecases
	c := controller.NewController(u)

	// router using controller
	router.NewRouter(c)
}
