package utils

// ====================
//  IMPORTS
// ====================

import (
	"bytes"
	"log"
)

// ====================
//  PUBLIC METHODS
// ====================

func ParseWSMessage(msg []byte) (string, []byte, bool) {
	// Check format
	if len(msg) < 3 || msg[0] != '#' || msg[1] != '/' {
		log.Printf("Awkward msg (%s)\n", msg)
		return "", nil, false
	}
	data := bytes.SplitN(msg[2:], []byte("/"), 2)
	if len(data) != 2 {
		log.Printf("No data in the msg\n")
		return "", nil, false
	}
	return string(data[0]), data[1], true
}

func ToWSResponse(response []string) []byte {
	str := ToString(response)
	return []byte(str)
}
