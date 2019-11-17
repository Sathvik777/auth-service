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
	"github.com/joho/godotenv"
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

	messageAPI := api.MessageAPI{DBOps: client}

	handleProduct := func(w http.ResponseWriter, r *http.Request) {
		log.Println("GOT here")
		log.Printf("Request type %s", r.Method)
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
	http.HandleFunc("/api/messages/", handleProduct)
}

//TODO: 
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
	if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
	configDir := os.Getenv("CONFIG_DIR")

	config := setUpConfig(configDir)
	dbClient := setupDb(config.Database)
	if err := db.Migrate(dbClient.DbClient, config.Database); err != nil {
		log.Fatal("Did not migrate")
	}

	setupRouting(dbClient)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Server.Port), nil); err != nil {
		log.Fatal("Serving exited with error")
	}
}
