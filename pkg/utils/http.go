package utils

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

func HEAD(addr string) (*http.Response, error) {
	return genericRequest("HEAD", addr, nil)
}

func PUT(addr string, body string) (*http.Response, error) {
	return genericRequest("PUT", addr, strings.NewReader(body))
}

func genericRequest(method string, addr string, body io.Reader) (*http.Response, error) {
	ctx, clean := context.WithTimeout(context.Background(), 10*time.Second)
	defer clean()
	var client http.Client
	req, err := http.NewRequestWithContext(ctx, method, addr, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Masterchef!")
	req.Header.Set("Content-Type", "application/json")
	return client.Do(req)
}
