package data

// Source is an interface for the modules in charge of getting information from different sources (files, APIs, databases, etc.)
type Source interface {
	GetData() (Data, error)
}

// Data is a struct to encapsulate the different raw results from the data sources
type Data struct {
	CsvData [][]string
}

// NewCsvData initializes the data got from a CsvDataSource
func NewCsvData(csvData [][]string) Data {
	data := Data{CsvData: csvData}
	return data
}
