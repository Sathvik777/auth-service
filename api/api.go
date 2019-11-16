package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sathvik777/go-api-skeleton/db"
	"github.com/Sathvik777/go-api-skeleton/httpbody"
)

type MessageAPI struct {
	DBOps db.DbOpsImpl
}

func (a *MessageAPI) Get(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/messages/")
	if idStr == "" {
		a.getAll(w, r)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Could not string cov")
	}
	response := httpbody.BasicResponse{Id: id}
	witeResponse(w, response, http.StatusOK)
}

func (a *MessageAPI) getAll(w http.ResponseWriter, r *http.Request) {
	witeResponse(w, nil, http.StatusOK)
}

func (a *MessageAPI) Create(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var req httpbody.MessageRequest
	err := decoder.Decode(&req)
	if err != nil {
		log.Panicf("Failed to decode %d", err)
		http.Error(w, "Unable to decode request", http.StatusInternalServerError)
		return
	}

	id, err := a.DBOps.InsertMessage(req)
	if err != nil {
		log.Panicf("Failed to insert %d", err)
		http.Error(w, "Unable to decode request", http.StatusInternalServerError)
		return
	}
	rep := httpbody.BasicResponse{Id: int(id)}
	witeResponse(w, rep, http.StatusOK)
}

func (a *MessageAPI) Update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req httpbody.MessageRequest
	err := decoder.Decode(&req)
	if err != nil {
		log.Panicf("Failed to decode %d", err)
		http.Error(w, "Unable to decode request", http.StatusInternalServerError)
		return
	}

	err = a.DBOps.UpdateMessage(req)
	if err != nil {
		log.Panicf("Failed to Update %d", err)
		http.Error(w, "Unable to decode request", http.StatusInternalServerError)
		return
	}

	witeResponse(w, nil, http.StatusOK)
}

func (a *MessageAPI) Delete(w http.ResponseWriter, r *http.Request) {
	witeResponse(w, nil, http.StatusOK)
}
