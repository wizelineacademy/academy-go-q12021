package aggregate

import (
	"testing"

	"github.com/maestre3d/academy-go-q12021/internal/valueobject"

	"github.com/stretchr/testify/assert"
)

var newMovieTestingSuite = []struct {
	inID        valueobject.MovieID
	inName      valueobject.DisplayName
	inYear      valueobject.ReleaseYear
	inDirectors []valueobject.DisplayName
	inIMDbID    valueobject.MovieID
}{
	{inID: valueobject.MovieID("1"), inName: valueobject.DisplayName("The Place Beyond the Pines"), inYear: valueobject.ReleaseYear(2012),
		inDirectors: []valueobject.DisplayName{valueobject.DisplayName("Derek Cianfrance")}, inIMDbID: valueobject.MovieID("tt12345678")},
	{inID: valueobject.MovieID("2"), inName: valueobject.DisplayName("Fargo"), inYear: valueobject.ReleaseYear(1996),
		inDirectors: []valueobject.DisplayName{valueobject.DisplayName("Ethan Coen"), valueobject.DisplayName("Joel Coen")},
		inIMDbID:    valueobject.MovieID("tt12345678")},
	{inID: valueobject.MovieID(""), inName: valueobject.DisplayName(""), inYear: valueobject.ReleaseYear(0),
		inDirectors: []valueobject.DisplayName{}, inIMDbID: valueobject.MovieID("tt12345678")},
}

func TestNewMovie(t *testing.T) {
	for _, tt := range newMovieTestingSuite {
		t.Run("New movie", func(t *testing.T) {
			movie := NewMovie(tt.inID, tt.inName, tt.inYear, tt.inIMDbID, tt.inDirectors...)
			events := movie.PullEvents()
			assert.Equal(t, 1, len(events))
			assert.Equal(t, string(movie.ID), events[0].AggregateID())
		})
	}
}

func TestNewEmptyMovie(t *testing.T) {
	movie := NewEmptyMovie()
	assert.Equal(t, 0, len(movie.PullEvents()))
}
