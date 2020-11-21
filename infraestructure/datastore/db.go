package datastore

import (
	"log"

	// "github.com/jinzhu/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewDB returns the instance of the connection to sqlite
func NewDB() *gorm.DB {
	DBName := "config/goolang-bootcamp-2020.db"
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
