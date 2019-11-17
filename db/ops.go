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

	tx, err := ops.DbClient.Begin()
	res, err := tx.Exec("INSERT messages (message) VALUES (?)", request.Message)
	if err != nil {
		log.Panic("Cannot Execute statement ", err)
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return res.LastInsertId()
}

func (ops *DbOpsImpl) UpdateMessage(request httpbody.MessageRequest) error {

	tx, err := ops.DbClient.Begin()
	_, err = tx.Exec("UPDATE messages SET message = ? WHERE id = ?", request.Message, request.Id)
	if err != nil {
		log.Panic("Cannot Execute statement", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (ops *DbOpsImpl) DeleteMessage(id int) error {

	return nil
}
