package datastore

import (
	"database/sql"
	_ "github.com/mithrandie/csvq-driver"
	"log"
)

func NewDB() *sql.DB {

	DBMS := "csvq"

	db, err := sql.Open(DBMS, ".")

	if err != nil {
		log.Fatalln(err)
	}
	return db
	/**
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	queryString := "SELECT id, first_name, country_code from `users.csv` WHERE id= "

	r := db.QueryRowContext(ctx, queryString)

	var (
		id int
		firstName string
		countryCode string
	)

	if err := r.Scan(&id, &firstName, &countryCode)

	 */
}

