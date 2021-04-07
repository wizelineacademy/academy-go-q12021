package model

// Data is a struct to encapsulate the different raw results from the data sources
type Data struct {
	CsvData  [][]string
	HttpData string
}

// NewCsvData initializes the data got from a CsvSource
func NewCsvData(csvData [][]string) *Data {
	return &Data{CsvData: csvData}
}

// NewHttpData initializes the data got from a HttpSource
func NewHttpData(data string) *Data {
	return &Data{HttpData: data}
}
