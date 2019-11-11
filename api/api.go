package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sathvik777/go-api-skeleton/db"
	"github.com/Sathvik777/go-api-skeleton/httpbody"
)

type MessageAPI struct {
	DBOps db.Ops
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
	witeResponse(w, nil, http.StatusOK)
}

func (a *MessageAPI) Update(w http.ResponseWriter, r *http.Request) {
	witeResponse(w, nil, http.StatusOK)
}

func (a *MessageAPI) Delete(w http.ResponseWriter, r *http.Request) {
	witeResponse(w, nil, http.StatusOK)
}
