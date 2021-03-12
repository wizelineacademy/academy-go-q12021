package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"

	"github.com/AlejandroSeguraWIZ/academy-go-q12021/registry"
	"github.com/AlejandroSeguraWIZ/academy-go-q12021/router"
)

func main() {
	address := ":8081"
	r := registry.NewRegistry("pokemon.csv")

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + address)
	if err := e.Start(address); err != nil {
		log.Fatalln(err)
	}
}
