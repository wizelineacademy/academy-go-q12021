package main

import (
	"encoding/json"
	"fmt"
	"go-api/config"
	"go-api/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TODO:
// Implement MVC pattern (refactor handlers into controllers)
// Implement repository pattern to support more datasources
func main() {
	config, err := config.New("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	cardBacks, err := data.InitializeCSVData(config.DataSources.CSV)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello, here's the data: %v", cardBacks)
	})

	router.HandleFunc("/cardbacks", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		res.WriteHeader(http.StatusOK)
		err := json.NewEncoder(res).Encode(cardBacks)
		if err != nil {
			fmt.Printf(err.Error())
			http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	})

	router.HandleFunc("/cardbacks/{id}", func(res http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		key, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Printf(err.Error())
			http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// TODO: Future... extract implementation details pass repository interface instead...
		for _, cardBack := range cardBacks {
			if cardBack.ID == key {
				res.Header().Set("Content-Type", "application/json; charset=utf-8")
				res.WriteHeader(http.StatusOK)
				err = json.NewEncoder(res).Encode(cardBack)
				if err != nil {
					fmt.Printf(err.Error())
					http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
			}
		}
	})

	port := config.Port
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("APP is listening on port: %d\n", port)
	log.Fatal(http.ListenAndServe(addr, router))
}
