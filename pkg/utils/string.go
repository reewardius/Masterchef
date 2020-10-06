package utils

// ====================
//  IMPORTS
// ====================

import "fmt"

// ====================
//  PUBLIC METHODS
// ====================

func ToAddr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func ContainsString(arr []string, str string) (int, bool) {
	for i, val := range arr {
		if val == str {
			return i, true
		}
	}
	return -1, false
}
