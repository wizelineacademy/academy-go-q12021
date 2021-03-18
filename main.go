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
	config.ReadConfig()


	db := datastore.NewDB()

	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost:"+config.C.Server.Address)
	if err := e.Start(":"+config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
	/* */
}
