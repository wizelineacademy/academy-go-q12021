package service

import (
	"github.com/wizelineacademy/academy-go/model"
)

// DataService is an interface for the modules to implement the specigic logic of the entities
type DataService interface {
	Init() error
	Get(id int) model.Response
	List(count, page int) model.Response
}
