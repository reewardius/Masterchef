package core

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func NewRouterChef(source *template.Template) *mux.Router {
	// Configuration
	router := basicConfiguration()
	// Handlers
	// -- SPA
	router.Path("/").
		Methods(http.MethodGet, http.MethodHead).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Headers
			w.WriteHeader(http.StatusOK)
			w.Header().Set("X-Powered-By", "Casseroles")
			// Content
			source.Execute(w, struct {
				Addr string
			}{
				Addr: r.Host,
			})
		})
	// Router
	return router
}

func basicConfiguration() *mux.Router {
	// Configuration
	router := mux.NewRouter().
		StrictSlash(true)
	// Handlers
	// -- Status
	router.Path("/_alive").
		Methods(http.MethodGet, http.MethodHead).
		HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			// Headers
			w.WriteHeader(http.StatusOK)
			w.Header().Set("X-Powered-By", "Knifes")
			w.Header().Set("Content-Type", "application/json")
			// Content
			json.NewEncoder(w).Encode(map[string]bool{"alive": true})
		})
	// Router
	return router
}
