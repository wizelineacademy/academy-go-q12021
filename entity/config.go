package entity

// Config contains aditional information for customizing runtime aspects
type Config struct {
	CSVFilepath string // Location of the CSV file.
	Port        string // API Port
}

// ConfigStorage manages persistency access to configuration information
type ConfigStorage interface {
	Load() (Config, error)
}
