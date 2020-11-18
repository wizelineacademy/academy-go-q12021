package router

import (
	"github.com/gorilla/mux"
)

//NewRouter creates a new router instance
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}
