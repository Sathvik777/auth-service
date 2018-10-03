package liveness

import (
	"net/http"

	"github.com/gorilla/mux"
)

// HealthAPI contains the config for the health api
type HealthAPI struct {
}

// InitRouter initializes the router with path specific handler functions
func (a *HealthAPI) InitRouter(router *mux.Router) {
	router.
		Methods("GET").
		Path("/health").
		Name("health").
		Handler(
			http.HandlerFunc(a.health()))
}

func (a *HealthAPI) health() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
