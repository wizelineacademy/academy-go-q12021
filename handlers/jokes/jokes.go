package jokes

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Joke struct {
	ID        int    `json:id`
	Setup     string `json:setup`
	Punchline string `json:punchline`
}

// type allJokes []Joke
var allJokes []Joke

//Load data from CSV
func Load() {
	// fmt.Println("This is the current working directory")
	// fmt.Println(os.Getwd())
	pwd, _ := os.Getwd()
	csvFile, err := os.Open(pwd + "/data/data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, rec := range records {
		var insert Joke
		insert.ID, err = strconv.Atoi(rec[0])
		insert.Setup = rec[1]
		insert.Punchline = rec[2]
		allJokes = append(allJokes, insert)
	}

}

//GetJokes all Jokes in the data
func GetJokes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allJokes)
}

//GetOneJoke only one joke by ID
func GetOneJoke(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jokeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	if len(allJokes) < jokeID {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "This joke doesn't exists")
	}
	for _, joke := range allJokes {
		if joke.ID == jokeID {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(joke)
		}
	}
}

//IndexRoute message
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Jokes API...")

}
