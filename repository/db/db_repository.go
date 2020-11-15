package db

type dbRepository struct {
}

type DataBaseRepository interface {
}

func NewDbRepository() DataBaseRepository {
	return &dbRepository{}
}
