package backend

// ====================
//  IMPORTS
// ====================

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// ====================
//  STRUCTS
// ====================

// Declare configuration
type configuration = struct {
	Host string
	Port int
	ctx  *context.Context
	src  *template.Template
	wait time.Duration
}

// ====================
//  GLOBALS
// ====================

// Init configuration
var server = &configuration{
	Host: "[::1]",
	Port: 9999,
	wait: time.Minute,
}

// ====================
//  PUBLIC METHODS
// ====================

// Serve is a method that configures and runs the web server
func Serve() error {
	// Signal
	catch := make(chan os.Signal, 1)
	defer close(catch)
	signal.Notify(catch, os.Interrupt)
	// Router
	router := mux.NewRouter().
		StrictSlash(true)
	// SPA
	router.Path("/").
		Methods(http.MethodGet).
		HandlerFunc(handlerDefault)
	// Status
	router.Path("/api/status").
		Methods(http.MethodGet).
		HandlerFunc(handlerStatus)
	// WebSockets
	router.Path("/ws").
		Methods(http.MethodGet).
		HandlerFunc(handlerWebSockets)
	// Templates
	source, err := template.New("index").Parse(source)
	if err != nil {
		return errors.New("Cannot compile index.html")
	}
	server.src = source
	// Server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", server.Port),
		Handler:      router,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	// Runner method
	srverr := make(chan struct{}, 1)
	defer close(srverr)
	go func() {
		log.Printf("[+] Server running on: http://%s:%d/\n", server.Host, server.Port)
		if err = srv.ListenAndServe(); err != nil {
			srverr <- struct{}{}
		}
	}()
	// Wait
	select {
	case <-catch:
		// Context
		ctx, cancel := context.WithTimeout(context.Background(), server.wait)
		defer cancel()
		// Shutdown
		srv.Shutdown(ctx)
		log.Println("[*] Interrupt received")
	case <-srverr:
		return errors.New("Server stopped")
	}

	return nil
}
