package aggregate

import (
	"github.com/maestre3d/academy-go-q12021/internal/domain"
	"github.com/maestre3d/academy-go-q12021/internal/event"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// Movie is a story or event recorded by a camera as a set of moving images and shown in a theater or on television; a motion picture.
//
//	Implements AggregateRoot
type Movie struct {
	ID          valueobject.MovieID
	DisplayName valueobject.DisplayName
	Directors   []valueobject.DisplayName
	ReleaseYear valueobject.ReleaseYear
	IMDbID      valueobject.MovieID

	events []domain.Event
}

var (
	// ErrMovieNotFound the specified movie was not found
	ErrMovieNotFound = domain.NewNotFound("movie")
)

// NewMovie creates a Movie and pushes the respective domain event
func NewMovie(id valueobject.MovieID, name valueobject.DisplayName,
	year valueobject.ReleaseYear, imdbID valueobject.MovieID, directors ...valueobject.DisplayName) *Movie {
	return &Movie{
		ID:          id,
		DisplayName: name,
		Directors:   directors,
		ReleaseYear: year,
		IMDbID:      imdbID,
		events: []domain.Event{
			event.NewMovieCreated(id, name, year, directors...),
		},
	}
}

// NewEmptyMovie allocates a Movie with default values to avoid nil pointer references
func NewEmptyMovie() *Movie {
	return &Movie{events: make([]domain.Event, 0)}
}

// PullEvents returns all the domain events that had happened inside the current aggregate root
func (m *Movie) PullEvents() []domain.Event {
	eventsMemo := m.events
	m.events = []domain.Event{}
	return eventsMemo
}
