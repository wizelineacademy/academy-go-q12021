package service

import (
	"github.com/JosueSdev/golang-bootcamp-2020/interface/model"

	"gorm.io/gorm"
)

type postService struct {
	db *gorm.DB
}

//PostService is an interface to communicate with the post's service
type PostService interface {
	GetAll() ([]*model.PostModel, error)
}

//NewPostService constructs a new PostService
func NewPostService(db *gorm.DB) PostService {
	return &postService{db}
}

func (ps *postService) GetAll() ([]*model.PostModel, error) {
	var posts []*model.PostModel

	result := ps.db.Find(&posts)

	if err := result.Error; err != nil {
		return nil, err
	}

	return posts, nil
}
