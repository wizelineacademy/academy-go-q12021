package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maestre3d/academy-go-q12021/internal/application"
	"github.com/maestre3d/academy-go-q12021/internal/query"
	"github.com/maestre3d/academy-go-q12021/pkg/httputil"
)

// MovieHTTP Movie's HTTP controller used to expose Movie use cases through an HTTP server
type MovieHTTP struct {
	app *application.Movie
}

// NewMovieHTTP allocates a MovieHTTP controller for Uber Fx modules
func NewMovieHTTP(app *application.Movie) httputil.ControllersFx {
	return httputil.ControllersFx{
		Controller: &MovieHTTP{app: app},
	}
}

// MapRoutes maps exposed use cases from the current aggregate using the given mux.Router
func (m MovieHTTP) MapRoutes(r *mux.Router) {
	r.Path("/movies").Methods(http.MethodGet).HandlerFunc(m.listMovies)
	r.Path("/movies/{id}").Methods(http.MethodGet).HandlerFunc(m.getMovieByID)
}

func (m *MovieHTTP) listMovies(w http.ResponseWriter, r *http.Request) {
	criteria, err := httputil.UnmarshalCriteriaJSON(r)
	if err != nil {
		criteria = httputil.UnmarshalCriteria(r)
	}

	movies, err := query.HandleListMovies(r.Context(), m.app, query.ListMovies{Criteria: criteria})
	if err != nil {
		httputil.RespondErrJSON(w, r, err)
		return
	}
	httputil.RespondJSON(w, http.StatusOK, movies)
}

func (m *MovieHTTP) getMovieByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	movie, err := query.HandleGetMovieByID(r.Context(), m.app, query.GetMovieByID{
		ID: id,
	})
	if err != nil {
		httputil.RespondErrJSON(w, r, err)
		return
	}

	httputil.RespondJSON(w, http.StatusOK, movie)
}
