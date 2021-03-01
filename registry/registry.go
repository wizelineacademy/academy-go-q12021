package registry

import "github.com/Topi99/academy-go-q12021/interface/controller"

type registry struct {
}

// Registry interface
type Registry interface {
	NewAppController() controller.AppController
}

// NewRegistry returns new Registry
func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Pokemon: r.NewPokemonController(),
	}
}
