package presenter

import "bootcamp/domain/model"

// JokePresenter interface for jokes
type JokePresenter interface {
	// ResponseJokes returns manipulated jokes
	ResponseJokes(jokes []*model.Joke) []*model.Joke
}
