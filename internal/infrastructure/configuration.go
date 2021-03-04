package infrastructure

import (
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// Configuration kernel/global configuration using OS environment variables if prod and yaml config file for the rest
// stages
type Configuration struct {
	Application   string
	Stage         string
	Version       string
	HTTPAddress   string
	HTTPPort      int
	MoviesDataset string
	OmdbAPIKey    string
}

func init() {
	viper.SetDefault("movies.application", "org.neutrinocorp.go-movies")
	viper.SetDefault("movies.stage", DevStage)
	viper.SetDefault("movies.version", "1.0.0")
	viper.SetDefault("movies.http", "")
	viper.SetDefault("movies.http.port", 8081)
	viper.SetDefault("movies.dataset.file", "./data/movies/movies_dataset.csv")
	viper.SetDefault("movies.omdb.api.key", "[YOUR_OMDB_API_KEY]")
}

const (
	// ProdStage Production deployment stage
	ProdStage = "prod"
	// DevStage Development deployment stage
	DevStage = "dev"
)

// NewConfiguration creates a Configuration with default configs or from sources
func NewConfiguration() Configuration {
	viper.AutomaticEnv()
	return Configuration{
		Application:   viper.GetString("movies.application"),
		Stage:         viper.GetString("movies.stage"),
		Version:       viper.GetString("movies.version"),
		HTTPAddress:   viper.GetString("movies.http"),
		HTTPPort:      viper.GetInt("movies.http.port"),
		MoviesDataset: viper.GetString("movies.dataset.file"),
		OmdbAPIKey:    viper.GetString("movies.omdb.api.key"),
	}
}

// IsProd returns if current config stage is in production stage
func (c Configuration) IsProd() bool {
	return c.Stage == ProdStage
}

// IsDev returns if current config stage is in development stage
func (c Configuration) IsDev() bool {
	return c.Stage == DevStage
}

// MajorVersion returns the current major version
func (c Configuration) MajorVersion() int {
	major, err := strconv.Atoi(strings.Split(c.Version, ".")[0])
	if err != nil {
		return 0
	}

	return major
}

// ReleaseStage returns the current release stage
func (c Configuration) ReleaseStage() string {
	stage := strings.Split(c.Version, "-")
	if len(stage) < 2 {
		return ""
	}

	return stage[1]
}
