package interactor

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/presenter"
	"bootcamp/usecase/repository"
)

// jokeInteractor struct for connecting layers
type jokeInteractor struct {
	JokeRepository repository.JokeRepository
	JokePresenter  presenter.JokePresenter
	DBRepository   repository.DBRepository
}

// JokeInteractor interface with main functions
type JokeInteractor interface {
	// Get return all jokes
	Get(jokes []*model.Joke) ([]*model.Joke, error)
}

// NewJokeInteractor returns a new struct interactor
func NewJokeInteractor(jr repository.JokeRepository, jp presenter.JokePresenter, db repository.DBRepository) JokeInteractor {
	return &jokeInteractor{jr,jp,db}
}

// Get func get all jokes from datastore and format they in the presenter layer
func (ji *jokeInteractor) Get(jokes []*model.Joke) ([]*model.Joke, error) {
	jks, err := ji.JokeRepository.GetAll(jokes)
	if err != nil {
		return nil, err
	}

	return ji.JokePresenter.ResponseJokes(jks), nil
}