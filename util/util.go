package util

import (
	"time"
)

// IsStringTimestamp will use time.Parse to try to verify that a string is a timestamp
func IsStringTimestamp(input string) bool {
	_, err := time.Parse("2006-01-02 15:04:05.999", input)
	if err != nil {
		return false
	}
	return true
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
