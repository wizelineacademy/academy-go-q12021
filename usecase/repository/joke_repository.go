package repository

import (
	"bootcamp/domain/model"
)

type JokeRepository interface {
	GetAll(jokes []*model.Joke) ([]*model.Joke, error)
}