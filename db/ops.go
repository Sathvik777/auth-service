package db

import (
	"database/sql"
	"os/exec"

	"github.com/Sathvik777/auth-service/request"
	"github.com/sirupsen/logrus"
)

type DbOps interface {
	InsertUser(request request.BlobRequest) error
}

type DbOpsImpl struct {
	DbClient *sql.DB `inject:""`
}

var _ DbOps = &DbOpsImpl{}

func (ops *DbOpsImpl) InsertUser(request request.SignUpRequest) string, error {

	uuid, err := exec.Command("uuidgen").Output()

	var email = "'" + request.email + "'"
	var password = "'" + request.password + "'"
	var token = "'" + uuid + "'"

	var sqlInsertQuery = "INSERT INTO USERS (email, password, token) VALUES (" + email + ", " + password + ", " + token + " )"
	if _, err := ops.DbClient.Exec(sqlInsertQuery); err != nil {
		logrus.Errorln("DB INSERT ERROR : ", err)
		return nil , err
	}
	return token, nil
}
