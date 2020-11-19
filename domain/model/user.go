package model

import "time"

type User struct {
	Name      string
	Password  string
	CreatedAt time.Time
}
