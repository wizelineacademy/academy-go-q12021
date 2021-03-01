package main

import (
	"fmt"
	"log"

	"github.com/Topi99/academy-go-q12021/infrastructure/router"
	"github.com/Topi99/academy-go-q12021/registry"
	"github.com/labstack/echo"
)

func main() {
	r := registry.NewRegistry()

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost:8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
