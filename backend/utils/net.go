package utils

// ====================
//  IMPORTS
// ====================

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ====================
//  PUBLIC METHODS
// ====================

func Get(base string, data string) ([]byte, error) {
	// Compose the URL
	api := fmt.Sprintf(base, data)
	// Request the data
	resp, err := http.Get(api)
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s is not available", api)
	}
	// Grab the content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s does not respond correctly", api)
	}
	return body, nil
}
