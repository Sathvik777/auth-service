package api

import (
	"net/http"
)

// Health contains the config for the health api
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
