package valueobject

// MarshalDirectorsPrimitive parses the given directors value object slice into a primitive slice
func MarshalDirectorsPrimitive(directors ...DisplayName) []string {
	ds := make([]string, 0)
	for _, d := range directors {
		ds = append(ds, string(d))
	}
	return ds
}

// MarshalDirectorsString parses the given directors value objects into a single string
func MarshalDirectorsString(directors ...DisplayName) string {
	directorsStr := ""
	for i, d := range directors {
		directorsStr += string(d)
		if i < len(directors)-1 {
			directorsStr += ","
		}
	}
	return directorsStr
}
