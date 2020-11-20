package model

import "time"

type Song struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	InterpreterID int       `json:"interpreterId"`
	Interpreter   string    `json:"interpreter"`
	AlbumID       int       `json:"albumId"`
	Album         string    `json:"album"`
	Length        time.Time `json:"length,omitempty"`
	Lyric         string    `json:"lyric"`
}
