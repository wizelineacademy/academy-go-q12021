package registry

import (
	"bootcamp/interface/controller"
	ip "bootcamp/interface/presenter"
	"bootcamp/usecase/interactor"
	up "bootcamp/usecase/presenter"
	ur "bootcamp/usecase/repository"
	ir "bootcamp/interface/repository"
)

func (r *registry) NewItemController() controller.ItemController {
	return controller.NewItemController(r.NewItemInteractor())
}

func (r *registry) NewItemInteractor() interactor.ItemInteractor {
	return interactor.NewItemInteractor(r.NewItemRepository(), r.NewItemPresenter(), ir.NewDBRepository(r.db))
}

func (r *registry) NewItemRepository() ur.ItemRepository {
	return ir.NewItemRepository(r.db)
}

func (r *registry) NewItemPresenter() up.ItemPresenter {
	return ip.NewItemPresenter()
}
