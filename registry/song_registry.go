package registry

import (
	"github.com/alexis-aguirre/golang-bootcamp-2020/interface/presenter"
	ir "github.com/alexis-aguirre/golang-bootcamp-2020/interface/repository"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/interactor"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

//NewSongInteractor creates a new instance of SongInteractor
func NewSongInteractor() interactor.SongInteractor {
	return interactor.NewSongInteractor(NewSongRepository(), NewSongPresenter())
}

//NewSongRepository creates a new instance of SongRepository
func NewSongRepository() repository.SongRepository {
	return ir.NewSongRepository()
}

//NewSongPresenter creates a new instance of SongPresenter
func NewSongPresenter() presenter.SongPresenter {
	return presenter.NewSongPresenter()
}
