package datastore

import (
	"strings"

	"github.com/JosueSdev/golang-bootcamp-2020/config"
	"github.com/JosueSdev/golang-bootcamp-2020/domain/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type postModel struct {
	gorm.Model
	model.Post
}

//NewMysqlDB creates a new mysql database instance
func NewMysqlDB() (*gorm.DB, error) {
	var dsn strings.Builder

	dsn.WriteString(config.Database.User)
	dsn.WriteString(":")
	dsn.WriteString(config.Database.Password)
	dsn.WriteString("@tcp(")
	dsn.WriteString(config.Database.Address)
	dsn.WriteString(")/")
	dsn.WriteString(config.Database.DBName)
	dsn.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := gorm.Open(mysql.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
