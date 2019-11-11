package db

import (
	"database/sql"

	"github.com/Sathvik777/go-api-skeleton/httpbody"
)

type Ops interface {
	GetProduct(id int) (httpbody.BasicResponse, error)
	InsertProduct(request httpbody.MessageRequest) (string, error)
	UpdateProduct(request httpbody.MessageRequest) error
	DeleteProduct(id int) error
}

type DbOpsImpl struct {
	DbClient *sql.DB
}

var _ Ops = &DbOpsImpl{}

func (ops *DbOpsImpl) GetProduct(id int) (httpbody.BasicResponse, error) {

	return httpbody.BasicResponse{}, nil
}

func (ops *DbOpsImpl) InsertProduct(request httpbody.MessageRequest) (string, error) {

	return "", nil
}

func (ops *DbOpsImpl) UpdateProduct(request httpbody.MessageRequest) error {

	return nil
}

func (ops *DbOpsImpl) DeleteProduct(id int) error {

	return nil
}
