package backend

// ====================
//  IMPORTS
// ====================

import (
	"context"
	"errors"
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
	ctx  *context.Context
	src  *template.Template
	wait time.Duration
}

// ====================
//  GLOBALS
// ====================

// Init configuration
var server = &configuration{
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
		Addr:         ":4141",
		Handler:      router,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	// Runner method
	srverr := make(chan struct{}, 1)
	defer close(srverr)
	go func() {
		log.Println("[+] Server running on: http://[::1]:4141/")
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
