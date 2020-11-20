package db

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/viper"
	"golang-bootcamp-2020/domain/model"
	_errors "golang-bootcamp-2020/utils/error"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	isCsvFetched  = false
	charactersMap map[string]*model.Character
)

const (
	chId = iota
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

type DataBaseRepository interface {
	CreateCharactersCSV(characters []model.Character) _errors.RestError
	GetCharacterFromId(id string) (*model.Character, _errors.RestError)
	GetCharacters() ([]model.Character, _errors.RestError)
	GetCharacterIdByName(name string) (string, _errors.RestError)
}

func Init() {
	isCsvFetched = readCharactersFromCSV()
}

func NewDbRepository() DataBaseRepository {
	Init()
	return &dbRepository{}
}

func (db *dbRepository) CreateCharactersCSV(characters []model.Character) _errors.RestError {
	file, err := os.Create("./resources/characters.csv")
	// TODO: handle this error
	defer file.Close()
	defer readCharactersFromCSV()

	if err != nil {
		return _errors.NewInternalServerError(errorWritingFile)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, ch := range characters {
		var row []string
		row = append(row, strconv.Itoa(ch.Id))
		row = append(row, ch.Name)
		row = append(row, ch.Status)
		row = append(row, ch.Species)
		row = append(row, ch.Type)
		row = append(row, ch.Gender)
		row = append(row, ch.Origin.Name)
		row = append(row, ch.Origin.Url)
		row = append(row, ch.Location.Name)
		row = append(row, ch.Location.Url)
		row = append(row, ch.Image)

		var episodesString string
		for _, episode := range ch.Episodes {
			episodesString += fmt.Sprintf("+%s", episode)
		}

		row = append(row, episodesString)
		//TODO: handle error
		writer.Write(row)
	}
	return nil
}

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
		return false
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
		return false
	}

	isCsvFetched = true
	return true
}

func parseCharacter(record []string) {
	var id string
	ch := &model.Character{}

	for pos, value := range record {
		switch pos {
		case chId:
			var err error
			id = value
			ch.Id, err = strconv.Atoi(value)
			//TODO: handle this error correctly
			if err != nil {
				panic(err)
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
			ch.Origin.Url = value
		case chLocationName:
			ch.Location.Name = value
		case chLocationUrl:
			ch.Location.Url = value
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
		row = append(row, strconv.Itoa(ch.Id))
		row = append(row, ch.Name)

		//TODO: handle error
		writer.Write(row)
	}
	return nil
}
