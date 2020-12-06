package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/jguerra6/API/errorsHandler"

	"github.com/jguerra6/API/infrastructure/datastore"
	"github.com/jguerra6/API/infrastructure/router"
)

type leagueController struct {
	router.Router
	datastore.Database
}

//LeagueController will create an interface to control all the League Operations
type LeagueController interface {
	GetAllLeagues(writer http.ResponseWriter, request *http.Request)
	Addleague(writer http.ResponseWriter, request *http.Request)
	GetLeague(writer http.ResponseWriter, request *http.Request)
	DeleteLeague(writer http.ResponseWriter, request *http.Request)
	Updateleague(writer http.ResponseWriter, request *http.Request)
}

//NewLeagueController returns a League Controller to handle the League Operations
func NewLeagueController(db datastore.Database, router router.Router) LeagueController {
	return &leagueController{router, db}
}

//GetAllLeagues returns all the items in the "leagues" table
func (lc *leagueController) GetAllLeagues(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	leagues, err := lc.GetAll("leagues")

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Error getting the leagues"})
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(leagues)
}

//validateLeague will validate that if item passed along it's valid or not
func validateLeague(league map[string]interface{}) error {

	if league == nil {
		err := errors.New("The league is empty")
		return err
	}

	if league["name"] == "" || league["name"] == nil {
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

func (lc *leagueController) Addleague(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	//Save the json body of the request in a map and handle the errors if any.
	tmpLeague := make(map[string]interface{})
	err := json.NewDecoder(request.Body).Decode(&tmpLeague)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Your JSON seems to be invalid. Please check again your input."})
		//log.Println("Failed decoding item: ", err)
		return
	}

	//Store the user input into a defined interface, this to remove any extra stuff the user might send
	league := map[string]interface{}{
		"name":              tmpLeague["name"],
		"country":           tmpLeague["country"],
		"current_season_id": tmpLeague["current_season_id"],
	}

	//Validate that the user input has the correct fields and types, otherwise handle the error and return a response to the user
	err1 := validateLeague(league)
	if err1 != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: err1.Error()})
		return
	}

	//Add the item to the database and handle the errors if any
	result, err2 := lc.AddItem("leagues", league)

	if err2 != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Error saving the league"})
		return
	}

	//Return the added object and a OK response
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(result)
}

//GetLeague will return the requested league
func (lc *leagueController) GetLeague(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	//Get the id from the request url
	vars := lc.GetVarsFromRequest(request)
	id := vars["id"]

	//Get the item from the DB and handle errors if any.
	league, err := lc.GetItemByID("leagues", id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "League not found"})
		return
	}

	//Return the item to the user
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(league)
}

//DeleteLeague will delete the specified item
func (lc *leagueController) DeleteLeague(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	//Get the id from the request url
	vars := lc.GetVarsFromRequest(request)
	id := vars["id"]

	//Perform the delete request and handle the error if any
	err := lc.DeleteItem("leagues", id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: err.Error()})
		fmt.Println("Error")
		return
	}

	//Return a success message.
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Succesfully deleted league"})
}

func (lc *leagueController) Updateleague(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	//Get the id from the request url
	vars := lc.GetVarsFromRequest(request)
	id := vars["id"]

	//Save the json body of the request in a map and handle the errors if any.
	tmpLeague := make(map[string]interface{})
	err := json.NewDecoder(request.Body).Decode(&tmpLeague)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: "Your JSON seems to be invalid. Please check again your input."})
		//log.Println("Failed decoding item: ", err)
		return
	}

	//Update the league, handle errors if any and get the new updated item
	var result map[string]interface{}
	result, err = lc.UpdateItem("leagues", id, tmpLeague)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorsHandler.ServiceError{Message: err.Error()})
		fmt.Println("Error")
		return
	}

	//Return the added object and a OK response
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(result)
}
