package registry

import (
	"github.com/AlejandroSeguraWIZ/academy-go-q12021/interface/controller"
	ip "github.com/AlejandroSeguraWIZ/academy-go-q12021/interface/presenter"
	ir "github.com/AlejandroSeguraWIZ/academy-go-q12021/interface/repository"
	"github.com/AlejandroSeguraWIZ/academy-go-q12021/usecase/interactor"
	"github.com/AlejandroSeguraWIZ/academy-go-q12021/usecase/presenter"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonRepository() ir.PokemonRepository {
	return ir.NewPokemonRepository(r.fileName)
}

func (r *registry) NewPokemonPresenter() presenter.PokemonPresenter {
	return ip.NewPokemonPresenter()
}
