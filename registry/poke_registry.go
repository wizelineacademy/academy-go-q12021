package registry

import (
	"github.com/ToteEmmanuel/academy-go-q12021/interface/controller"
	interfacePresenter "github.com/ToteEmmanuel/academy-go-q12021/interface/presenter"
	interfaceRepository "github.com/ToteEmmanuel/academy-go-q12021/interface/repository"
	"github.com/ToteEmmanuel/academy-go-q12021/usecase/interactor"
	usecasePresenter "github.com/ToteEmmanuel/academy-go-q12021/usecase/presenter"
	usecaseRepository "github.com/ToteEmmanuel/academy-go-q12021/usecase/repository"
)

func (r *registry) NewPokeController() controller.PokeController {
	return controller.NewPokeController(r.NewPokeInteractor())
}

func (r *registry) NewPokeInteractor() interactor.PokeInteractor {
	return interactor.NewPokeInteractor(r.NewPokeRepository(), r.NewPokePresenter())
}

func (r *registry) NewPokeRepository() usecaseRepository.PokeRepository {
	return interfaceRepository.NewPokeRepository(r.storage)
}

func (r *registry) NewPokePresenter() usecasePresenter.PokePresenter {
	return interfacePresenter.NewPokePresenter()
}
