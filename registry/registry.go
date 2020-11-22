package registry

import (
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/datastore"
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/services"
	"github.com/alexis-aguirre/golang-bootcamp-2020/interface/presenter"
	ir "github.com/alexis-aguirre/golang-bootcamp-2020/interface/repository"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/interactor"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

//NewUserInteractor creates a new instance of UserInteractor
func NewUserInteractor() interactor.UserInteractor {
	registry := services.ServicesRegistry
	db := &datastore.MySQL{}
	registry.FetchService(db)
	logger := &datastore.Logger{}
	registry.FetchService(logger)
	return interactor.NewUserInteractor(NewUserRepository(db, logger), NewUserPresenter())
}

//NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db services.Database, logger services.Logger) repository.UserRepository {
	return ir.NewUserRepository(db, logger)
}

//NewUserPresenter creates a new instance of UserPresenter
func NewUserPresenter() presenter.UserPresenter {
	return presenter.NewUserPresenter()
}
