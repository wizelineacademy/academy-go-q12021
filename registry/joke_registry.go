package registry

import (
	ic "bootcamp/interface/controller"
	ip "bootcamp/interface/presenter"
	ir "bootcamp/interface/repository"
	ui "bootcamp/usecase/interactor"
	up "bootcamp/usecase/presenter"
	ur "bootcamp/usecase/repository"
)

// NewJokeController attached func to Registry for interfacing Controller and Interactor.
// Returns a JokeController
func (r *registry) NewJokeController() ic.JokeController  {
	return ic.NewJokeController(r.NewJokeInteractor())
}

// NewJokeInteractor returns a new JokeInteractor
func (r *registry) NewJokeInteractor() ui.JokeInteractor  {
	return ui.NewJokeInteractor(r.NewJokeRepository(), r.NewJokePresenter(), ir.NewDBRepository(r.db))
}

// NewJokeRepository returns a JokeRepository
func (r *registry) NewJokeRepository() ur.JokeRepository {
	return ir.NewJokeRepository(r.db)
}

// NewJokePresenter returns a JokePresenter
func (r *registry) NewJokePresenter() up.JokePresenter {
	return ip.NewJokePresenter()
}
