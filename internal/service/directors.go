package service

import "github.com/maestre3d/academy-go-q12021/internal/valueobject"

// MarshalDirectorsPrimitive parses the given directors value object slice into a primitive slice
func MarshalDirectorsPrimitive(directors ...valueobject.DisplayName) []string {
	ds := make([]string, 0)
	for _, d := range directors {
		ds = append(ds, string(d))
	}
	return ds
}
