package presenter

import "bootcamp/domain/model"

type JokerPesenter interface {
	ResponseJokes(jokes []*model.Joke) []*model.Joke
}
