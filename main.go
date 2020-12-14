package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jguerra6/API/controller"
	"github.com/jguerra6/API/infrastructure/datastore"
	"github.com/jguerra6/API/infrastructure/router"
	"github.com/jguerra6/API/service"
)

var (
	database         datastore.Database          = datastore.NewFirestoreDB()
	httpRouter       router.Router               = router.NewMuxRouter()
	leagueService    service.LeagueService       = service.NewLeagueService(database, httpRouter)
	leagueController controller.LeagueController = controller.NewLeagueController(leagueService)
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
