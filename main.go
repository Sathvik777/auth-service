package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Sathvik777/go-api-skeleton/api"
	"github.com/Sathvik777/go-api-skeleton/db"
	"gopkg.in/yaml.v2"
)

type serverConfig struct {
	Port int `yaml:"port"`
}

//Config is a yaml replication
type Config struct {
	Server   serverConfig `yaml:"server"`
	Database db.DbConfig  `yaml:"database"`
}

func setupDb(config db.DbConfig) db.DbOpsImpl {
	sqlCLient, err := db.Init(config)
	if err != nil {
		log.Fatal("Could not connect to DB")
	}
	return db.DbOpsImpl{DbClient: sqlCLient}
}

func setupRouting(client db.DbOpsImpl) {

	messageAPI := api.MessageAPI{}

	handleProduct := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			messageAPI.Get(w, r)
		case http.MethodPost:
			messageAPI.Create(w, r)
		case http.MethodPut:
			messageAPI.Update(w, r)
		case http.MethodDelete:
			messageAPI.Delete(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}

	http.HandleFunc("/health", api.Health)
	http.HandleFunc("/api/products/", handleProduct)
}

func main() {
	var config Config
	configYaml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("failed to find config file : %s", err)
	}
	if err = yaml.Unmarshal(configYaml, &config); err != nil {
		log.Fatalf("failed to unmarshal config file : %s", err)
	}
	dbClient := setupDb(config.Database)
	setupRouting(dbClient)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), nil); err != nil {
		log.Fatal("Serving exited with error")
	}
}
