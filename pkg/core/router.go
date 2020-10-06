package core

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

	"github.com/gorilla/mux"
)

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
