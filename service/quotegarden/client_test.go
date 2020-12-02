package quotegarden

import "testing"

// Test if client is able to get a quote from quote-garden
func TestGetQuote(t *testing.T) {
	c := NewClient()
	if _, err := c.GetQuote(); err != nil {
		t.Error(err)
	}
}
