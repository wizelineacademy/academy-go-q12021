package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"example.com/me/csv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("{status: OK}")
	fmt.Println("Endpoint Hit: homePage")
}

func csvRetrieveAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(csv.RetrieveFromCSV(""))
	fmt.Println("Endpoint Hit: csvRetrieveAll")
}

func csvSingleItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(csv.RetrieveFromCSV(id))
	fmt.Println("Endpoint Hit: csvSingleItem")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/allitems", csvRetrieveAll).Methods("GET")
	router.HandleFunc("/item/{id}", csvSingleItem).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
}
