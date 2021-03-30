package datastore

import (
	"database/sql"
	_ "github.com/mithrandie/csvq-driver"
	"log"
)

// NewDB sets up the SQL driver and settings
func NewDB() *sql.DB {

	DBMS := "csvq"

	db, err := sql.Open(DBMS, ".")

	if err != nil {
		log.Fatalln(err)
	}
	return db
}

