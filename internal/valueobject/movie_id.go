package valueobject

import "github.com/maestre3d/academy-go-q12021/internal/domain"

const (
	movieIDMinLength = 1
	movieIDMaxLength = 128 // in case of uuid
)

// ErrMovieIDOutOfRange the given movie ID is out of the defined range [1,128)
var ErrMovieIDOutOfRange = domain.NewOutOfRange("movie_id", movieIDMinLength, movieIDMaxLength)

// MovieID Movie's unique identifier
type MovieID string

// NewMovieID allocates a new valid MovieID
func NewMovieID(v string) (MovieID, error) {
	id := MovieID(v)
	if err := id.ensureLength(); err != nil {
		return "", err
	}

	return id, nil
}

func (i MovieID) ensureLength() error {
	if length := len(i); length < movieIDMinLength || length > movieIDMaxLength {
		return ErrMovieIDOutOfRange
	}

	return nil
}
