package modules

import (
	"model"
	"utils"
	"service/db"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"gopkg.in/mgo.v2/bson"
)

var collection = db.GetSession()

func responseOne(w http.ResponseWriter, poke model.Pokemon) {
	w = setHeaders(w)
	json.NewEncoder(w).Encode(poke)
}

func responseSome(w http.ResponseWriter, pokes []model.Pokemon) {
	w = setHeaders(w)
	json.NewEncoder(w).Encode(pokes)
}

func responseWithError(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
	return
}

func setHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
	return w
}

func GetPokemonCsv(w http.ResponseWriter, r *http.Request) {
	pokeList := utils.ReadCSV()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithError(w, http.StatusBadRequest)
	}

	pokemonId := id - 1

	if pokemonId <= len(pokeList) - 1 {
		responseOne(w, pokeList[pokemonId])
	} else {
		responseWithError(w, http.StatusNotFound)
	}
}

func GetPokemonListCsv(w http.ResponseWriter, r *http.Request) {
	pokeList := utils.ReadCSV()
	responseSome(w, pokeList)
}

func AddPokemon(w http.ResponseWriter, r *http.Request) {
	var data model.Pokemon
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	if err != nil {
		responseWithError(w, http.StatusNotFound)
	}

	defer r.Body.Close()

	err = collection.Insert(data)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError)
	} else {
		responseOne(w, data)
	}
}

func GetPokemonList(w http.ResponseWriter, r *http.Request) {
	var pokeList model.PokemonList

	err := collection.Find(nil).Sort("_id").All(&pokeList)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError)
	} else {
		responseSome(w, pokeList)
	}
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if !bson.IsObjectIdHex(id) {
		responseWithError(w, http.StatusBadRequest)
	}

	var poke model.Pokemon
	objectId := bson.ObjectIdHex(id)

	err := collection.FindId(objectId).One(&poke)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError)
	} else {
		responseOne(w, poke)
	}
}

func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if !bson.IsObjectIdHex(id) {
		responseWithError(w, http.StatusBadRequest)
	}

	var poke model.Pokemon
	decoder := json.NewDecoder(r.Body)
	objectId := bson.ObjectIdHex(id)

	err := decoder.Decode(&poke)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError)
	}
	
	defer r.Body.Close()

	document := bson.M{"_id": objectId}
	change := bson.M{"$set":poke}

	err = collection.Update(document, change)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError)
	} else {
		responseOne(w, poke)
	}
}

func DeletePokemon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if !bson.IsObjectIdHex(id) {
		responseWithError(w, http.StatusBadRequest)
	}

	var poke model.Pokemon
	objectId := bson.ObjectIdHex(id)

	err := collection.FindId(objectId).One(&poke)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError)
	}

	err = collection.RemoveId(objectId)

	if err != nil {
		responseWithError(w, http.StatusInternalServerError)
	} else {
		responseOne(w, poke)
	}
}