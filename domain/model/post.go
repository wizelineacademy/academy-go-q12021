package model

import "time"

//Post is a storage-agnostic definition of a post for the board
type Post struct {
	Message string
	History int
	Time    time.Time
}

//PostGetter allows to extract a Post from future compositions
type PostGetter interface {
	Get() Post
}

//Get implements PostGetter for Post
func (p Post) Get() Post {
	return p
}
