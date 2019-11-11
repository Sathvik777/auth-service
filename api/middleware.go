package api

import (
	"encoding/json"
	"net/http"
)

func witeResponse(w http.ResponseWriter, body interface{}, respCode int) {
	if body != nil {
		response, err := json.Marshal(body)
		if err != nil {
			http.Error(w, "Unable to decode request", http.StatusInternalServerError)
			return
		}
		w.Write(response)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(respCode)
}
