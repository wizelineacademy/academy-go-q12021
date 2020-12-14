package registry

import (
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/services"
	"github.com/alexis-aguirre/golang-bootcamp-2020/interface/presenter"
	ir "github.com/alexis-aguirre/golang-bootcamp-2020/interface/repository"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/interactor"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

//NewSongInteractor creates a new instance of SongInteractor
func NewSongInteractor() interactor.SongInteractor {
	registry := services.ServicesRegistry
	var service interface{}
	service = registry.FetchService(services.LOGGER)
	logger, _ := service.(services.Logger)
	return interactor.NewSongInteractor(NewSongRepository(logger), NewSongPresenter())
}

//NewSongRepository creates a new instance of SongRepository
func NewSongRepository(logger services.Logger) repository.SongRepository {
	return ir.NewSongRepository(logger)
}

//NewSongPresenter creates a new instance of SongPresenter
func NewSongPresenter() presenter.SongPresenter {
	return presenter.NewSongPresenter()
}
