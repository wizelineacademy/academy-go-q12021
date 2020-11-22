package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"digimons/config"
	"digimons/domain/model"
	"digimons/infraestructure/datastore"
	"digimons/infraestructure/router"
	"digimons/registry"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

func main() {
	// This is the methods used for the first delivery
	write()
	read()

	// From here it begins the clean architecture for the final delivery
	config.ReadConfig()

	db := datastore.NewDB(config.C.Sqlitedb.DBPath)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	defer sqlDB.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
}

// write Obtain data from an external API convert it to an array and save it into csv file
func write() {
	resp, err := http.Get(config.C.Sources.DigimonAPI)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bodyBytes)

	var DigimonStructArray []model.Digimon
	json.Unmarshal(bodyBytes, &DigimonStructArray)

	csvFile, err := os.Create(config.C.Dest.DigimonCSV)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	for _, digimon := range DigimonStructArray {
		var row []string
		row = append(row, digimon.Name)
		row = append(row, digimon.Level)
		row = append(row, digimon.Image)
		writer.Write(row)
	}
	// remember to flush!
	writer.Flush()
}

// read Takes the information from a csv file, convert it to an array of Digimon structure and convert it to json.
func read() {
	csvFile, err := os.Open(config.C.Dest.DigimonCSV)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var digimon model.Digimon
	var digimons []model.Digimon
	for _, rec := range records {
		digimon.Name = string(rec[0])
		digimon.Level = string(rec[1])
		digimon.Image = string(rec[2])
		digimons = append(digimons, digimon)
	}

	jsonData, err := json.Marshal(digimons)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(jsonData)
}
