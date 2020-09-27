package utils

// ====================
//  IMPORTS
// ====================

import (
	"fmt"
	"strings"
)

// ====================
//  PUBLIC METHODS
// ====================

// ToString converts a string slice into a string
func ToString(data []string) string {
	result := ""
	for _, str := range data {
		result = fmt.Sprintf("%s\n%s", result, str)
	}
	return strings.TrimSpace(result)
}

// Unique removes duplicates from a string slice
func Unique(data []string) []string {
	unique := []string{}
	duplicates := make(map[string]int)
	// Iterate over all the data
	for _, d := range data {
		duplicates[d]++
		if duplicates[d] == 1 {
			unique = append(unique, d)
		}
	}
	return unique
}
