package data

import (
	"encoding/csv"
	"go-api/models"
	"io"
	"os"
	"strconv"
)

func InitializeCSVData(filePath string) ([]models.CardBack, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	data := csv.NewReader(csvFile)
	cardBacks := []models.CardBack{}

	isHeader := true
	for {
		record, err := data.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if isHeader {
			isHeader = !isHeader
			continue
		}

		cardBack := models.CardBack{}
		cardBack.ID, err = strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		cardBack.SortCategory, err = strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		cardBack.Text = record[2]
		cardBack.Name = record[3]
		cardBack.Image = record[4]
		cardBack.Slug = record[5]

		cardBacks = append(cardBacks, cardBack)
	}

	return cardBacks, nil
}
