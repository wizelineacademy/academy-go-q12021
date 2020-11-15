package app

import (
	"golang-bootcamp-2020/controllers"
	"golang-bootcamp-2020/infrastructure/router"
	"golang-bootcamp-2020/repository/db"
	"golang-bootcamp-2020/repository/rest"
	"golang-bootcamp-2020/services"
)

func StartApp() {

	handler := controllers.NewAppController(services.NewService(rest.NewRickAndMortyApiRepository(), db.NewDbRepository()))

	if err := router.StartRouter(handler); err != nil {
		panic(err)
	}
}
