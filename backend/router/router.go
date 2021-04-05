package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	GetMovies(w http.ResponseWriter, r *http.Request)
	GetMovieById(w http.ResponseWriter, r *http.Request)
	GetMoviesConcurrently(w http.ResponseWriter, r *http.Request)
	GetLanguages(w http.ResponseWriter, r *http.Request)
	GetLanguageById(w http.ResponseWriter, r *http.Request)
}

func New(c Controller) *mux.Router {
	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/getMovies", c.GetMovies).Methods("GET")
	r.HandleFunc("/getMoviesConcurrently", c.GetMoviesConcurrently).Methods("GET")
	r.HandleFunc("/getMovieById", c.GetMovieById).Methods("GET")
	r.HandleFunc("/getTechStack", c.GetLanguages).Methods("GET")
	r.HandleFunc("/getTechStackById", c.GetLanguageById).Methods("GET")
	return r
}
