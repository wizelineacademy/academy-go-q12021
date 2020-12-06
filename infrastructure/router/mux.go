package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

//NewMuxRouter returns a gorilla mux router
func NewMuxRouter() Router {
	return &muxRouter{}
}

//GET will handle the requests and redirect them to the specified function
func (*muxRouter) GET(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

//POST will handle the requests and redirect them to the specified function
func (*muxRouter) POST(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

//PATCH will handle the requests and redirect them to the specified function
func (*muxRouter) PATCH(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("PATCH")
}

//DELETE will handle the requests and redirect them to the specified function
func (*muxRouter) DELETE(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")
}

//SERVE will start the Router with the specified port
func (*muxRouter) SERVE(port string) {
	log.Println("Listening to port ", port)
	http.ListenAndServe(port, muxDispatcher)
}

//GetVarsFromRequest will return the vars in the url
func (*muxRouter) GetVarsFromRequest(request *http.Request) map[string]string {
	vars := mux.Vars(request)
	return vars
}
