package internal

// ====================
//  IMPORTS
// ====================

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// ====================
//  TYPES
// ====================

type ChefServer struct {
	listener *http.Server
}

// ====================
//  CONSTRUCTOR
// ====================

func NewChefServer(host string, port int, handler http.Handler) *ChefServer {
	// Server
	return &ChefServer{
		listener: &http.Server{
			Addr:         fmt.Sprintf("%s:%d", host, port),
			IdleTimeout:  60 * time.Second,
			Handler:      handler,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
		},
	}
}

// ====================
//  STRUCTURE METHODS
// ====================

func (srv *ChefServer) Listen(red context.Context, stop context.CancelFunc) {
	// Lights
	yellow := make(chan error, 1)
	defer close(yellow)
	// Run
	log.Printf("|+| Server running on: http://%s\n", srv.listener.Addr)
	go func() {
		if err := srv.listener.ListenAndServe(); err != nil {
			select {
			case <-red.Done():
			case <-yellow:
			default:
				yellow <- err
			}
		}
	}()
	// Catch
	select {
	case err := <-yellow:
		stop()
		log.Printf("|-| Server suddently stopped: %s\n", err.Error())
	case <-red.Done():
		ctx, gracefulShutdown := context.WithTimeout(context.Background(), time.Minute)
		defer gracefulShutdown()
		log.Println("|*| Starting shutdown process (Max time: 60 seconds)")
		srv.listener.Shutdown(ctx)
		log.Println("|+| Server successfully stopped")
	}
}
