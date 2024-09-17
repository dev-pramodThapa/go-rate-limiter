package utils

import (
	"strings"
)

// SplitPath splits a URL path into segments based on "/"
func SplitPath(path string) []string {
	return strings.Split(path, "/")
}

// IsStringEmpty checks if a string is empty
func IsStringEmpty(s string) bool {
	return len(s) == 0
}
