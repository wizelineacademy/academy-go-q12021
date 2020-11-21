package registry

import (
	"digimons/interface/controller"
	ip "digimons/interface/presenter"
	ir "digimons/interface/repository"
	"digimons/usecase/interactor"
	up "digimons/usecase/presenter"
	ur "digimons/usecase/repository"
)

func (r *registry) NewDigimonController() controller.DigimonController {
	return controller.NewDigimonController(r.NewDigimonInteractor())
}

func (r *registry) NewDigimonInteractor() interactor.DigimonInteractor {
	return interactor.NewDigimonInteractor(r.NewDigimonRepository(), r.NewDigimonPresenter())
}

func (r *registry) NewDigimonRepository() ur.DigimonRepository {
	return ir.NewDigimonRepository(r.db)
}

func (r *registry) NewDigimonPresenter() up.DigimonPresenter {
	return ip.NewDigimonPresenter()
}
