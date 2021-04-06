package registry

import (
	ic "bootcamp/interface/controller"
	ip "bootcamp/interface/presenter"
	ir "bootcamp/interface/repository"
	ui "bootcamp/usecase/interactor"
	up "bootcamp/usecase/presenter"
	ur "bootcamp/usecase/repository"
)

// NewItemController attached func to Registry for interfacing Controller and Interactor.
// Returns a ItemController.
func (r *registry) NewItemController() ic.ItemController {
	return ic.NewItemController(r.NewItemInteractor())
}

// NewItemInteractor returns an ItemInteractor.
func (r *registry) NewItemInteractor() ui.ItemInteractor {
	return ui.NewItemInteractor(r.NewItemRepository(), r.NewItemPresenter(), ir.NewDBRepository(r.db))
}

// NewItemRepository returns an ItemRepository.
func (r *registry) NewItemRepository() ur.ItemRepository {
	return ir.NewItemRepository(r.db)
}

// NewItemPresenter returns an ItemPresenter.
func (r *registry) NewItemPresenter() up.ItemPresenter {
	return ip.NewItemPresenter()
}
