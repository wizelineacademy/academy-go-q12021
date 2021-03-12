package registry

import "github.com/AlejandroSeguraWIZ/academy-go-q12021/interface/controller"

type registry struct {
	fileName string
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(fn string) Registry {
	return &registry{fn}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewPokemonController()
}
