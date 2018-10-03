package liveness

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ReadinessAPI contains the config for the readiness api
type ReadinessAPI struct {

}

// InitRouter initializes the router with path specific handler functions
func (a *ReadinessAPI) InitRouter(router *mux.Router) {
	router.
		Methods("GET").
		Path("/readiness").
		Name("readiness").
		Handler(
			http.HandlerFunc(a.readiness()))
}

func (a *ReadinessAPI) readiness() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
