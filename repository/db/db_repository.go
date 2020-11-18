package db

import (
	"encoding/csv"
	"errors"
	"fmt"
	"golang-bootcamp-2020/domain/model"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	IsCsvFetched  = false
	CharactersMap map[string]model.Character
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
)

type dbRepository struct {
}

type DataBaseRepository interface {
	CreateCharactersCSV(characters []model.Character) error
}

func Init() {
	IsCsvFetched = readCharactersFromCSV()
}

func NewDbRepository() DataBaseRepository {
	Init()
	return &dbRepository{}
}

func (db *dbRepository) CreateCharactersCSV(characters []model.Character) error {
	file, err := os.Create("./resources/characters.csv")
	// TODO: handle this error
	defer file.Close()
	defer readCharactersFromCSV()

	if err != nil {
		return errors.New("error writing file")
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

func readCharactersFromCSV() bool {

	// empty map

	CharactersMap = make(map[string]model.Character)

	file, err := os.Open("./resources/characters.csv")
	if err != nil {
		IsCsvFetched = false
		return false
	}

	r := csv.NewReader(file)
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		parseCharacter(record)
	}

	IsCsvFetched = true
	return true
}

func parseCharacter(record []string) {
	var id string
	ch := model.Character{}

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

	CharactersMap[id] = ch
}
