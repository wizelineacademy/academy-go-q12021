package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/etyberick/golang-bootcamp-2020/entity"
	"github.com/gorilla/mux"
)

//ResponseError represents error output feedback
type ResponseError struct {
	Message string `json:"message"`
}

//QuoteHandler is the http handler for entity.QuoteManager
type QuoteHandler struct {
	QuoteStorage entity.QuoteStorage
}

//NewQuoteHandler initializes /v0/quote endpoint
func NewQuoteHandler(r *mux.Router, quoteStorage entity.QuoteStorage) {
	handler := &QuoteHandler{
		QuoteStorage: quoteStorage,
	}

	r.HandleFunc("/quote", handler.UpdateQuotes).Methods("POST")
	r.HandleFunc("/quote", handler.GetQuotes).Methods("GET")
	http.ListenAndServe(":8080", r)
}

//UpdateQuotes will update the quotes repository on each call
func (q *QuoteHandler) UpdateQuotes(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://quote-garden.herokuapp.com/api/v2/quotes/random")
	if err != nil {
		log.Printf("error fetching quote from third party - %v", err)
	}
	defer resp.Body.Close()

	//convert response to JSON
	var quote entity.QuoteGarden
	responseMessage, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(responseMessage, &quote)
	w.Write(responseMessage)
	q.QuoteStorage.Write(&quote.Quote)
}

//GetQuotes will return all the available quotes on each call
func (q *QuoteHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	panic("UpdateQuotes not implemented") // TODO: Implement
}
