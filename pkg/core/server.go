package core

// ====================
//  IMPORTS
// ====================

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/cosasdepuma/masterchef/pkg/utils"
)

// ====================
//  TYPES
// ====================

type (
	Server interface {
		Listen(context.Context, context.CancelFunc)
	}

	ChefServer struct {
		listener *http.Server
	}

	CookerServer struct {
		chef string
		host string
		port int
	}
)

// ====================
// =                  =
// =      COOKER      =
// =      SERVER      =
// =                  =
// ====================

// ====================
//  CONSTRUCTOR
// ====================

func NewCookerServer(host string, port int, chef string) *CookerServer {
	var srv CookerServer
	// Chef alive
	if utils.IsAlive(chef) {
		srv.chef = chef
	}
	// Server
	srv.host = host
	srv.port = port
	return &srv
}

// ====================
//  STRUCTURE METHODS
// ====================

func (srv CookerServer) Listen(red context.Context, stop context.CancelFunc) {
	addr := utils.ToAddr(srv.host, srv.port)
	// Listener
	lstnr, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("|!| Cannot start the server: %s\n", addr)
		return
	}
	log.Printf("|+| Server running on: %s\n", addr)
	// Meet chef
	_, err = utils.HEAD(fmt.Sprintf("http://%s/_hello/%d", srv.chef, srv.port))
	if err != nil {
		log.Println("|!| Cannot meet the chef!")
		return
	}
	log.Printf("|*| Meeting the chef: %s\n", srv.chef)
	go func() {
		defer stop()
		conn, err := lstnr.Accept()
		if err != nil {
			if !strings.Contains(err.Error(), "use of closed network connection") {
				log.Printf("|!| Unexpected error: %s\n", err.Error())
			}
			return
		}
		defer conn.Close()
		log.Printf("|+| Cooking for: %s\n", conn.RemoteAddr())
		// Get data
		for {
			buffer := make([]byte, 1024)
			_, err := conn.Read(buffer)
			if err != nil {
				fmt.Println(err.Error())
				log.Println("|!| I cannot hear your orders, Masterchef...")
				return
			}
			fmt.Println(string(buffer))
			switch string(buffer) {
			case "a coffee":
				conn.Write([]byte("sure"))
			}
		}
	}()
	select {
	case <-red.Done():
		lstnr.Close()
	}
}

// ====================
// =                  =
// =       CHEF       =
// =      SERVER      =
// =                  =
// ====================

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
