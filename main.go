package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jguerra6/API/controller"
	"github.com/jguerra6/API/infrastructure/datastore"
	"github.com/jguerra6/API/infrastructure/router"
)

var (
	httpRouter       = router.NewMuxRouter()
	db               = datastore.NewFirestoreDB()
	leagueController = controller.NewLeagueController(db, httpRouter)
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Homepage Endpoint Hit")
}

func main() {
	//Read the flag for the port so users can define their own
	port := flag.String("port", ":8081", "HTTP network address")
	flag.Parse()

	httpRouter.GET("/", homePage)
	httpRouter.GET("/leagues", leagueController.GetAllLeagues)
	httpRouter.POST("/leagues", leagueController.Addleague)
	httpRouter.GET("/leagues/{id}", leagueController.GetLeague)
	httpRouter.DELETE("/leagues/{id}", leagueController.DeleteLeague)
	httpRouter.PATCH("/leagues/{id}", leagueController.Updateleague)

	httpRouter.SERVE(*port)

}
