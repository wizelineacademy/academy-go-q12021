package registry

import (
	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/services"
	"github.com/alexis-aguirre/golang-bootcamp-2020/interface/presenter"
	ir "github.com/alexis-aguirre/golang-bootcamp-2020/interface/repository"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/interactor"
	"github.com/alexis-aguirre/golang-bootcamp-2020/usecase/repository"
)

//NewAdminInteractor creates a new instance of AdminInteractor
func NewAdminInteractor() interactor.AdminInteractor {
	return interactor.NewAdminInteractor(NewAdminRepository(), NewAdminPresenter())
}

//NewAdminRepository creates a new instance of AdminRepository
func NewAdminRepository() repository.AdminRepository {
	registry := services.ServicesRegistry
	var service interface{}
	service = registry.FetchService(services.LOGGER)
	logger, _ := service.(services.Logger)
	return ir.NewAdminRepository(logger)
}

//NewAdminPresenter creates a new instance of AdminPresenter
func NewAdminPresenter() presenter.AdminPresenter {
	return presenter.NewAdminPresenter()
}
