package repository

import (
	"bootcamp/domain/model"
)

// JokeRepository interface for jokes
type JokeRepository interface {
	// GetAll function returns all jokes
	GetAll(jokes []*model.Joke) ([]*model.Joke, error)
}
