package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"main/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

const UintSize = 32 << (^uint(0) >> 32 & 1)
const MaxInt  = 1<<(UintSize-1) - 1


// UseCase interface
type UseCase interface {
	GetConcurrently(*model.QueryParameters, bool, string) ([]*model.Movie, error)
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

func (t *MovieUseCase) GetMovies(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")
	
	movies, err := t.useCase.GetMovies()
	if err != nil {
		log.Fatal("Failed on GetMovies : %w", err)
		t.render.JSON(w, http.StatusInternalServerError, movies)
	}
	
	totalTime :=  fmt.Sprintf("%d%s", time.Since(start).Microseconds(), " Microseconds.")
	
	jsonObject := model.Response{ 
		Title: "model.Response", 
		Results: 1,
		Message: "Data",
		Data: movies,
		Errors: nil, // TODO: Send errors too
		ExecutionTime: totalTime,
	}
	t.render.JSON(w, http.StatusOK, jsonObject)
}

// GET /movies_concurrently
func (t *MovieUseCase) GetConcurrently(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	// GET QUERY PARAMS AND VALIDATE
	// var queryParams model.model.QueryParameters = GetQueryParams(r)
	// GetConcurrently(nil, false, "")
	
	queryParams := model.QueryParameters{	
		ItemPerWorkers: 1,
		Items: MaxInt,
		Type: "",
	}

	movies, err := t.useCase.GetConcurrently(&queryParams, true, "") // TODO: send complete boolean and id string

	if err != nil {
		log.Fatal("Failed on GetMovies : %w", err)
		t.render.JSON(w, http.StatusInternalServerError, movies)
	}
	
	totalTime :=  fmt.Sprintf("%d%s", time.Since(start).Microseconds(), " Microseconds.")
	
	jsonObject := model.Response{ 
		Title: "model.Response", 
		Results: 1,
		Message: "Data",
		Data: movies,
		Errors: nil, // TODO: Send errors too
		ExecutionTime: totalTime,
	}

	t.render.JSON(w, http.StatusOK, jsonObject)
}

// GET /movies/{id}
func (t *MovieUseCase) GetMovieById(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	params := mux.Vars(r)
	
	movie, err := t.useCase.GetMovieById(params["id"])
	if err != nil {
		log.Fatal("Failed on GetMovieById : %w", err)
		t.render.JSON(w, http.StatusInternalServerError, movie)
	}

	totalTime :=  fmt.Sprintf("%d%s", time.Since(start).Microseconds(), " Microseconds.")
	
	jsonObject := model.Response{
		Title: "model.Response", 
		Results: 1,
		Message: "Data",
		Data: movie,
		Errors: nil, // TODO: Send errors too
		ExecutionTime: totalTime,
	}

	t.render.JSON(w, http.StatusOK, jsonObject)
}
