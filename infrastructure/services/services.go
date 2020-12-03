package services

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"golang-bootcamp-2020/config"
	"golang-bootcamp-2020/domain/model"
)

func (c *Client) GetStudentsService() ([]model.Student, error)  {
	//c.client.R()
	// open csv
	csvFile, err := os.Open(config.C.CsvPath.Path)
	checkError("Can't open csv",err)

	var students []model.Student
	csvReader := csv.NewReader(csvFile)
	csvReader.Comment = '#'
	for {
		//read csv
		dataRow, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		checkError("Can't read csv file",err)

		//parse str to int
		id, err := strconv.Atoi(dataRow[0])
		checkError("Can't convert ID to int",err)
		zip, err := strconv.Atoi(dataRow[6])
		checkError("Can't convert Zip to int",err)
		//fill struct with data
		student := model.Student{
			ID:       id,
			Name:     dataRow[1],
			LastName: dataRow[2],
			Gender:   dataRow[3],
			City:     dataRow[4],
			State:    dataRow[5],
			Zip:      zip,
			Email:    dataRow[7],
			Age:      dataRow[8],
		}
		//add struct student to array  student
		students = append(students, student)
	}
	return students, err
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
func (c *Client) ReadUrl()   {
	//consule url
	//unmarshall
	//writetoCSV
}
// write info to CSV
func (c *Client) WriteToCsv(students []model.Student){
	//file, err := os.Create(config.C.CsvPath.Path)
	//checkError("Cannot open file", err)
	//defer file.Close()
	//
	//writer := csv.NewWriter(file)
	//
	//defer writer.Flush()

	//for _, value := range students {
	//	var err, _ = writer.Write( value )
	//	checkError("Cannot write to file", err)
	//}
}
