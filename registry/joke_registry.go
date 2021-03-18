package registry

import (
	ic "bootcamp/interface/controller"
	ip "bootcamp/interface/presenter"
	ui "bootcamp/usecase/interactor"
	up "bootcamp/usecase/presenter"
	ur "bootcamp/usecase/repository"
	ir "bootcamp/interface/repository"
)

func (r *registry) NewJokeController() ic.JokeController  {
	return ic.NewJokeController(r.NewJokeInteractor())
}

func (r *registry) NewJokeInteractor() ui.JokeInteractor  {
	return ui.NewJokeInteractor(r.NewJokeRepository(), r.NewJokePresenter(), ir.NewDBRepository(r.db))
}

func (r *registry) NewJokeRepository() ur.JokeRepository {
	return ir.NewJokeRepository(r.db)
}

func (r *registry) NewJokePresenter() up.JokerPesenter {
	return ip.NewJokePresenter()
}
