package db

import (
	"database/sql"
	"log"

	"github.com/Sathvik777/go-api-skeleton/httpbody"
)

type Ops interface {
	GetMessage(id int) (httpbody.BasicResponse, error)
	InsertMessage(request httpbody.MessageRequest) (int64, error)
	UpdateMessage(request httpbody.MessageRequest) error
	DeleteMessage(id int) error
}

type DbOpsImpl struct {
	DbClient *sql.DB
}

var _ Ops = &DbOpsImpl{}

func (ops *DbOpsImpl) GetMessage(id int) (httpbody.BasicResponse, error) {

	return httpbody.BasicResponse{}, nil
}

func (ops *DbOpsImpl) InsertMessage(request httpbody.MessageRequest) (int64, error) {
	stmt, err := ops.DbClient.Prepare("INSERT messages (message) VALUES (?)")
	if err != nil {
		log.Panicf("Cannot create statement d%", err)
		return 0, err
	}

	res, err := stmt.Exec(request.Message)
	if err != nil {
		log.Panicf("Cannot Execute statement d%", err)
		return 0, err
	}

	return res.LastInsertId()
}

func (ops *DbOpsImpl) UpdateMessage(request httpbody.MessageRequest) error {

	stmt, err := ops.DbClient.Prepare("UPDATE messages SET message = ? WHERE id = ?")
	if err != nil {
		log.Panicf("Cannot create statement d%", err)
		return err
	}

	_, err = stmt.Exec(request.Message, request.Id)
	if err != nil {
		log.Panicf("Cannot Execute statement d%", err)
		return err
	}

	return nil
}

func (ops *DbOpsImpl) DeleteMessage(id int) error {

	return nil
}
