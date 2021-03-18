package controller

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/interactor"
	"net/http"
)

type jokeController struct {
	jokeInteractor interactor.JokeInteractor
}

type JokeController interface {
	GetJokes(c Context) error
}

func NewJokeController(ji interactor.JokeInteractor) JokeController {
	return &jokeController{ji}
}

func (jc *jokeController) GetJokes(c Context) error {
	var jokes []*model.Joke

	jokes, err := jc.jokeInteractor.Get(jokes)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, jokes)
}