package db

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang-bootcamp-2020/domain/model"
	_errors "golang-bootcamp-2020/utils/error"

	"github.com/spf13/viper"
)

var (
	isCsvFetched  = false
	charactersMap map[string]*model.Character
)

const (
	chID = iota
	chName
	chStatus
	chSpecies
	chType
	chGender
	chOriginName
	chOriginUrl
	chLocationName
	chLocationUrl
	chImage
	chEpisodes

	errorDbEmpty     = "db empty, fetch is needed"
	errorWritingFile = "error writing csv file"
)

type dbRepository struct {
	mapReader *io.Reader
}

//DataBaseRepository - database repository methods
type DataBaseRepository interface {
	CreateCharactersCSV(characters []model.Character) _errors.RestError
	GetCharacterFromId(id string) (*model.Character, _errors.RestError)
	GetCharacters() ([]model.Character, _errors.RestError)
	GetCharacterIdByName(name string) (string, _errors.RestError)
}

//Init - Set the initial state of db repository
func Init() {
	isCsvFetched = readCharactersFromCSV()
}

//NewDbRepository - Return new db repository and init it
func NewDbRepository() DataBaseRepository {
	Init()
	return &dbRepository{}
}

// CreateCharactersCSV - Make csv file given array of characters
func (db *dbRepository) CreateCharactersCSV(characters []model.Character) _errors.RestError {
	file, err := os.Create(viper.GetString("db.charactersPath"))

	// ignoring close error it's safe on this point: https://www.joeshaw.org/dont-defer-close-on-writable-files/
	defer file.Close()
	defer readCharactersFromCSV()

	if err != nil {
		return _errors.NewInternalServerError(errorWritingFile)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, ch := range characters {
		var row []string
		row = append(row, strconv.Itoa(ch.ID))
		row = append(row, ch.Name)
		row = append(row, ch.Status)
		row = append(row, ch.Species)
		row = append(row, ch.Type)
		row = append(row, ch.Gender)
		row = append(row, ch.Origin.Name)
		row = append(row, ch.Origin.URL)
		row = append(row, ch.Location.Name)
		row = append(row, ch.Location.URL)
		row = append(row, ch.Image)

		var episodesString string
		for _, episode := range ch.Episodes {
			episodesString += fmt.Sprintf("+%s", episode)
		}

		row = append(row, episodesString)
		if err := writer.Write(row); err != nil {
			return _errors.NewInternalServerError(errorWritingFile)
		}
	}
	return nil
}

//GetCharacterFromId - Get character from map (complex O(1) )
func (db *dbRepository) GetCharacterFromId(id string) (*model.Character, _errors.RestError) {
	if !isCsvFetched {
		return nil, _errors.NewInternalServerError(errorDbEmpty)
	}

	ch, ok := charactersMap[id]
	if ch == nil || !ok {
		return nil, _errors.NewInternalServerError(fmt.Sprintf("character with id %s not found, fetch with more pages to update the db", id))
	}

	return ch, nil
}

//GetCharacters - Get all characters from map
func (db *dbRepository) GetCharacters() ([]model.Character, _errors.RestError) {
	if !isCsvFetched {
		return nil, _errors.NewInternalServerError(errorDbEmpty)
	}

	var characters []model.Character
	for _, ch := range charactersMap {
		characters = append(characters, *ch)
	}

	return characters, nil
}

//GetCharacterIdByName - Get character id from csv map (complex O(n) )
func (db *dbRepository) GetCharacterIdByName(name string) (string, _errors.RestError) {
	if !isCsvFetched {
		return "", _errors.NewInternalServerError(errorDbEmpty)
	}

	file, err := os.Open(viper.GetString("db.mapPath"))
	defer file.Close()

	if err != nil {
		return "", _errors.NewInternalServerError(errorDbEmpty)
	}

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", _errors.NewInternalServerError(errorWritingFile)
		}

		if strings.TrimSpace(strings.ToLower(record[1])) == strings.TrimSpace(strings.ToLower(name)) {
			return record[0], nil
		}
	}

	return "", _errors.NewNotFoundError(fmt.Sprintf("character with name %s not found", name))
}

func readCharactersFromCSV() bool {
	// empty map
	charactersMap = make(map[string]*model.Character)

	file, err := os.Open(viper.GetString("db.charactersPath"))
	defer file.Close()

	if err != nil {
		isCsvFetched = false
		return isCsvFetched
	}

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return false
		}

		parseCharacter(record)
	}

	err = createMapTable()
	if err != nil {
		isCsvFetched = false
		return isCsvFetched
	}

	isCsvFetched = true
	return isCsvFetched
}

func parseCharacter(record []string) {
	var id string
	ch := &model.Character{}

	for pos, value := range record {
		switch pos {
		case chID:
			var err error
			id = value
			ch.ID, err = strconv.Atoi(value)
			if err != nil {
				continue
			}
		case chName:
			ch.Name = value
		case chStatus:
			ch.Status = value
		case chSpecies:
			ch.Species = value
		case chType:
			ch.Type = value
		case chGender:
			ch.Gender = value
		case chOriginName:
			ch.Origin.Name = value
		case chOriginUrl:
			ch.Origin.URL = value
		case chLocationName:
			ch.Location.Name = value
		case chLocationUrl:
			ch.Location.URL = value
		case chImage:
			ch.Image = value
		case chEpisodes:
			episodes := strings.Split(value, "+")
			ch.Episodes = episodes
		}
	}

	charactersMap[id] = ch
}

func createMapTable() _errors.RestError {
	file, err := os.Create(viper.GetString("db.mapPath"))
	if err != nil {
		return _errors.NewInternalServerError(errorWritingFile)
	}
	// ignoring close error it's safe on this point: https://www.joeshaw.org/dont-defer-close-on-writable-files/
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, ch := range charactersMap {
		var row []string
		row = append(row, strconv.Itoa(ch.ID))
		row = append(row, ch.Name)

		if err := writer.Write(row); err != nil {
			return _errors.NewInternalServerError(errorWritingFile)
		}
	}
	return nil
}
