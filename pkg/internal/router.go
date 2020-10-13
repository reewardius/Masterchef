package internal

// ====================
//  IMPORTS
// ====================

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"text/template"

	"github.com/cosasdepuma/masterchef/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// ====================
//  GLOBALS
// ====================

var wsUpgrader = websocket.Upgrader{}

// ====================
//  CONSTRUCTOR
// ====================

func NewRouter(source *template.Template, green chan string) *mux.Router {
	// Configuration
	router := mux.NewRouter().
		StrictSlash(true)
	// Handlers
	// -- Alive
	router.Path("/_alive").
		Methods(http.MethodGet, http.MethodHead).
		HandlerFunc(handlerAlive)
	// -- Hello
	router.Path("/_hello/{port:[0-9]{1,5}}").
		Methods(http.MethodHead).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handlerHello(w, r, green) })
	// -- Kitchen
	router.Path("/_kitchen").
		Methods(http.MethodGet).
		HandlerFunc(handlerKitchenWS)
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

func handlerHello(w http.ResponseWriter, r *http.Request, green chan string) {
	// New cooker
	vars := mux.Vars(r)
	if port, ok := vars["port"]; ok && len(r.RemoteAddr) > 0 {
		rport := regexp.MustCompile("^\\d{1,5}$")
		rhost := regexp.MustCompile("^(.+):\\d{1,5}$")
		host := rhost.FindStringSubmatch(r.RemoteAddr)
		port := rport.FindString(port)
		if len(host) == 2 && len(port) > 0 {
			cooker := fmt.Sprintf("%s:%s", host[1], port)
			log.Printf("[*] Someone wants to be a cooker: %s\n", cooker)
			go func() { green <- cooker }()
		}
	}
	// Headers
	w.Header().Set("X-Powered-By", "Forks")
	w.WriteHeader(http.StatusOK)
}

func handlerKitchenWS(w http.ResponseWriter, r *http.Request) {
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
		orderType, order, err := client.ReadMessage()
		if err != nil {
			log.Printf("|-| Cannot retrieve order from %s\n", r.RemoteAddr)
			continue
		}
		// Parse the order
		fmt.Println(string(order))
		cmd, data, ok := utils.ParseWSMessage(order)
		if !ok {
			log.Printf("|-| There was an error in the order from %s: ", r.RemoteAddr)
			continue
		}
		fmt.Println(orderType, cmd, data)
	}
}
