package infrastructure

import (
	"log"
	"testing"

	"github.com/jesus-mata/academy-go-q12021/infrastructure/dto"
	"github.com/stretchr/testify/assert"
)

func TestGetAllLines(t *testing.T) {

	csv := NewCsvSource("../resources/data_test.csv", log.Default())
	lines, err := csv.GetAllLines()
	if err != nil {
		t.Errorf("Test failed due to %s", err)
	}

	uuids := []string{"f5433032-6d9b-4c28-a353-1517adb2fecd", "2d2bfd26-df9f-4568-b686-2779b9c9747e", "6e73ea5c-ce39-45fb-8aaf-9ff062698fe3"}

	for i, v := range lines {
		assert.Equal(t, v[0], uuids[i])
	}
}

func TestWriteLines(t *testing.T) {
	data := []dto.NewItem{{Id: "f5433032-6d9b-4c28-a353-1517adb2fecd", Title: "Title 1", Description: "Description 1", Url: "http://url.com", Author: "Some Ahthor 1", Image: "htpp://storage/image.png", Language: "en", Category: []string{"general"}, Published: "2021-03-17 15:57:47 +0000"},
		{Id: "2d2bfd26-df9f-4568-b686-2779b9c9747e", Title: "Title 2", Description: "Description 2", Url: "http://url.com", Author: "Some Ahthor 2", Image: "htpp://storage/image.png", Language: "en", Category: []string{"general"}, Published: "2021-03-17 15:57:47 +0000"},
		{Id: "6e73ea5c-ce39-45fb-8aaf-9ff062698fe3", Title: "Title 3", Description: "Description 3", Url: "http://url.com", Author: "Some Ahthor 3", Image: "htpp://storage/image.png", Language: "en", Category: []string{"general"}, Published: "2021-03-17 15:57:47 +0000"}}

	csvFileName := "../resources/data_test.csv"
	csv := NewCsvSource(csvFileName, log.Default())
	err := csv.WriteLines(data)
	if err != nil {
		t.Errorf("Test failed due to %s", err)
	}

	lines, _ := csv.GetAllLines()

	assert.Equal(t, 3, len(lines))

}
