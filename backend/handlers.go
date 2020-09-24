package backend

// ====================
//  IMPORTS
// ====================

import (
	"encoding/json"
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
	server.src.Execute(w, "index")
}

func handlerStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// JSON Response
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func handlerWebSockets(w http.ResponseWriter, r *http.Request) {
	// WebSocket
	ws, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		// Receive message
		mtype, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("[-] Error in websocket")
			continue
		}
		log.Printf("[*] Order: %s\n", message)
		// Response message
		if err = ws.WriteMessage(mtype, message); err != nil {
			log.Printf("[-] Error sending message:\n%s\n", message)
		}
	}
}
