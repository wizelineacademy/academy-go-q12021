package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
)

func TestSongToString(t *testing.T) {
	song := &model.Song{
		ID:            14519164,
		Name:          "Fiesta pagana 2.0",
		InterpreterID: 22635,
		Interpreter:   "Mägo de Oz",
		AlbumID:       1455876,
		Album:         "Fiesta pagana 2.0",
	}

	expected := "14519164,Fiesta pagana 2.0,22635,Mägo de Oz,1455876,Fiesta pagana 2.0"
	result := song.ToString()
	assert.Equal(t, expected, result)
}
