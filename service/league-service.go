package service

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/jguerra6/API/infrastructure/datastore"
	"github.com/jguerra6/API/infrastructure/router"
)

type LeagueService interface {
	ValidateLeague(league map[string]interface{}) error
	GetAllLeagues() ([]map[string]interface{}, error)
	Addleague(league map[string]interface{}) (map[string]interface{}, error)
	GetLeague(id string) (map[string]interface{}, error)
	DeleteLeague(id string) error
	Updateleague(id string, league map[string]interface{}) (map[string]interface{}, error)
	GetID(request *http.Request) string
}

type leagueService struct{}

var (
	leagueDatabase datastore.Database
	leagueRouter   router.Router
)

func NewLeagueService(ld datastore.Database, lr router.Router) LeagueService {
	leagueDatabase = ld
	leagueRouter = lr
	return &leagueService{}
}

func (*leagueService) ValidateLeague(league map[string]interface{}) error {
	if league == nil {
		err := errors.New("The league is empty")
		return err
	}

	if league["name"] == "" {
		err := errors.New("The league name can't be empty")
		return err
	}

	//Validate that it's a float parse it to int
	if reflect.TypeOf(league["current_season_id"]) != reflect.TypeOf(1.0) {
		err := errors.New("The current_season_id must be an integer")
		return err
	}
	league["current_season_id"] = int(league["current_season_id"].(float64))
	return nil
}

func (*leagueService) GetAllLeagues() ([]map[string]interface{}, error) {
	return leagueDatabase.GetAll("leagues")
}

func (*leagueService) Addleague(league map[string]interface{}) (map[string]interface{}, error) {
	return leagueDatabase.AddItem("leagues", league)
}

func (*leagueService) GetLeague(id string) (map[string]interface{}, error) {

	return leagueDatabase.GetItemByID("leagues", id)
}

func (*leagueService) DeleteLeague(id string) error {
	return leagueDatabase.DeleteItem("leagues", id)
}

func (*leagueService) Updateleague(id string, league map[string]interface{}) (map[string]interface{}, error) {
	return leagueDatabase.UpdateItem("leagues", id, league)
}

func (*leagueService) GetID(request *http.Request) string {
	vars := leagueRouter.GetVarsFromRequest(request)
	return vars["id"]

}
