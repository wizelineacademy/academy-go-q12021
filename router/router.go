package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joseantoniovz/academy-go-q12021/controller"
)

func SetRoutes(r *mux.Router) {
	r.HandleFunc("/book", controller.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/book/{id}", controller.GetBookById).Methods(http.MethodGet)
	r.HandleFunc("/books/consume", controller.ConsumeAPI).Methods(http.MethodGet)
}
