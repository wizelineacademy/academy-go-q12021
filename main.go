package main

import (
	"fmt"
	"strconv"

	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/infrastructure/controller"
	"golang-bootcamp-2020/infrastructure/router"
	"golang-bootcamp-2020/infrastructure/services"
	"golang-bootcamp-2020/interface/usecase"
)

func main() {
	fmt.Println("Starting the app ...")
	config.ReadConfig()
	//db, ctx := datastore.NewDb()
	//defer db.Disconnect(ctx)
	//fmt.Println("DB ready")

	//resty en main
	fmt.Println("Server listen at " + config.C.Server.Address + ":" + strconv.Itoa(config.C.Server.Port))

	s := services.NewClient()
	u := usecase.New(s)
	c := controller.New(u)
	router.NewRouter(c)
}
