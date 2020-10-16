package utils

// ====================
//  IMPORTS
// ====================

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// ====================
//  PUBLIC METHODS
// ====================

func GETBody(addr string) ([]byte, error) {
	// Request the data
	resp, err := genericRequest("GET", addr, nil, nil)
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s is not available", addr)
	}
	// Grab the content
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s does not respond correctly", addr)
	}
	return body, nil
}

func HEAD(addr string) (*http.Response, error) {
	return genericRequest("HEAD", addr, nil, map[string]string{"Connection": "close"})
}

func IsAlive(addr string) bool {
	conn, err := net.DialTimeout("tcp", addr, time.Second*10)
	if err == nil {
		conn.Close()
	}
	return err == nil
}

// ====================
//  PRIVATE METHODS
// ====================

func genericRequest(method string, addr string, body io.Reader, headers map[string]string) (*http.Response, error) {
	ctx, clean := context.WithTimeout(context.Background(), 10*time.Second)
	defer clean()
	req, err := http.NewRequestWithContext(ctx, method, addr, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Masterchef!")
	for header, val := range headers {
		req.Header.Set(header, val)
	}
	if err != nil {
		return nil, err
	}
	var client http.Client
	return client.Do(req)
}
