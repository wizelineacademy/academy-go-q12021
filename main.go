package main

import (
  "log"
  "github.com/andrecrts/go-bootcamp/infrastructure/router"
	"github.com/labstack/echo/v4"
)

func handleRequests() {

  e := echo.New()
  e = router.NewRouter(e)

	if err := e.Start(":1000"); err != nil {
		log.Fatalln(err)
	}
}

func main() {
  handleRequests()
}
