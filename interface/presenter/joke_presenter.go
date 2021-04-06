package presenter

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/presenter"
)

// jokePresenter struct for JokePresenter
type jokePresenter struct {}

// NewJokePresenter return a JokePresenter
func NewJokePresenter() presenter.JokePresenter {
	return &jokePresenter{}
}

// ResponseJokes return an array of Jokes
func (jp *jokePresenter) ResponseJokes(jokes []*model.Joke) []*model.Joke {

	for _, joke := range jokes {
		joke.Joke = joke.Joke + "..."
	}
	return jokes
}