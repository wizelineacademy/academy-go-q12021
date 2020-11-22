package registry

import (
	"digimons/interface/controller"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

// Registry contain a pointer to an instance of form DB
type Registry interface {
	NewAppController() controller.AppController
}

// NewRegistry returns a new registry with a pointer to the database.
func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

// NewAppController Creates an app controller that includes the controller for other models.
func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Digimon: r.NewDigimonController(),
	}
}
