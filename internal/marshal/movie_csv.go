package marshal

import (
	"strconv"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/domain"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

// ErrCannotParseMovie the given movie could not get parsed
var ErrCannotParseMovie = domain.NewInfrastructure("cannot parse movie")

const movieTotalFields = 4

// UnmarshalMovieCSV parses the given csv data into a Movie
func UnmarshalMovieCSV(movie *aggregate.Movie, records ...string) (err error) {
	if len(records) != movieTotalFields {
		return ErrCannotParseMovie
	}
	movie.ID = valueobject.MovieID(records[0])
	movie.DisplayName = valueobject.DisplayName(records[1])
	movie.Director = valueobject.DisplayName(records[2])
	year, _ := strconv.Atoi(records[3])
	movie.ReleaseYear, err = valueobject.NewReleaseYear(year) // avoid integer overflow at runtime
	if err != nil {
		return err
	}
	return nil
}
