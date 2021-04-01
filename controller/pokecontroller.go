package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ToteEmmanuel/academy-go-q12021/usecase/interactor"
	"github.com/gorilla/mux"
)

type pokeController struct {
	pokeInteractor interactor.PokeInteractor
}

type PokeController interface {
	GetPokemon(http.ResponseWriter, *http.Request)
	GetPokemons(http.ResponseWriter, *http.Request)
	CatchPokemon(http.ResponseWriter, *http.Request)
	GetPokemonsWithWorkers(http.ResponseWriter, *http.Request)
}

func NewPokeController(pokeInteractor interactor.PokeInteractor) PokeController {
	return &pokeController{pokeInteractor}
}

func (pI *pokeController) GetPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p, err := pI.pokeInteractor.Get(int32(id))
	w.Header().Add("Content-Type", "Applicaiton/Json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error:":err.Error()})
		return
	}
	if err = json.NewEncoder(w).Encode(p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error:":err.Error()})
		log.Fatal("Error encoding... ", err)
	}
}

func (pI *pokeController) GetPokemons(w http.ResponseWriter, r *http.Request) {
	p, err := pI.pokeInteractor.GetAll()
	w.Header().Add("Content-Type", "Applicaiton/Json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error:":err.Error()})
		return
	}
	if err = json.NewEncoder(w).Encode(p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error:":err.Error()})
		log.Fatal("Error encoding... ", err)
	}
}

func (pI *pokeController) CatchPokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p, err := pI.pokeInteractor.CatchOne(int32(id))
	w.Header().Add("Content-Type", "Applicaiton/Json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error:":err.Error()})
		return
	}
	if err = json.NewEncoder(w).Encode(p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error:":err.Error()})
		log.Fatal("Error encoding... ", err)
	}
}

func (pI *pokeController) GetPokemonsWithWorkers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeQuery := vars["type"]
	if ! regexp.MustCompile(`(\bodd\b|\beven\b)`).MatchString(typeQuery) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error:":"invalid or missing [type]"})
		return
	}
	itemsQuery := vars["items"]
	var items int
	var err error
	if items, err = strconv.Atoi(itemsQuery); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error:":"invalid or missing [items]"})
		return
	}
	itemsPerWorkerQuery := vars["items_per_workers"]
	var itemsPerWorker int
	if itemsPerWorker, err = strconv.Atoi(itemsPerWorkerQuery); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error:":"invalid or missing [items_per_workers]"})
		return
	}

	pokemons, err := pI.pokeInteractor.GetAllWorkers(typeQuery, items, itemsPerWorker)

	w.Header().Add("Content-Type", "Applicaiton/Json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error:":err.Error()})
		return
	}
	if err = json.NewEncoder(w).Encode(pokemons); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error:":err.Error()})
		log.Fatal("Error encoding... ", err)
	}
}