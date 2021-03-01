package valueobject

import "github.com/maestre3d/academy-go-q12021/internal/domain"

const (
	displayNameMinLength = 1
	displayNameMaxLength = 256
)

// ErrDisplayNameOutOfRange the given display name is out of the defined range [1,256)
var ErrDisplayNameOutOfRange = domain.NewOutOfRange("display_name", displayNameMinLength, displayNameMaxLength)

// DisplayName represents a simplified name of an entity to be shown at the presentation layer
type DisplayName string

// NewDisplayName allocates a new valid DisplayName
func NewDisplayName(v string) (DisplayName, error) {
	name := DisplayName(v)
	if err := name.ensureLength(); err != nil {
		return "", err
	}

	return name, nil
}

func (n DisplayName) ensureLength() error {
	if length := len(n); length < displayNameMinLength || length > displayNameMaxLength {
		return ErrDisplayNameOutOfRange
	}

	return nil
}
