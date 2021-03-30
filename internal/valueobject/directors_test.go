package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var marshalDirectorStringTestingSuite = []struct {
	in  []DisplayName
	exp string
}{
	{[]DisplayName{"Ethan Coen", "Joel Coen"}, "Ethan Coen,Joel Coen"},
	{[]DisplayName{"Auguste Lumiere", "Louis Lumiere"}, "Auguste Lumiere,Louis Lumiere"},
	{[]DisplayName{"Paul Thomas Anderson"}, "Paul Thomas Anderson"},
	{[]DisplayName{}, ""},
}

func TestMarshalDirectorsString(t *testing.T) {
	for _, tt := range marshalDirectorStringTestingSuite {
		t.Run("MarshalDirectorsString", func(t *testing.T) {
			str := MarshalDirectorsString(tt.in...)
			assert.Equal(t, tt.exp, str)
		})
	}
}
