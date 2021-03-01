package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var movieIDTestingSuite = []struct {
	in  string
	exp error
}{
	{"", ErrMovieIDOutOfRange},
	{"", ErrMovieIDOutOfRange}, // will be above 128 char long
	{"1234567890123456", nil},
	{"", nil}, // will be 128 char long
	{"1", nil},
}

func TestNewMovieID(t *testing.T) {
	for i, tt := range movieIDTestingSuite {
		if i == 1 {
			tt.in = populateString(129)
		} else if i == 3 {
			tt.in = populateString(128)
		}

		t.Run("New movie id", func(t *testing.T) {
			id, err := NewMovieID(tt.in)
			if err != nil {
				assert.Equal(t, tt.exp, err)
				return
			}
			assert.Equal(t, tt.in, string(id))
		})
	}
}

func BenchmarkNewMovieID(b *testing.B) {
	id := "123"
	b.Run("Bench New movie id", func(b *testing.B) {
		var v MovieID
		defer func() {
			// avoids v non-used
			if v != "" {
			}
		}()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			v, _ = NewMovieID(id)
		}
	})
}
