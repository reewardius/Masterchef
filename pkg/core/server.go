package core

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	http *http.Server
}

func NewServer(host string, port int, handler http.Handler) *Server {
	// Server
	return &Server{
		http: &http.Server{
			Addr:         fmt.Sprintf("%s:%d", host, port),
			IdleTimeout:  60 * time.Second,
			Handler:      handler,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
		},
	}
}

func (srv Server) ListenAndServe(red chan os.Signal) {
	// Lights
	yellow := make(chan error, 1)
	defer close(yellow)
	// Run
	log.Printf("|+| Server running on: http://%s\n", srv.http.Addr)
	go func() {
		if err := srv.http.ListenAndServe(); err != nil {
			select {
			case <-yellow:
			default:
				yellow <- err
			}
		}
	}()
	// Catch
	select {
	case err := <-yellow:
		log.Printf("|-| Server suddently stopped: %s\n", err.Error())
	case <-red:
		ctx, gracefulShutdown := context.WithTimeout(context.Background(), time.Minute)
		defer gracefulShutdown()
		log.Println("|*| Starting shutdown process (Max time: 60 seconds)")
		srv.http.Shutdown(ctx)
		log.Println("|+| Server successfully stopped")
	}
}
