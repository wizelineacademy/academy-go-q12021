package services

//Database is the interface which should accomplish each object to be used as a DB
type Database interface {
	Create(interface{}) (interface{}, error)
	Update(interface{}) (interface{}, error)
	Delete() error
	Get(interface{}) (interface{}, error)
}
