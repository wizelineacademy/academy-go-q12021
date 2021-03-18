package presenter

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/presenter"
)

type jokePresenter struct {}

func NewJokePresenter() presenter.JokerPesenter {
	return &jokePresenter{}
}

func (jp *jokePresenter) ResponseJokes(jokes []*model.Joke) []*model.Joke {

	for _, joke := range jokes {
		joke.Joke = joke.Joke + "..."
	}
	return jokes
}