package model

import "strings"

var validTypeFilters = [2]string{"Odd", "Even"}

type TypeFilter string

func (tf TypeFilter) isValid() bool {
	homoVal := strings.ToLower(string(tf))
	for _, enumValue := range validTypeFilters {
		if strings.ToLower(enumValue) == homoVal {
			return true
		}
	}
	return false
}

func (tf TypeFilter) isOdd() bool {
	homoVal := strings.ToLower(string(tf))
	return homoVal == strings.ToLower(validTypeFilters[0])
}

func Odd() string {
	return validTypeFilters[0]
}

func Even() string {
	return validTypeFilters[1]
}
