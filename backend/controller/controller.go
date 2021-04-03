package controller

import (
	"net/http"

	"main/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// UseCase interface
type UseCase interface {
	GetMovies() ([]*model.Movie, error)
	GetMovieById(string) (*model.Movie, error)
}

// MovieUseCase struct
type MovieUseCase struct {
	useCase UseCase
	render  *render.Render
}

// New returns a controller
func New(
	u UseCase,
	r *render.Render,
) *MovieUseCase {
	return &MovieUseCase{u, r}
}

// GET /movies
func (t *MovieUseCase) GetMovies(w http.ResponseWriter, r *http.Request) {
	body, _ := t.useCase.GetMovies()
	w.Header().Set("Content-Type", "application/json")
	t.render.JSON(w, http.StatusOK, body)
}

// GET /movies/{id}
func (t *MovieUseCase) GetMovieById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	body, _ := t.useCase.GetMovieById(params["id"])
	w.Header().Set("Content-Type", "application/json")
	t.render.JSON(w, http.StatusOK, body)
}
