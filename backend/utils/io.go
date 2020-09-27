package utils

// ====================
//  IMPORTS
// ====================

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ====================
//  PUBLIC METHODS
// ====================

func ReadFile(path string) ([]string, error) {
	lines := []string{}
	// Read the file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return lines, fmt.Errorf("Can't read %s file", path)
	}
	str := string(file)
	str = strings.TrimSpace(str)
	lines = strings.Split(str, "\n")
	return lines, nil
}
