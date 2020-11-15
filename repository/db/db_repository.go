package db

type dbRepository struct {
}

type DbRepository interface {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
