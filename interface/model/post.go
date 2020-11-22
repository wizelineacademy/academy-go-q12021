package model

import (
	"github.com/JosueSdev/golang-bootcamp-2020/domain/model"
)

//PostModel adapts the domain model to a gorm aware environment
type PostModel struct {
	ID uint
	model.Post
}

//TableName defines the table to lookup in the database
func (PostModel) TableName() string {
	return "post"
}

func (pm *PostModel) Get() model.Post {
	return pm.Post
}
