package marshal

import (
	"strconv"

	"github.com/maestre3d/academy-go-q12021/internal/aggregate"
	"github.com/maestre3d/academy-go-q12021/internal/valueobject"
)

const movieDirectorDelimeterPatternOmdb = ", "

// MovieOmdb OMDb API Movie schema
type MovieOmdb struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Director string `json:"Director"`
	IMDbID   string `json:"imdbID"`
}

// UnmarshalMovieOmdb parses the given MovieOMDb into an aggregate Movie
func UnmarshalMovieOmdb(movieOmdb MovieOmdb, movie *aggregate.Movie) (err error) {
	movie.DisplayName = valueobject.DisplayName(movieOmdb.Title)
	movie.Directors = unmarshalMovieDirectors(movieDirectorDelimeterPatternOmdb, movieOmdb.Director)
	year, err := strconv.Atoi(movieOmdb.Year)
	if err != nil {
		return err
	}
	movie.ReleaseYear, err = valueobject.NewReleaseYear(year) // avoid integer overflow at runtime
	if err != nil {
		return err
	}
	movie.IMDbID = valueobject.MovieID(movieOmdb.IMDbID)
	return nil
}
