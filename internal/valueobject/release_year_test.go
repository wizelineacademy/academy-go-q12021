package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var releaseYearTestingSuite = []struct {
	in  int
	exp error
}{
	{-1, ErrReleaseYearOutOfRange},
	{1799, ErrReleaseYearOutOfRange},
	{2101, ErrReleaseYearOutOfRange},
	{1800, nil},
	{2100, nil},
	{2013, nil},
}

func TestNewReleaseYear(t *testing.T) {
	for _, tt := range releaseYearTestingSuite {
		t.Run("New Release year", func(t *testing.T) {
			year, err := NewReleaseYear(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, int(year))
		})
	}
}

func BenchmarkNewReleaseYer(b *testing.B) {
	b.Run("Bench New Release year", func(b *testing.B) {
		var v ReleaseYear
		defer func() {
			if v != 0 {
			}
		}()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v, _ = NewReleaseYear(2021)
		}
	})
}
