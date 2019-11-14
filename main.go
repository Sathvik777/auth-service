package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

func setUpConfig(filename string) Config {
	var config Config
	configYaml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("failed to find config file : %s", err)
	}
	if err = yaml.Unmarshal(configYaml, &config); err != nil {
		log.Fatalf("failed to unmarshal config file : %s", err)
	}
	return config
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
	http.HandleFunc("/api/message/", handleProduct)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// Block until we receive our signal.
	<-interruptChan
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

func main() {
	config := setUpConfig("config.yaml")
	dbClient := setupDb(config.Database)
	setupRouting(dbClient)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("Serving exited with error")
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}
