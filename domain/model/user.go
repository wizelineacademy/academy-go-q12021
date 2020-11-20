package model

import "time"

type User struct {
	ID        string
	Name      string
	Password  string
	CreatedAt time.Time
	IsAdmin   bool
}

func (user *User) HashPassword() {

}
