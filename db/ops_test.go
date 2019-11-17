package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Sathvik777/go-api-skeleton/httpbody"
)

func TestDbOpsImpl_InsertMessage(t *testing.T) {

	testReq := httpbody.MessageRequest{Message: "Hello"}
	// Creates sqlmock database connection and a mock to manage expectations.
	dbm, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbm.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT messages").WithArgs("Hello").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	dbOps := DbOpsImpl{DbClient: dbm}

	_, err = dbOps.InsertMessage(testReq)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}
