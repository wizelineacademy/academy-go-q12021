package models

import (
	"log"

	"../../db"
)

// Bitcoin -> Structure of the database
type Bitcoin struct {
	ID       int
	Base     string
	Currency string
	Amount   string
}

// InsertBitcoinValue -> This inserts a new entry to the db
func InsertBitcoinValue(base string, currency string, amount string) (Bitcoin, bool) {
	database := db.GetConnection()

	res, err := database.Exec("INSERT INTO bitcoins (base, currency, amount) VALUES(?, ?, ?); ", base, currency, amount)

	if err != nil {
		println(err.Error())
		return Bitcoin{}, false
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		println(err.Error())
		return Bitcoin{}, false
	}

	println(lastID)

	return Bitcoin{int(lastID), base, currency, amount}, true
}

// GetValue -> This should get one value from the db based on the id
func GetValue(id int) (Bitcoin, bool) {
	database := db.GetConnection()
	row := database.QueryRow("SELECT * FROM bitcoins where id = ?", id)

	var ID int
	var base string
	var currency string
	var amount string

	err := row.Scan(&ID, &base, &currency, &amount)
	if err != nil {
		return Bitcoin{}, false
	}

	return Bitcoin{ID, base, currency, amount}, true
}

// GetAllValues -> Return all values from the db
func GetAllValues() []Bitcoin {
	database := db.GetConnection()

	var btc []Bitcoin
	rows, err := database.Query("SELECT * FROM bitcoins ORDER BY id;")
	if err != nil {
		log.Println(err)
		return btc
	}

	defer rows.Close()

	for rows.Next() {
		b := Bitcoin{}

		var ID int
		var base string
		var amount string
		var currency string

		err := rows.Scan(&ID, &base, &currency, &amount)
		if err != nil {
			log.Fatal(err)
		}

		b.ID = ID
		b.Currency = currency
		b.Base = base
		b.Amount = amount

		btc = append(btc, b)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return btc
}

// DeleteValue -> This should delete one value from the db based on the id
func DeleteValue(id int) (Bitcoin, bool) {
	database := db.GetConnection()
	entry, success := GetValue(id)

	if success != true {
		return Bitcoin{}, false
	}
	_, err := database.Exec("DELETE FROM bitcoins where id = ?;", id)

	if err != nil {
		return Bitcoin{}, false
	}

	return entry, true
}
