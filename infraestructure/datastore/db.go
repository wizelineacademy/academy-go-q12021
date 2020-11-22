package datastore

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewDB returns the instance of the connection to sqlite
func NewDB(dbpath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
