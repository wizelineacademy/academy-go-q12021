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
}{
	{inID: valueobject.MovieID("1"), inName: valueobject.DisplayName("The Place Beyond the Pines"), inYear: valueobject.ReleaseYear(2012),
		inDirectors: []valueobject.DisplayName{valueobject.DisplayName("Derek Cianfrance")}},
	{inID: valueobject.MovieID("2"), inName: valueobject.DisplayName("Fargo"), inYear: valueobject.ReleaseYear(1996),
		inDirectors: []valueobject.DisplayName{valueobject.DisplayName("Ethan Coen"), valueobject.DisplayName("Joel Coen")}},
	{inID: valueobject.MovieID(""), inName: valueobject.DisplayName(""), inYear: valueobject.ReleaseYear(0),
		inDirectors: []valueobject.DisplayName{}},
}

func TestNewMovie(t *testing.T) {
	for _, tt := range newMovieTestingSuite {
		t.Run("New movie", func(t *testing.T) {
			movie := NewMovie(tt.inID, tt.inName, tt.inYear, tt.inDirectors...)
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
