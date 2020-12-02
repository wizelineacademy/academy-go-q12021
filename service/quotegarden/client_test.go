package quotegarden

import "testing"

func TestGetQuote(t *testing.T) {
	c := NewClient()
	if _, err := c.GetQuote(); err != nil {
		t.Error(err)
	}
}
