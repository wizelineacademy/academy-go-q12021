package marshal

import (
	"strconv"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

const (
	movieTotalFields                  = 5
	movieDirectorsDelimiterPatternCSV = ","
)

// UnmarshalMovieCSV parses the given csv data into a Movie
func UnmarshalMovieCSV(movie *aggregate.Movie, records ...string) (err error) {
	if len(records) < movieTotalFields {
		return ErrCannotParseMovie
	}
	movie.ID = valueobject.MovieID(records[0])
	movie.DisplayName = valueobject.DisplayName(records[1])
	movie.Directors = unmarshalMovieDirectors(movieDirectorsDelimiterPatternCSV, records[2])
	year, err := strconv.Atoi(records[3])
	if err != nil {
		return err
	}
	movie.ReleaseYear, err = valueobject.NewReleaseYear(year) // avoid integer overflow at runtime
	if err != nil {
		return err
	}
	movie.IMDbID = valueobject.MovieID(records[4])
	return nil
}
