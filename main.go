package main

import (
	"context"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pankecho/golang-bootcamp-2020/store"
	"github.com/pankecho/golang-bootcamp-2020/transport"
	"github.com/pankecho/golang-bootcamp-2020/usecase"
)

func main() {
	ctx := context.Background()
	personStore, _ := store.NewStore(ctx)
	personUseCase := usecase.NewPerson(personStore)
	personTransport := transport.NewPerson(personUseCase)

	e := transport.NewRouter(personTransport)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}
