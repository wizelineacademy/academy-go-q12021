package jokes

import (
	"academy/services/dataload"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetJokes all Jokes in the data
func GetJokes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataload.ReadData())
}

//GetOneJoke only one joke by ID
func GetOneJoke(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jokeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	if len(dataload.ReadData()) < jokeID {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "This joke doesn't exists")
	}

	for _, joke := range dataload.ReadData() {
		if joke.ID == jokeID {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(joke)
		}
	}
}
