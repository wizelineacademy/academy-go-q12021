package db

import (
	"database/sql"
	"log"

	// Registering sqlite3 driver into database/sql
	_ "github.com/mattn/go-sqlite3"

	"github.com/adantop/golang-bootcamp-2020/pokemon"
)

// SQLite3 is a datasource type for sqlite3 connection
type SQLite3 struct {
	database *sql.DB
}

// UseSQLite3 overwrites the db module DB var with an SQLite3 implementation
func UseSQLite3(dbfile string) {

	db, err := sql.Open("sqlite3", dbfile)

	if err != nil {
		log.Fatalf("Unable to create db object: %v\n", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not establish connection to the database: %v", err)
	}

	DS = &SQLite3{db}
}

// GetPokemonByName Get pokemon by name
func (ds *SQLite3) GetPokemonByName(name string) (p pokemon.Pokemon, err error) {

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
func (ds *SQLite3) Close() {
	ds.database.Close()
}
