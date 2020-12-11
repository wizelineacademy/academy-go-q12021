package main

import (
	"fmt"
	"log"
	"net/http"

	"./routes"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // here
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service Status: Ok")
}

func main() {
	// fmt.Println("Start service")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/bitcoin", routes.GetBicoins).Methods("GET")
	router.HandleFunc("/bitcoin", routes.CreateBitcoinFromAPI).Methods("POST")
	router.HandleFunc("/bitcoin/manual", routes.CreateBitcoin).Methods("POST")
	router.HandleFunc("/bitcoin/{id}", routes.GetBicoin).Methods("GET")
	router.HandleFunc("/bitcoin/{id}", routes.DeleteBitcoin).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
