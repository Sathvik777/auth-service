package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBClient struct {
}

func Init() (*sql.DB, error) {

	serverName := os.Getenv("MYSQL_SERVER")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+serverName+")/"+dbName+"?charset=utf8&parseTime=true&multiStatements=true")
	if err != nil {
		log.Fatal("Cannot open database connection. ", err)
	}

	return db, nil
}
