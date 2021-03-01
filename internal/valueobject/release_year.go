package valueobject

import "github.com/maestre3d/academy-go-q12021/internal/domain"

const (
	releaseYearMinLength = 1800 // before this time, movies didn't exist
	releaseYearMaxLength = 2100
)

// ErrReleaseYearOutOfRange the given release year is out of the defined epoch range [1800,2100)
var ErrReleaseYearOutOfRange = domain.NewOutOfRange("release_year", releaseYearMinLength, releaseYearMaxLength)

// ReleaseYear the specific year a Movie was released to the public on its origin country
type ReleaseYear uint16

// NewReleaseYear allocates a new valid ReleaseYear
func NewReleaseYear(v int) (ReleaseYear, error) {
	year := ReleaseYear(0) // we must avoid assigning the incoming data from here to avoid signed integer overflow at runtime
	if err := year.ensureEpoch(v); err != nil {
		return 0, err
	}
	year = ReleaseYear(v)
	return year, nil
}

func (y ReleaseYear) ensureEpoch(v int) error {
	if v < releaseYearMinLength || v > releaseYearMaxLength {
		return ErrReleaseYearOutOfRange
	}

	return nil
}
