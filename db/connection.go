package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// GetConnection should return connection to the db
func GetConnection() *sql.DB {
	connStr := "root:password@tcp(localhost:4545)/bitcoin_db"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
