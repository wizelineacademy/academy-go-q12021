package httputil

import (
	"net/http"

	muxhandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// NewGorillaRouter allocates a new gorilla.mux pre-configured router
func NewGorillaRouter() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(false)
	injectMiddlewares(r)
	return r
}

func injectMiddlewares(r *mux.Router) {
	r.Use(muxhandlers.RecoveryHandler())
	r.Use(muxhandlers.CORS(
		muxhandlers.AllowedMethods([]string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		}),
		muxhandlers.AllowedOrigins([]string{"*"}),
	))
	r.Use(muxhandlers.CompressHandler)
}
