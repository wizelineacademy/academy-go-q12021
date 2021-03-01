package event

import (
	"time"

	"github.com/maestre3d/academy-go-q12021/internal/service"
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

func NewMovieCreated(id valueobject.MovieID, name valueobject.DisplayName,
	year valueobject.ReleaseYear, directors ...valueobject.DisplayName) MovieCreated {
	return MovieCreated{
		ID:          string(id),
		DisplayName: string(name),
		Directors:   service.MarshalDirectorsPrimitive(directors...),
		ReleaseYear: int(year),
		CreateTime:  time.Now().UTC(),
	}
}

func (c MovieCreated) Kind() string {
	return "movie-created"
}

func (c MovieCreated) AggregateID() string {
	return c.ID
}
