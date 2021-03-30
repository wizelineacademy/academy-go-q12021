package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/grethelBello/academy-go-q12021/config"
	"github.com/grethelBello/academy-go-q12021/constant"
	"github.com/grethelBello/academy-go-q12021/model"
	"github.com/grethelBello/academy-go-q12021/model/errs"
)

var retriesCache = make(map[string]int)
var maxRetry, timeRetry int

// CsvSource is a module to get information from a CSV file. To init, indicate the path to the file
type CsvSource string

// GetData is an implementation of Source interface which returns a Data struct with the data from the CSV file
func (source CsvSource) GetData(csvConfig ...*model.SourceConfig) (*model.Data, error) {
	path := string(source)
	if len(csvConfig) > 0 {
		path = *&csvConfig[0].CsvConfig
	}

	file, err := os.Open(path)
	if err != nil {
		storageError := errs.StorageError{TechnicalError: err}
		return &model.Data{}, storageError
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		storageError := errs.StorageError{TechnicalError: err}
		return &model.Data{}, storageError
	}

	data := model.NewCsvData(lines)
	return data, nil
}

func (source CsvSource) SetData(generalData *model.Data) error {
	// Open the file to append at the end
	file, err := os.OpenFile(string(source), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return errs.StorageError{TechnicalError: err}
	}
	defer file.Close()

	csvData := generalData.CsvData
	for _, row := range csvData {
		line := strings.Join(row, ",")
		_, err := file.WriteString(fmt.Sprintf("%v\n", line))
		checkRetries(line, string(source), err)
	}

	if len(retriesCache) > 0 {
		if timeRetry == 0 {
			timeRetry = convertEnvVar(constant.CsvTimeRetryVarName)
		}

		go func() {
			time.Sleep(time.Second * time.Duration(timeRetry))
			source.retrySetData()
			return
		}()
	}
	return nil
}

func (source CsvSource) retrySetData() {
	dataRaw := make([][]string, len(retriesCache))
	dataIndex := 0
	for cacheKey := range retriesCache {
		line := strings.Split(cacheKey, ",")
		dataRaw[dataIndex] = line
		dataIndex++
	}
	csvData := model.NewCsvData(dataRaw)
	source.SetData(csvData)
}

func checkRetries(cacheKey, file string, err error) {
	if maxRetry == 0 {
		maxRetry = convertEnvVar(constant.CsvMaxRetryVarName)
	}
	counter, ok := retriesCache[cacheKey]

	if err != nil {
		if !ok || counter < maxRetry {
			retriesCache[cacheKey] = counter + 1
			log.Printf("Error writing '%v' in '%v' file: %v, retry: %v", cacheKey, file, err, counter+1)
		} else if counter >= maxRetry {
			log.Printf("Last retry writing '%v' in '%v' file: %v\n", cacheKey, file, err)
			delete(retriesCache, cacheKey)
		}
	} else if ok {
		log.Printf("Success writing '%v' in '%v' file, retry: %v\n", cacheKey, file, counter)
		delete(retriesCache, cacheKey)
	}
}

func convertEnvVar(envName string) int {
	envVar, getEnvError := config.GetEnvVar(envName)
	if getEnvError != nil {
		envVar = constant.DefaultMaxRetries
	}
	if convVal, convError := strconv.Atoi(envVar); convError == nil {
		return convVal
	}

	return 0
}
