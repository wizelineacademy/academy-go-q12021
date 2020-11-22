package main

import (
    "encoding/csv"
    "github.com/gin-gonic/gin"
    "io"
    "log"
    "os"
    "strconv"
)

const CSV_ERROR_MESSAGE = "There was an error when attempting to read the csv file. Please try again later."
const FILE_PATH = "covid.csv"

type Country struct {
    Name   string
    Cases  int
    Deaths int
}

func csvError(context *gin.Context, err error) {
    log.Println(err)
    context.JSON(500, gin.H{
        "message": CSV_ERROR_MESSAGE,
    })
}

func readCSV(context *gin.Context) {
    file, err := os.Open(FILE_PATH)
    log.Println("%T", err)
    if err != nil {
        csvError(context, err)
        return
    }

    var data []Country
    reader := csv.NewReader(file)
    reader.Read() // skip csv headers, maybe there's a better way...
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            csvError(context, err)
            return
        }

        cases, _ := strconv.Atoi(record[1])
        deaths, _ := strconv.Atoi(record[2])

        data = append(data, Country{
            Name:   record[0],
            Cases:  cases,
            Deaths: deaths,
        })
    }

    context.JSON(200, data)
}

func main() {
    r := gin.Default()
    r.GET("/readcsv", readCSV)
    r.Run()
}
