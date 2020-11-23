package repositories

import (
	"database/sql"
	"errors"

	"github.com/wizelineacademy/golang-bootcamp-2020/domain/models"
)

// Errors

// ErrNoRecord is a custom error used in case no Champion is found in the database
var ErrNoRecord = errors.New("repositories: no matching record found")

// ChampionRepository defines the interface used by a Champion struct to access its repositories methods
type ChampionRepository interface {
	GetSingle(id int) (*models.Champion, error)
	GetMultiple(limit int) ([]*models.Champion, error)
	Insert(champion *models.Champion) (int, error)
}

// ChampRepo defines the link between the Champion and the Database
type ChampRepo struct {
	DB *sql.DB
}

// NewChampRepo returns an inituialized ChampRepo struct
func NewChampRepo(DB *sql.DB) *ChampRepo {
	return &ChampRepo{DB}
}

// GetSingle gets a single database row and returns it as a Champion
func (cr *ChampRepo) GetSingle(id int) (*models.Champion, error) {
	// SQL statement
	stmt := `SELECT  name, lore, created FROM champions WHERE id = ?`

	// This returns a pointer to a sql.Row object which holds the result from the database.
	row := cr.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Champion struct.
	c := &models.Champion{}

	// Copy the values from each field in sql.Row to the corresponding field in the Champion struct.
	err := row.Scan(&c.Name, &c.Lore, &c.DateCreated)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	//If everything went OK then return the Champion struct.
	return c, nil
}

// GetMultiple query the DB and returns a slice of Champions
func (cr *ChampRepo) GetMultiple(limit int) ([]*models.Champion, error) {
	// SQL statement
	stmt := `SELECT name, lore, created FROM champions LIMIT ?`

	// This returns a pointer to a sql.Row object which holds the result from the database.
	rows, err := cr.DB.Query(stmt, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize an empty slice to hold the Champion structs.
	champions := []*models.Champion{}

	for rows.Next() {
		// Initialize a pointer to a new zeroed Champion struct.
		c := &models.Champion{}

		// Use row.Scan() to copy the values from each field in sql.Row to the corresponding field in the Champion struct.
		err := rows.Scan(&c.Name, &c.Lore, &c.DateCreated)
		if err != nil {
			return nil, err
		}

		champions = append(champions, c)

	}

	// When the rows.Next() loop has finished we call rows.Err() to retrieve any error that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	//If everything went OK then return the Champion struct.
	return champions, nil
}

// Insert a new Champion into the database.
func (cr *ChampRepo) Insert(champion *models.Champion) (int, error) {
	stmt := `INSERT INTO champions (name, lore) VALUES (?, ?)`

	// Execute the SQL statement
	result, err := cr.DB.Exec(stmt, champion.Name, champion.Lore)
	if err != nil {
		return 0, err
	}

	// Get the ID of the newly inserted record.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type before returning it.
	return int(id), nil
}

func (cr *ChampRepo) Update(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (cr *ChampRepo) Delete(db *sql.DB) error {
	return errors.New("Not implemented")
}
