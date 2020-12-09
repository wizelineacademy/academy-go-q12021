package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wizelineacademy/golang-bootcamp-2020/domain/models"
	"github.com/wizelineacademy/golang-bootcamp-2020/domain/repositories"
	"github.com/wizelineacademy/golang-bootcamp-2020/helpers"
)

// ChampionHandler defines the Champion Handler properties
type ChampionHandler struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

// NewChampionHandler returns a ChampModel struct with a logger.
func NewChampionHandler(infoLog, errorLog *log.Logger) *ChampionHandler {
	return &ChampionHandler{infoLog, errorLog}
}

// GetChamp returns a single Champion by id
func (ch *ChampionHandler) GetChamp(w http.ResponseWriter, r *http.Request) {
	// Get the Gorilla mux vars (path params)
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Create a new chap repo to access the db
	cr, err := repositories.NewChampRepo(ch.errorLog)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		ch.errorLog.Printf("message: error while creating a new champ repo, handler: champions, method: GET, id: %v, err: %v\n", id, err)
	}
	defer cr.DB.Close()

	// Query the database to get a single Champion
	c, err := cr.GetSingle(id)
	if err != nil {
		if errors.Is(err, repositories.ErrNoRecord) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			helpers.ServerError(w, ch.errorLog, err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Marshal the Champion struct to JSON
	err = c.ToJSON(w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		ch.errorLog.Printf("message: error while parsing the champions data, handler: champions, method: GET, id: %v, err: %v\n", id, err)
	}
}

// GetChamps returns multiple Champions. It defaults to 10 if a limit isn't specified or if the limit is < 1.
func (ch *ChampionHandler) GetChamps(w http.ResponseWriter, r *http.Request) {
	const defaultLimit = 10

	limit := defaultLimit
	params := r.URL.Query()
	// If query param "limit" is found
	if params.Get("limit") != "" {
		paramLimit, err := strconv.Atoi(params.Get("limit"))
		if err != nil {
			ch.errorLog.Printf("message: error parsing limit, handler: champions, method: GET, err: %v\n", err)
			helpers.NotFound(w)
			return
		}
		if paramLimit < 1 {
			ch.infoLog.Printf("message: limit < 1, using default limit %v, handler: champions, method: GET", defaultLimit)
		} else {
			limit = paramLimit
		}
	}

	// Create a new chap repo to access the db
	cr, err := repositories.NewChampRepo(ch.errorLog)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		ch.errorLog.Printf("message: error while creating a new champ repo, handler: champions, method: GET, err: %v\n", err)
	}
	defer cr.DB.Close()

	// Query the database and get a slice of Champions
	champions, err := cr.GetMultiple(limit)
	if err != nil {
		helpers.ServerError(w, ch.errorLog, err)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")

	// Marshal the Champions slice to JSON
	for _, c := range champions {
		err = c.ToJSON(w)
		if err != nil {
			ch.errorLog.Printf("message: error while parsing the champions data, handler: champions, method: GET, err: %v\n", err)
			helpers.ServerError(w, ch.errorLog, err)
		}
	}
}

// AddChamp inserts a new Champ into the DB
func (ch *ChampionHandler) AddChamp(w http.ResponseWriter, r *http.Request) {
	// Initialize a pointer to a new zeroed Champion struct.
	champ := &models.Champion{}

	err := champ.FromJSON(r.Body)
	if err != nil {
		ch.errorLog.Printf("message: error while decoding the JSON data, handler: champions, method: POST, err: %v\n", err)
		helpers.ServerError(w, ch.errorLog, err)
		return
	}

	// Create a new chap repo to access the db
	cr, err := repositories.NewChampRepo(ch.errorLog)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		ch.errorLog.Printf("message: error while creating a new champ repo, handler: champions, method: GET, err: %v\n", err)
	}
	defer cr.DB.Close()
	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := cr.Insert(champ)
	if err != nil {
		fmt.Printf("err %v\n", err)
		helpers.ServerError(w, ch.errorLog, err)
		return
	}

	fmt.Fprintf(w, "Success creating the Champion: %v\n", id)
}
