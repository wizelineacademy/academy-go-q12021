package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	database *sql.DB
)

// Bitcoin -> Structure of the database
type Bitcoin struct {
	ID       int
	Base     string
	Currency string
	Amount   string
}

func init() {
	GetConnection()
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// GetConnection should return connection to the db
func GetConnection() {
	connStr := goDotEnvVariable("DB_CONNECTION")
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	database = db
}

// InsertBitcoinValue -> This inserts a new entry to the db
func InsertBitcoinValue(base string, currency string, amount string) (Bitcoin, error) {

	res, err := database.Exec("INSERT INTO bitcoins (base, currency, amount) VALUES(?, ?, ?); ", base, currency, amount)

	if err != nil {
		log.Fatal(err)
		return Bitcoin{}, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return Bitcoin{}, err
	}

	println(lastID)

	return Bitcoin{int(lastID), base, currency, amount}, nil
}

// GetValue -> This should get one value from the db based on the id
func GetValue(id int) (Bitcoin, error) {
	row := database.QueryRow("SELECT * FROM bitcoins where id = ?", id)

	var ID int
	var base string
	var currency string
	var amount string

	err := row.Scan(&ID, &base, &currency, &amount)
	if err != nil {
		return Bitcoin{}, err
	}

	return Bitcoin{ID, base, currency, amount}, nil
}

// GetAllValues -> Return all values from the db
func GetAllValues() []Bitcoin {

	var btc []Bitcoin
	rows, err := database.Query("SELECT * FROM bitcoins ORDER BY id;")
	if err != nil {
		log.Fatal(err)
		return btc
	}

	defer rows.Close()

	for rows.Next() {

		var ID int
		var base string
		var amount string
		var currency string

		err := rows.Scan(&ID, &base, &currency, &amount)
		if err != nil {
			log.Fatal(err)
		}

		b := Bitcoin{
			ID:       ID,
			Currency: currency,
			Base:     base,
			Amount:   amount,
		}

		btc = append(btc, b)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return btc
}

// DeleteValue -> This should delete one value from the db based on the id
func DeleteValue(id int) (Bitcoin, error) {
	entry, err := GetValue(id)

	if err != nil {
		return Bitcoin{}, err
	}
	_, err = database.Exec("DELETE FROM bitcoins where id = ?;", id)

	if err != nil {
		return Bitcoin{}, err
	}

	return entry, nil
}
