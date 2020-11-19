package db

import (
	"database/sql"
	"fmt"
	"log"

	// Required by database/sql
	_ "github.com/lib/pq"

	"github.com/adantop/golang-bootcamp-2020/pokemon"
)

// PostgreSQL is a datasource type for pg connection
type PostgreSQL struct {
	database *sql.DB
}

// UsePostgreSQL overwrites the db module DB var with a pg implementation
func UsePostgreSQL() {
	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "passwd"
		dbname   = "postgres"
		pgDSN    = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	)

	db, err := sql.Open("postgres", pgDSN)

	if err != nil {
		log.Fatalf("Unable to create db object: %v\n", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not establish connection to the database: %v", err)
	}

	DS = &PostgreSQL{db}
}

// GetPokemonByName Get pokemon by name
func (ds *PostgreSQL) GetPokemonByName(name string) (p pokemon.Pokemon, err error) {

	err = ds.database.QueryRow(queryPokemonByName, name).Scan(
		&p.Number,
		&p.Name,
		&p.Type1,
		&p.Type2,
		&p.HeightM,
		&p.WeightKg,
		&p.Male,
		&p.Female,
		&p.CaptRate,
		&p.HP,
		&p.Attack,
		&p.Defense,
		&p.Special,
		&p.Speed)

	return
}

// Close terminates the database connection
func (ds *PostgreSQL) Close() {
	ds.database.Close()
}
