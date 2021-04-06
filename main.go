package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joseantoniovz/academy-go-q12021/router"
)

func main() {
	r := mux.NewRouter()

	router.SetRoutes(r)

	srv := http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	log.Println("Running on port 8081")
	log.Println(srv.ListenAndServe())
}
