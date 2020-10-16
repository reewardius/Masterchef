package utils

// ====================
//  IMPORTS
// ====================

import (
	"fmt"
	"regexp"
	"strings"
)

// ====================
//  PUBLIC METHODS
// ====================

func ContainsString(arr []string, str string) (int, bool) {
	for i, val := range arr {
		if val == str {
			return i, true
		}
	}
	return -1, false
}

func SplitContentLines(in string) []string {
	pattern := regexp.MustCompile("[\\t\\r\\n]+")
	onenl := pattern.ReplaceAllString(in, "\n")
	return strings.Split(strings.TrimSpace(onenl), "\n")
}

func ToAddr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

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
