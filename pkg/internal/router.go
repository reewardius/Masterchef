package internal

// ====================
//  IMPORTS
// ====================

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"github.com/cosasdepuma/masterchef/pkg/utils"
)

// ====================
//  GLOBALS
// ====================

var wsUpgrader = websocket.Upgrader{}

// ====================
//  CONSTRUCTOR
// ====================

func NewRouter(source *template.Template, opts map[string]interface{}) *mux.Router {
	// Configuration
	router := mux.NewRouter().
		StrictSlash(true)
	// Handlers
	// -- Alive
	router.Path("/_alive").
		Methods(http.MethodGet, http.MethodHead).
		HandlerFunc(handlerAlive)
	// -- Kitchen
	router.Path("/_kitchen").
		Methods(http.MethodGet).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlerKitchenWS(w, r, opts)
		})
	// -- SPA
	router.Path("/").
		Methods(http.MethodGet, http.MethodHead).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handlerIndex(w, r, source) })
	// Router
	return router
}

// ====================
//  PRIVATE METHODS
// ====================

func handlerIndex(w http.ResponseWriter, r *http.Request, source *template.Template) {
	// Headers
	w.Header().Set("X-Powered-By", "Casseroles")
	w.WriteHeader(http.StatusOK)
	// Content
	source.Execute(w, struct {
		Addr string
	}{
		Addr: r.Host,
	})
}

func handlerAlive(w http.ResponseWriter, _ *http.Request) {
	// Headers
	w.Header().Set("X-Powered-By", "Knifes")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Content
	json.NewEncoder(w).Encode(map[string]bool{"alive": true})
}

func handlerKitchenWS(w http.ResponseWriter, r *http.Request, opts map[string]interface{}) {
	// New Client
	client, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("|-| %s entered the restaurant but there was an error: %s\n", r.RemoteAddr, err.Error())
		return
	}
	defer client.Close()
	// Parse the orders
	for {
		// New order
		t, order, err := client.ReadMessage()
		if err != nil {
			log.Printf("|-| No more orders from %s\n", r.RemoteAddr)
			return
		}
		// Parse the order
		cmd, request, ok := utils.ParseWSMessage(order)
		if !ok {
			log.Printf("|-| There was an error in the order from %s\n", r.RemoteAddr)
			continue
		}
		// Cook
		var response []byte
		switch cmd {
		case "cook":
			response = utils.ToWSResponse(cookInFurnace(request, opts))
		case "cancel":
		}
		// Deliver the dish
		err = client.WriteMessage(t, response)
		if err != nil {
			log.Printf("|-| There was an error delivering dish to %s\n", r.RemoteAddr)
		}
	}
}
