package db

import (
	"database/sql"
	"fmt"
	// Required by database/sql
	_ "github.com/lib/pq"
  )

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "passwd"
	dbname   = "postgres"
)

// PgPing checks if we are able to connect
func PgPing() {
	pgInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}