package config

import "testing"

// Test if the runtime config is populated
func TestNew(t *testing.T) {
	ok := true
	switch c := New(); "" {
	case c.CSVFilepath, c.Port:
		ok = false
	}
	if !ok {
		t.FailNow()
	}
}
