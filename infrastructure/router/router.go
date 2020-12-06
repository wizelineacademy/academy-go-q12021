package router

import (
	"net/http"
)

//Router is the object that will process and route the API requests
type Router interface {
	GET(uri string, f func(writer http.ResponseWriter, request *http.Request))
	POST(uri string, f func(writer http.ResponseWriter, request *http.Request))
	PATCH(uri string, f func(writer http.ResponseWriter, request *http.Request))
	DELETE(uri string, f func(writer http.ResponseWriter, request *http.Request))
	SERVE(port string)
	GetVarsFromRequest(request *http.Request) map[string]string
}
