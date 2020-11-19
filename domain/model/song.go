package model

import "time"

type Song struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Interpreter string    `json:"interpreter"`
	Length      time.Time `json:"length"`
	Lyric       string    `json:"lyric"`
}
