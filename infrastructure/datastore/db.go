package datastore

//Database is the db object that will process all the db tasks
type Database interface {
	GetAll(tableName string) ([]map[string]interface{}, error)
	GetItemByID(tableName string, id string) (map[string]interface{}, error)
	DeleteItem(tableName string, id string) error
	UpdateItem(tableName string, id string, item map[string]interface{}) (map[string]interface{}, error)
	AddItem(tableName string, item map[string]interface{}) (map[string]interface{}, error)
}
