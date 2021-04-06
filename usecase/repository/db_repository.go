package repository

// DBRepository interface for data manipulating transactions
type DBRepository interface {
	// Transaction
	Transaction(func(interface{}) (interface{}, error)) (interface{}, error)
}
