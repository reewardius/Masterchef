package utils

// ====================
//  IMPORTS
// ====================

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"
)

// ====================
//  PUBLIC METHODS
// ====================

func IsAlive(addr string) bool {
	conn, err := net.DialTimeout("tcp", addr, time.Second*10)
	if err == nil {
		conn.Close()
	}
	return err == nil
}

func HEAD(addr string) (*http.Response, error) {
	return genericRequest("HEAD", addr, nil, map[string]string{"Connection": "close"})
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
