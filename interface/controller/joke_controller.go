package controller

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/interactor"
	"net/http"
)

// jokeController struct for JokeInteractor
type jokeController struct {
	jokeInteractor interactor.JokeInteractor
}

// JokeController interface
type JokeController interface {
	GetJokes(c Context) error
}

// NewJokeController returns a JokeController
func NewJokeController(ji interactor.JokeInteractor) JokeController {
	return &jokeController{ji}
}

// GetJokes return an array of Jokes
func (jc *jokeController) GetJokes(c Context) error {
	var jokes []*model.Joke

	jokes, err := jc.jokeInteractor.Get(jokes)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, jokes)
}