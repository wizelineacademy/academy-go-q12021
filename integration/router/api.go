package router

import (
	"fmt"
	"net/http"

	"github.com/etyberick/golang-bootcamp-2020/entity"
	"github.com/etyberick/golang-bootcamp-2020/service/handler"
	"github.com/etyberick/golang-bootcamp-2020/usecase/interactor"
	"github.com/gorilla/mux"
)

const quoteRoute = "/v0/quote"

// Initialize a new web server
func Initialize(config entity.Config) {
	// Initialize the quote handler
	qi := interactor.NewQuoteInteractor(config)
	quoteHandler := handler.NewQuote(qi)

	// Initialize the router and its endpoints
	r := mux.NewRouter()
	r.HandleFunc(quoteRoute, quoteHandler.GetAll).Methods("GET")
	r.HandleFunc(quoteRoute, quoteHandler.Update).Methods("POST")

	//Serve
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r)
	return
}
