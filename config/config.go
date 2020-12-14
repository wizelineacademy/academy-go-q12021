package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type Config struct {
	DSN      string
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	DB       *sql.DB
	Addr     *string
	CSVFile  *os.File
}

// Init the application's configuration
func Init(infoLog, errorLog *log.Logger, csvFile *os.File) (*Config, error) {

	//Read config file
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Get the server address/port
	addr := viper.GetString("addr")

	// Create the Data Source Name
	// *We need to use the parseTime=true parameter in our
	// DSN to force it to convert TIME and DATE fields to time.Time. Otherwise it returns these as
	// []byte objects.
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	dbHost := viper.GetString("db.host")
	dbPort := viper.GetString("db.port")
	dbName := viper.GetString("db.name")
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", username, password, dbHost, dbPort, dbName)

	// Open db connection
	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatalf("message: unable to open db connection, type: database, err: %v", err)
		return nil, err
	}

	// Retrun the config
	return &Config{dsn, infoLog, errorLog, db, &addr, csvFile}, nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//Check if the DB is responding
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
