package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DbConfig struct {
	Port     string
	Name     string
	Password string
	Address  string
}

func Migrate(instance *sql.DB, config DbConfig) error {
	migrationConfig := mysql.Config{
		MigrationsTable: "schema_migration",
		DatabaseName:    config.Name,
	}

	driver, err := mysql.WithInstance(instance, &migrationConfig)
	if err != nil {
		log.Fatal("Cannot open migrate ", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:/Users/sathvikkatam/src/go-api-skeleton/sql/",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal("Cannot open migrate connection. ", err)
	}

	m.Steps(1)
	return nil
}

func Init(config DbConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:"+config.Password+"@tcp("+config.Address+":"+config.Port+")/"+config.Name+"?charset=utf8&parseTime=true&multiStatements=true")
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot open database connection. ", err)
	}
	return db, nil
}
