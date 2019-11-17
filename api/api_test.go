package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Sathvik777/go-api-skeleton/db"
	"github.com/Sathvik777/go-api-skeleton/httpbody"
)

func setUp(t *testing.T) (error, *MessageAPI) {
	dbm, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbm.Close()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT messages").WithArgs("Hello").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	a := &MessageAPI{DBOps: db.DbOpsImpl{DbClient: dbm}}
	return err, a
}

func TestMessageAPI_Create(t *testing.T) {
	reqbody := httpbody.MessageRequest{Message: "Test"}
	write, err := json.Marshal(reqbody)
	err, a := setUp(t)
	req := httptest.NewRequest("POST", "http://localhost:8080/api/messages/", bytes.NewReader(write))
	w := httptest.NewRecorder()

	a.Create(w, req)
	resp := w.Body
	_, err = ioutil.ReadAll(resp)
	if err != nil {
		fmt.Println(err)
	}
}
