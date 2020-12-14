package repositories

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
	"github.com/wizelineacademy/golang-bootcamp-2020/domain/models"
)

const fileOffset = 0

// Errors

// ErrNoChar is a custom error used in case no Character is found in the CSV file
var ErrNoChar = errors.New("repositories: no matching char found")

// CharacterRepository defines the interface used by a Character struct to access its repositories methods
type CharacterRepository interface {
	Insert(char *models.Character) error
	Get(id int) (*models.Character, error)
}

// CharRepo defines the link between the Character and the CSV
type CharRepo struct {
	file *os.File
}

// NewCharRepo returns an initialized CharRepo struct
func NewCharRepo(file *os.File) *CharRepo {
	return &CharRepo{file}
}

// Insert inserts a new Rick and Morty Character into a local CSV file.
func (cr *CharRepo) Insert(character *models.Character) error {

	writer := csv.NewWriter(cr.file)
	var row []string
	row = append(row, strconv.Itoa(character.ID))
	row = append(row, character.Name)
	row = append(row, character.Status)
	row = append(row, character.Species)
	err := writer.Write(row)
	if err != nil {
		return err
	}

	writer.Flush()

	return nil

}

// Get gets a new Rick and Morty Character from a local CSV file.
func (cr *CharRepo) Get(id int) (*models.Character, error) {
	// Setting the file pointer to the beginning
	if _, err := cr.file.Seek(fileOffset, io.SeekStart); err != nil {
		return nil, err
	}

	// Empty struct to hold the unmarshaled data
	characters := []*models.Character{}
	// Empty struct to hold the char data
	character := &models.Character{}

	err := gocsv.UnmarshalFile(cr.file, &characters)
	if err != nil {
		return nil, err
	}

	for _, char := range characters {
		if char.ID == id {
			character = char
			break
		}
	}

	if character.ID == 0 { // 0 id means an zeroed struct with no values
		return nil, ErrNoChar
	}

	return character, nil
}
