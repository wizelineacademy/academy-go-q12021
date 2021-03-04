package repository

import (
	"bootcamp/usecase/repository"
	"database/sql"
	"errors"
	"log"
)

type dbRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) repository.DBRepository {
	return &dbRepository{db}
}

func (r *dbRepository) Transaction(txFunc func(interface{}) (interface{}, error)) (data interface{}, err error) {
	tx, err := r.db.Begin()
	if !errors.Is(err, nil) {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			log.Print("recover")
			tx.Rollback()
			panic(p)
		} else if !errors.Is(err, nil) {
			log.Print("rollback")
			tx.Rollback()
			panic("error")
		} else {
			err = tx.Commit()
		}
	}()

	data, err = txFunc(tx)
	return data, err
}