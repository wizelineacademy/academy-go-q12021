package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	GetMovies(w http.ResponseWriter, r *http.Request)
	GetMovieById(w http.ResponseWriter, r *http.Request)
	GetConcurrently(w http.ResponseWriter, r *http.Request)
}

func New(c Controller) *mux.Router {
	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/movies", c.GetMovies).Methods("GET")
	r.HandleFunc("/getMovies", c.GetConcurrently).Methods("GET")
	r.HandleFunc("/getMovieById", c.GetMovieById).Methods("GET")
	return r
}
