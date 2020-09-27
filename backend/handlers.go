package backend

// ====================
//  IMPORTS
// ====================

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// ====================
//  GLOBALS
// ====================

var wsUpgrader = websocket.Upgrader{}

// ====================
//  PRIVATE METHODS
// ====================

func handlerDefault(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// Template
	server.src.Execute(w, server)
}

func handlerStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// JSON Response
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

// ----- Web Sockets -----

func handlerWebSockets(w http.ResponseWriter, r *http.Request) {
	// WebSocket
	client, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer client.Close()
	for {
		// Receive message
		mtype, raw, err := client.ReadMessage()
		if err != nil {
			log.Println("[-] Error in websocket")
			continue
		}
		// Parse the response
		if len(raw) < 3 || raw[0] != '#' || raw[1] != '/' {
			log.Printf("[-] The message is not a command: %s", raw)
			continue
		}
		raw = raw[2:]
		data := bytes.SplitN(raw, []byte("/"), 2)
		if len(data) != 2 {
			log.Printf("[-] No data in the message")
			continue
		}
		// Run the commands
		switch string(data[0]) {
		case "cook":
			// Execute the modules
			log.Printf("[*] Order: %s\n", raw)
			result, err := Runner(data[1])
			fmt.Println("Outside")
			fmt.Println(result)
			fmt.Println(err)
			log.Printf("[+] Execution completed")
			if err != nil {
				log.Println("[-] One or more errors found")
				result = err.Error()
			}
			log.Println(result)
			if err = client.WriteMessage(mtype, []byte(result)); err != nil {
				log.Printf("[-] Error sending message:\n%s\n", result)
			}
		case "cancel":
			// Cancel the execution
			log.Println("[+] Cancel command")
		}
	}
}
