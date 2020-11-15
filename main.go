package main

import (
	"golang-bootcamp-2020/controllers"
	"golang-bootcamp-2020/infrastructure/router"
)

func main() {

	if err := router.StartRouter(controllers.NewAppController()); err != nil {
		panic(err)
	}
}
