package db

import (
	"database/sql"
	"os/exec"

	"github.com/Sathvik777/go-api-skeleton/request"
	"github.com/sirupsen/logrus"
)

type DbOps interface {
	InsertUser(request request.SignUpRequest) (string, error)
}

type DbOpsImpl struct {
	DbClient *sql.DB `inject:""`
}

var _ DbOps = &DbOpsImpl{}

func (ops *DbOpsImpl) InsertUser(request request.SignUpRequest) (string, error) {

	uuid, err := exec.Command("uuidgen").Output()

	if err != nil {
		logrus.Errorln("No token created : ", err)
		return "", err
	}

	var email = "'" + request.Email + "'"
	var password = "'" + request.Password + "'"
	var token = "'" + string(uuid) + "'"

	var sqlInsertQuery = "INSERT INTO USERS (email, password, token) VALUES (" + email + ", " + password + ", " + token + " )"
	if _, err := ops.DbClient.Exec(sqlInsertQuery); err != nil {
		logrus.Errorln("DB INSERT ERROR : ", err)
		return "", err
	}
	return token, nil
}
