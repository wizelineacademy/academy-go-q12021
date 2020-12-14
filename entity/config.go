package entity

import "fmt"

// Config contains aditional information for customizing runtime aspects
type Config struct {
	CSVFilepath string
	Port        string
}

// ConfigStorage manages persistency access to configuration information
type ConfigStorage interface {
	Load() (Config, error)
}

// String prints the configuration values as a key, value string array
func (c Config) String() string {
	return fmt.Sprintf("[database: %q, api_port: %q]", c.CSVFilepath, c.Port)
}
