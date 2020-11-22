package repositories

import (
	"database/sql"
	"errors"

	"github.com/wizelineacademy/golang-bootcamp-2020/domain/models"
)

type ChampRepo struct {
	DB *sql.DB
}

func NewChampRepo(DB *sql.DB) *ChampRepo {
	return &ChampRepo{DB}
}

func (cr *ChampRepo) GetSingle(id int) (*models.Champion, error) {
	stmt := `SELECT  name, lore FROM champions WHERE id = ?`

	// This returns a pointer to a sql.Row object which holds the result from the database.
	row := cr.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Post struct.
	c := &models.Champion{}

	// Use row.Scan() to copy the values from each field in sql.Row to the
	// corresponding field in the Post struct. Notice that the arguments
	// to row.Scan are *pointers* to the place you want to copy the data into,
	// and the number of arguments must be exactly the same as the number of
	// columns returned by your statement.
	err := row.Scan(&c.Name, &c.Lore)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	//If everything went OK then return the Post struct.
	return c, nil
}

func (cr *ChampRepo) GetMultiple() ([]*models.Champion, error) {
	stmt := `SELECT name, lore FROM champions`

	// This returns a pointer to a sql.Row object which holds the result from the database.
	rows, err := cr.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize an empty slice to hold the models.Post objects.
	champions := []*models.Champion{}

	for rows.Next() {
		// Initialize a pointer to a new zeroed Post struct.
		c := &models.Champion{}

		// Use row.Scan() to copy the values from each field in sql.Row to the
		// corresponding field in the Post struct. Notice that the arguments
		// to row.Scan are *pointers* to the place you want to copy the data into,
		// and the number of arguments must be exactly the same as the number of
		// columns returned by your statement.
		err := rows.Scan(&c.Name, &c.Lore)

		if err != nil {
			return nil, err
		}
		champions = append(champions, c)

	}

	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
	// error that was encountered during the iteration. It's important to
	// call this - don't assume that a successful iteration was completed
	// over the whole resultset.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	//If everything went OK then return the Post struct.
	return champions, nil
}

func (cr *ChampRepo) Update(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (cr *ChampRepo) Delete(db *sql.DB) error {
	return errors.New("Not implemented")
}
