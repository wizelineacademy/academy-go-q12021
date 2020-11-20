package interactor

import (
	"log"

	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/presenter"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

type songInteractor struct {
	SongRepository repository.SongRepository
	SongPresenter  presenter.SongPresenter
}

type SongInteractor interface {
	Get(s *model.Song) (*model.Song, error)
}

//NewSongInteractor generates a new instance of a song interactor
func NewSongInteractor(r repository.SongRepository, p presenter.SongPresenter) SongInteractor {
	return &songInteractor{r, p}
}

func (si *songInteractor) Get(s *model.Song) (*model.Song, error) {
	log.Println("Here in Song_interactor.Get")
	s, err := si.SongRepository.Find(s)
	if err != nil {
		return nil, err
	}

	return si.SongPresenter.ResponseSong(s), nil
}
