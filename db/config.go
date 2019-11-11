package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DbConfig struct {
	Port     string
	Name     string
	Password string
}

func Init(config DbConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:"+config.Password+"@tcp(localhost:"+config.Port+")/"+config.Name+"?charset=utf8&parseTime=true&multiStatements=true")
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot open database connection. ", err)
	}
	return db, nil
}
