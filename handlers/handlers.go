package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	pokemonMap := make(map[int]string)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	file, err := os.Open("assets/pokemon.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, pokemon := range records {
		index, _ := strconv.Atoi(pokemon[0])
		pokemonMap[index] = pokemon[1]
	}
	response, _ := json.Marshal(pokemonMap)
	w.Write(response)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	pokemonMap := make(map[string]string)
	pokemonId := pathParams["pokemonId"]
	w.Header().Set("Content-Type", "application/json")

	file, err := os.Open("assets/pokemon.csv")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("An error occured while opening the file"))
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, pokemon := range records {
		pokemonMap[pokemon[0]] = pokemon[1]
	}

	if _, ok := pokemonMap[pokemonId]; ok {
		response, _ := json.Marshal(pokemonMap[pokemonId])
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Resource Not Found"))
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func Params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	fmt.Println(pathParams)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}
