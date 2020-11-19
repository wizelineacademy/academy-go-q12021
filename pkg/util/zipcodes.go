package util

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/AlonSerrano/GolangBootcamp/pkg/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/text/encoding/charmap"
)

func GetAndSave(client *mongo.Client) *mongo.InsertManyResult {
	zipCodes, zipCodesModel := getCSVCodes()
	fmt.Println("Se han insertado %i codigos postales", len(zipCodesModel))
	dropZipCodes(client)
	mr := saveZipCodes(zipCodes, client)
	return mr
}

func dropZipCodes(client *mongo.Client) bool {
	collection := client.Database("bootcamp").Collection("ZipCodes")
	collection.Drop(context.TODO())
	fmt.Printf("Droppped table")
	return true
}

func SearchZipCodes(cp string, client *mongo.Client) []models.ZipCodeCSV {
	collection := client.Database("bootcamp").Collection("ZipCodes")
	findOpts := options.Find()
	var results []models.ZipCodeCSV
	filter := bson.D{{"codigoPostal", cp}}
	cur, err := collection.Find(context.TODO(), filter, findOpts)

	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var s models.ZipCodeCSV
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, s)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return results
}

func saveZipCodes(zip []interface{}, client *mongo.Client) *mongo.InsertManyResult {
	collection := client.Database("bootcamp").Collection("ZipCodes")
	insertManyResult, err := collection.InsertMany(context.TODO(), zip)
	if err != nil {
		log.Fatal(err)
	}
	return insertManyResult
}

func getCSVCodes() ([]interface{}, []models.ZipCodeCSV) {
	url := "https://www.correosdemexico.gob.mx/datosabiertos/cp/cpdescarga.txt"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	stream := charmap.ISO8859_1.NewDecoder().Reader(response.Body)
	zipCodes, zipCodesModel := csvToMap(stream)
	return zipCodes, zipCodesModel
}

func edoToISO(stateCode int) string {
	rows := []map[string]string{
		{"isoCode": "MX-AGU"},
		{"isoCode": "MX-BCN"},
		{"isoCode": "MX-BCN"},
		{"isoCode": "MX-BCS"},
		{"isoCode": "MX-CAM"},
		{"isoCode": "MX-COA"},
		{"isoCode": "MX-COL"},
		{"isoCode": "MX-CHP"},
		{"isoCode": "MX-CHH"},
		{"isoCode": "MX-CMX"},
		{"isoCode": "MX-DUR"},
		{"isoCode": "MX-GUA"},
		{"isoCode": "MX-GRO"},
		{"isoCode": "MX-HID"},
		{"isoCode": "MX-JAL"},
		{"isoCode": "MX-MEX"},
		{"isoCode": "MX-MIC"},
		{"isoCode": "MX-MOR"},
		{"isoCode": "MX-NAY"},
		{"isoCode": "MX-NLE"},
		{"isoCode": "MX-OAX"},
		{"isoCode": "MX-PUE"},
		{"isoCode": "MX-QUE"},
		{"isoCode": "MX-ROO"},
		{"isoCode": "MX-SLP"},
		{"isoCode": "MX-SIN"},
		{"isoCode": "MX-SON"},
		{"isoCode": "MX-TAB"},
		{"isoCode": "MX-TAM"},
		{"isoCode": "MX-TLA"},
		{"isoCode": "MX-VER"},
		{"isoCode": "MX-YUC"},
		{"isoCode": "MX-ZAC"},
	}
	return rows[stateCode]["isoCode"]
}

func csvToMap(reader io.Reader) ([]interface{}, []models.ZipCodeCSV) {
	var zipCodes []interface{}
	var zipCodesModel []models.ZipCodeCSV
	r := csv.NewReader(reader)
	r.Comma = '|'
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	rows := []map[string]string{}
	var header []string
	firstLine := true
	for {
		if firstLine == false {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			if header == nil {
				header = record
			} else {

				dict := map[string]string{}
				for i := range header {
					dict[header[i]] = record[i]
				}
				u := strings.Replace(uuid.New().String(), "-", "", -1)
				isoCode, err := strconv.Atoi(dict["c_estado"])
				if err != nil {
					continue
				}
				parsingZipCode := models.ZipCodeCSV{
					Id:           u,
					CodigoPostal: dict["d_codigo"],
					Estado:       dict["d_estado"],
					EstadoISO:    edoToISO(isoCode),
					Municipio:    dict["D_mnpio"],
					Ciudad:       dict["d_ciudad"],
					Barrio:       dict["d_asenta"],
				}
				zipCodes = append(zipCodes, parsingZipCode)
				zipCodesModel = append(zipCodesModel, parsingZipCode)
				rows = append(rows, dict)

			}
		} else {
			r.Read()
			firstLine = false
		}
	}
	return zipCodes, zipCodesModel
}
