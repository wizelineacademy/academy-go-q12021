package service

type DataService interface {
	Init() error
	Get(id int) (interface{}, error)
	List(count, page int) ([]interface{}, error)
}
