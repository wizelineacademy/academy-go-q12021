package handler

import (
	"fmt"

	"github.com/JosueSdev/golang-bootcamp-2020/interface/service"

	"gorm.io/gorm"
)

type postHandler struct {
	service.PostService
}

//PostHandler details the available handling methods
type PostHandler interface {
	GetAll() error
}

//NewPostHandler constructs a handler injecting the required dependencies
func NewPostHandler(db *gorm.DB) PostHandler {
	return &postHandler{service.NewPostService(db)}
}

func (pc *postHandler) GetAll() error {
	posts, err := pc.PostService.GetAll()

	if err != nil {
		return err
	}

	for _, p := range posts {
		fmt.Println(p)
	}

	return nil
}
