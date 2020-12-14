package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wizelineacademy/golang-bootcamp-2020/domain/models"
	"github.com/wizelineacademy/golang-bootcamp-2020/domain/repositories"
	"github.com/wizelineacademy/golang-bootcamp-2020/helpers"
)

const CharactersEndpoint = "https://rickandmortyapi.com/api/character/"

type RickMortyHandler struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	client   *http.Client
	charRepo repositories.CharacterRepository
}

// NewRickMortyHandler returns a RickMortyHandler struct with a logger.
func NewRickMortyHandler(infoLog, errorLog *log.Logger, c *http.Client, charRepo repositories.CharacterRepository) *RickMortyHandler {
	return &RickMortyHandler{infoLog, errorLog, c, charRepo}
}

// InsertCharacterByID returns information about a Rick and Morty character and inserts it into a local CSV file
func (rm *RickMortyHandler) InsertCharacterByID(w http.ResponseWriter, r *http.Request) {
	// Get the route variables
	vars := mux.Vars(r)
	id := vars["id"]
	// Complete the endpoint with the character id
	url := CharactersEndpoint + id

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		rm.errorLog.Printf("msg: error getting character, err: %v\n", err)
		helpers.ServerError(w, rm.errorLog, err)
		return
	}
	// Send the HTTP request
	resp, err := rm.client.Do(req)
	if err != nil {
		rm.errorLog.Printf("msg: error sending http request, err: %v\n", err)
		helpers.ServerError(w, rm.errorLog, err)
		return
	}
	// Close the body at the end of execution
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rm.errorLog.Printf("msg: error while reading the body, err: %v\n", err)
		helpers.ServerError(w, rm.errorLog, err)
		return
	}

	// Empty struct to hold the response data
	char := &models.Character{}

	err = json.Unmarshal(body, char)
	if err != nil {
		rm.errorLog.Printf("msg: error while unmarshaling character data, err: %v\n", err)
		helpers.ServerError(w, rm.errorLog, err)
		return
	}

	jsonChar, err := json.Marshal(char)
	if err != nil {
		rm.errorLog.Printf("msg: error while marshaling character data, err: %v\n", err)
		helpers.ServerError(w, rm.errorLog, err)
		return
	}

	// Insert to CSV
	err = rm.charRepo.Insert(char)
	if err != nil {
		rm.errorLog.Printf("msg: error while inserting character data, err: %v\n", err)
		helpers.ServerError(w, rm.errorLog, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v\n", string(jsonChar))

}

// GetCharacterByID returns information about a Rick and Morty character read from a local CSV file
func (rm *RickMortyHandler) GetCharacterByID(w http.ResponseWriter, r *http.Request) {
	// Get the route variables
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rm.errorLog.Printf("msg: error while casting character id, err: %v\n", err)
		helpers.ServerError(w, rm.errorLog, err)
		return
	}

	character, err := rm.charRepo.Get(id)
	if err != nil {
		if errors.Is(err, repositories.ErrNoChar) {
			fmt.Fprintf(w, "Character with id %v wasn't found in local CSV file \n", id)
			return
		} else {
			rm.errorLog.Printf("msg: error while obtaining character %v data, err: %v\n", id, err)
			helpers.ServerError(w, rm.errorLog, err)
			return
		}

	}

	jsonChar, err := json.Marshal(character)
	if err != nil {
		rm.errorLog.Printf("msg: error while marshaling character data, err: %v\n", err)
		helpers.ServerError(w, rm.errorLog, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v\n", string(jsonChar))

}
