package model

import (
	"fmt"
)

type Song struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	InterpreterID int    `json:"interpreterId"`
	Interpreter   string `json:"interpreter"`
	AlbumID       int    `json:"albumId"`
	Album         string `json:"album"`
	Length        int    `json:"length,omitempty"`
	Lyric         string `json:"lyric"`
}

func (song *Song) ToString() string {
	return fmt.Sprintf("%d,%s,%d,%s,%d,%s", song.ID, song.Name, song.InterpreterID, song.Interpreter, song.AlbumID, song.Album)
}
