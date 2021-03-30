package event

import (
	"time"

	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// MovieCreated a movie was created
type MovieCreated struct {
	ID          string    `json:"movie_id"`
	DisplayName string    `json:"display_name"`
	Directors   []string  `json:"director"`
	ReleaseYear int       `json:"release_year"`
	CreateTime  time.Time `json:"create_time"`
}

// NewMovieCreated allocates a MovieCreated event
func NewMovieCreated(id valueobject.MovieID, name valueobject.DisplayName,
	year valueobject.ReleaseYear, directors ...valueobject.DisplayName) MovieCreated {
	return MovieCreated{
		ID:          string(id),
		DisplayName: string(name),
		Directors:   valueobject.MarshalDirectorsPrimitive(directors...),
		ReleaseYear: int(year),
		CreateTime:  time.Now().UTC(),
	}
}

// Kind returns the Event key
func (c MovieCreated) Kind() string {
	return "movie.created"
}

// AggregateID returns the aggregate ID from the current aggregage whose state was modified
func (c MovieCreated) AggregateID() string {
	return c.ID
}
