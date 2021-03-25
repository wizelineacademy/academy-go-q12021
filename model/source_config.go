package model

// SourceConfig contains configuration for any source
type SourceConfig struct {
	CsvConfig  string
	HttpConfig HttpData
}

// NewHttpConfig initializes a HTTP config
func NewHttpConfig(httpConfig HttpData) *SourceConfig {
	return &SourceConfig{HttpConfig: httpConfig}
}

// NewCsvConfig initializes a CSV config
func NewCsvConfig(csvConfig string) *SourceConfig {
	return &SourceConfig{CsvConfig: csvConfig}
}
