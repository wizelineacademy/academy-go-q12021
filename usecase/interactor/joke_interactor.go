package interactor

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/presenter"
	"bootcamp/usecase/repository"
)

type jokeInteractor struct {
	JokeRepository repository.JokeRepository
	JokePresenter presenter.JokerPesenter
	DBRpository repository.DBRepository
}

type JokeInteractor interface {
	Get(jokes []*model.Joke) ([]*model.Joke, error)
}

func NewJokeInteractor(jr repository.JokeRepository, jp presenter.JokerPesenter, db repository.DBRepository) JokeInteractor {
	return &jokeInteractor{jr,jp,db}
}

// Get: get all jokes from datastore and format it in the presenter layer
func (ji *jokeInteractor) Get(jokes []*model.Joke) ([]*model.Joke, error) {
	jks, err := ji.JokeRepository.GetAll(jokes)
	if err != nil {
		return nil, err
	}

	return ji.JokePresenter.ResponseJokes(jks), nil
}