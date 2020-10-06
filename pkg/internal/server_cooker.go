package internal

// ====================
//  IMPORTS
// ====================

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/cosasdepuma/masterchef/pkg/utils"
)

// ====================
//  TYPES
// ====================

type CookerServer struct {
	chef string
	host string
	port int
}

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
		fmt.Println("Hi!")
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
			var data []byte
			for {
				buffer := make([]byte, 1024)
				_, err := conn.Read(buffer)
				if err != nil {
					if err.Error() == "EOF" {
						break
					} else {
						fmt.Println(err.Error())
						log.Println("|!| I cannot hear your orders, Masterchef...")
						return
					}
				}
				data = append(data, buffer...)
			}
			data = bytes.Trim(data, "\x00")
			switch string(data) {
			case "fired!":
				log.Println("|-| The masterchef has gone")
				return
			}
		}
	}()
	select {
	case <-red.Done():
		lstnr.Close()
		log.Println("|+| Server successfully stopped")
	}
}
