package model

import "time"

type User struct {
	ID        string
	Name      string
	Password  string
	CreatedAt time.Time
}

func (user *User) HashPassword() {

}
