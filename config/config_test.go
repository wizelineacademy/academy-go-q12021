package config_test

import (
	"os"
	"testing"

	"github.com/alexis-aguirre/golang-bootcamp-2020/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	//Cleaning env variables
	os.Setenv("PORT", "")
	os.Setenv("DBHOST", "")
	os.Setenv("DBPORT", "")
	os.Setenv("DBNAME", "")

	Config := &config.C
	Config.ReadConfig()
	assert.Equal(t, "8080", Config.PORT)
	assert.Equal(t, "localhost", Config.DB.HOST)
	assert.Equal(t, "3306", Config.DB.PORT)
	assert.Equal(t, "srms", Config.DB.NAME)
}
