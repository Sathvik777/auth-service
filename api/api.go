package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Sathvik777/go-api-skeleton/db"
	"github.com/Sathvik777/go-api-skeleton/request"
	"github.com/sirupsen/logrus"

	"github.com/Sathvik777/go-api-skeleton/response"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type ServiceAPI struct {
	DBOps db.Ops `inject:""`
}

// InitRouter initializes the router with path specific handler functions
func (a *ServiceAPI) InitRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/signUp").
		Handler(
			negroni.New(negroni.HandlerFunc(a.signUp())))
}

func (a *ServiceAPI) signUp() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		signUpReq := request.SignUpRequest{}
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			logrus.Fatal(err)
		}

		err = json.Unmarshal(body, &signUpReq)
		if err != nil {
			logrus.Info(err)
			http.Error(w, "Unable to decode request", http.StatusInternalServerError)
			return
		}
		basicResponse := response.BasicResponse{}
		if len(signUpReq.Password) == 0 || len(signUpReq.Email) == 0 {

			basicResponse.Err = "Request parameter missing"
			basicResponseJson, err := json.Marshal(basicResponse)
			if err != nil {
				logrus.Error(err)
			}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(basicResponseJson)
			return
		}

		token, err := a.DBOps.InsertUser(signUpReq)
		if err != nil {
			// Silent fail
			logrus.Errorf("Log insert failed", err)
		}

		basicResponse.Token = token
		basicResponseJson, err := json.Marshal(basicResponse)
		if err != nil {
			http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(basicResponseJson)

	}
}
