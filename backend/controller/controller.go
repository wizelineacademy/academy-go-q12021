package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"main/model"

	"github.com/unrolled/render"
)

const UintSize = 32 << (^uint(0) >> 32 & 1)
const MaxInt = 1<<(UintSize-1) - 1

var requestErrors []string

// UseCase interface
type UseCase interface {
	GetConcurrently(model.QueryParameters, bool, string) ([]interface{}, error)
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

	totalTime := fmt.Sprintf("%d%s", time.Since(start).Microseconds(), " Microseconds.")

	jsonObject := model.Response{
		Title:         "model.Response",
		Results:       1,
		Message:       "Data",
		Data:          movies,
		Errors:        requestErrors,
		ExecutionTime: totalTime,
	}
	t.render.JSON(w, http.StatusOK, jsonObject)
}

// GET /movies_concurrently
func (t *MovieUseCase) GetConcurrently(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	// GET QUERY PARAMS AND VALIDATE
	var queryParams model.QueryParameters = GetQueryParams(r)

	log.Println("\n\t QUERYPARAMS", queryParams.Items, queryParams.ItemPerWorkers, queryParams.Type)

	willRequireMovieComplete := false

	movies, err := t.useCase.GetConcurrently(queryParams, willRequireMovieComplete, "")
	if err != nil {
		log.Fatal("Failed on GetMovies : %w", err)
		t.render.JSON(w, http.StatusInternalServerError, movies)
	}

	totalTime := fmt.Sprintf("%d%s", time.Since(start).Microseconds(), " Microseconds.")

	jsonObject := model.Response{
		Title:         "model.Response",
		Results:       len(movies),
		Message:       "Data",
		Data:          movies,
		Errors:        requestErrors,
		ExecutionTime: totalTime,
	}

	t.render.JSON(w, http.StatusOK, jsonObject)
}

// GET /movies/{id}
func (t *MovieUseCase) GetMovieById(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")

	// GET QUERY PARAMS AND VALIDATE
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys) <= 0 {
		errorMessage := "Id query param is required but missing"
		requestErrors = append(requestErrors, errorMessage)
		response := model.Response{
			Data:          nil,
			Title:         "Error",
			Message:       errorMessage,
			Errors:        requestErrors,
			Results:       0,
			ExecutionTime: fmt.Sprintf("%d%s", time.Since(start).Microseconds(), " Microseconds."),
		}
		t.render.JSON(w, http.StatusInternalServerError, response)
		log.Println(errorMessage)
		return
	}
	var id string
	if ok {
		id = keys[0]
	} else {
		id = ""
	}

	queryParams := model.QueryParameters{Items: 1, Type: "", ItemPerWorkers: 1}
	willRequireMovieComplete := true

	log.Println("Will call the GetConcurrently function with params: ", queryParams, willRequireMovieComplete, id)
	movies, err := t.useCase.GetConcurrently(queryParams, true, id)
	if err != nil {
		log.Println("Failed on GetMovieById : %w", err)
		t.render.JSON(w, http.StatusInternalServerError, movies)
		return
	}

	jsonObject := model.Response{
		Title:         "model.Response",
		Results:       1,
		Message:       "Data",
		Data:          movies,
		Errors:        requestErrors,
		ExecutionTime: fmt.Sprintf("%d%s", time.Since(start).Microseconds(), " Microseconds."),
	}

	t.render.JSON(w, http.StatusOK, jsonObject)
}

func GetQueryParams(r *http.Request) (queryParams model.QueryParameters) {
	keys := r.URL.Query()

	if val, ok := keys["type"]; ok {
		log.Println("Type query provided")
		queryParams.Type = val[0]
		if queryParams.Type != "odd" && queryParams.Type != "even" {
			log.Println("Type defafult empty")
			queryParams.Type = ""
		}
	} else {
		requestErrors = append(requestErrors, "`type` was not provided as query param. Should be rather odd or even.")
		log.Println("Type not provided as query param.")
	}
	if val, ok := keys["item_per_workers"]; ok {
		IntItemPerWorkers, err := strconv.Atoi(val[0]) // parse string to int
		if err != nil {
			requestErrors = append(requestErrors, err.Error())
			queryParams.ItemPerWorkers = 1
		} else {
			log.Println("item_per_workers query provided")
			queryParams.ItemPerWorkers = IntItemPerWorkers
		}
	} else {
		requestErrors = append(requestErrors, "`items_per_workers` was not provided as query param.")
		log.Println("item_per_workers not provided as query param")
	}

	if val, ok := keys["items"]; ok {
		IntItems, err := strconv.Atoi(val[0]) // parse string to int
		if err != nil {
			requestErrors = append(requestErrors, err.Error()+". Number should be positive integer. The items param will be considered as 0. ")
			queryParams.Items = 0
		} else {
			queryParams.Items = IntItems
			log.Println("items query provided: value ", IntItems)
		}
	} else {
		requestErrors = append(requestErrors, "`items` was not provided as query param: MaxValue")
		queryParams.Items = MaxInt
	}
	return
}
