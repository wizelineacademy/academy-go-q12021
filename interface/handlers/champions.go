package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wizelineacademy/golang-bootcamp-2020/domain/models"
	"github.com/wizelineacademy/golang-bootcamp-2020/helpers"
)

// ChampionHandler defines the Champion Handler properties
type ChampionHandler struct {
	infoLog   *log.Logger
	errorLog  *log.Logger
	champRepo models.ChampionRepository
}

// NewChampHandler returns a ChampModel struct with a logger.
func NewChampHandler(infoLog, errorLog *log.Logger, cr models.ChampionRepository) *ChampionHandler {
	return &ChampionHandler{infoLog, errorLog, cr}
}

func (ch *ChampionHandler) GetChamp(w http.ResponseWriter, r *http.Request) {
	// Get the Gorilla mux vars
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id < 1 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	c, err := ch.champRepo.GetSingle(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = c.ToJSON(w)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		ch.errorLog.Printf("message: error while parsing the champions data, handler: posts, method: GET, id: %v, err: %v\n", id, err)
	}
}

func (ch *ChampionHandler) GetChamps(w http.ResponseWriter, r *http.Request) {

	posts, err := ch.champRepo.GetMultiple()
	if err != nil {
		helpers.ServerError(w, ch.errorLog, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	for _, p := range posts {
		err = p.ToJSON(w)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			ch.errorLog.Printf("message: error while parsing the champions data, handler: posts, method: GET, err: %v\n", err)
		}
	}

}
