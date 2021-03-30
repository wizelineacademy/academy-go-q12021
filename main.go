package main

import (
	"bootcamp/config"
	"bootcamp/infrastructure/datastore"
	"bootcamp/infrastructure/router"
	"bootcamp/registry"
	"fmt"
	"github.com/labstack/echo"
	"log"
)


func main () {

	// Read confiig settings
	config.ReadConfig()

	// startup sql connection
	db := datastore.NewDB()

	defer db.Close()

	//
	r := registry.NewRegistry(db)

	// New instance of Echo router library
	e := echo.New()
	// setting up routing
	e = router.NewRouter(e, r.NewAppController())

	// start server on 8080 port
	fmt.Println("Server listen at http://localhost:"+config.C.Server.Address)
	if err := e.Start(":"+config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
