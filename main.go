package main

import (
	"net/http"
	"os"

	"github.com/Sathvik777/auth-service/db"

	"github.com/Sathvik777/auth-service/api"
	"github.com/Sathvik777/auth-service/liveness"

	"github.com/facebookgo/inject"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"github.com/Sathvik777/auth-service/middleware"
	"github.com/Storytel/go-logrus-fluentd"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func setupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Debug("Error loading .env files. ", err)
	}
}

func setupLogging() {
	log.SetFormatter(logrusfluentd.NewFormatter("logger-service"))
	if os.Getenv("ENV") == "dev" {
		log.SetFormatter(&log.TextFormatter{})
	}

	if level := os.Getenv("LOG_LEVEL"); len(level) > 0 {
		ll, err := log.ParseLevel(level)
		if err == nil {
			log.SetLevel(ll)
		} else {
			log.Info("Failed to parse log level. Will use default: %s", err.Error())
		}
	}
}

func setupRouting(app *App) *negroni.Negroni {
	router := mux.NewRouter().StrictSlash(true)

	app.HealthAPI.InitRouter(router)
	app.ReadinessAPI.InitRouter(router)
	app.ServiceAPI.InitRouter(router)

	n := negroni.New()
	n.Use(middleware.Logger{
		ExludePaths: []string{"/health", "/readiness"},
	})
	n.UseHandler(router)

	return n
}

// App contains the context and configuration for the project
type App struct {
	HealthAPI    *liveness.HealthAPI    `inject:""`
	ReadinessAPI *liveness.ReadinessAPI `inject:""`
	ServiceAPI   *api.ServiceAPI        `inject:""`
}

func bootstrapApp() App {
	g := inject.Graph{}
	g.Logger = log.StandardLogger()

	app := App{}
	DbClient, err := db.Init()

	if err != nil {
		log.Fatal("DB Initialize failed:", err)
	}

	err = g.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: DbClient},
		&inject.Object{Value: &db.DbOpsImpl{}},
	)
	if err != nil {
		log.Fatal(err)
	}

	if err = g.Populate(); err != nil {
		log.Fatal(err)
	}

	return app
}

func main() {
	setupEnv()
	setupLogging()
	app := bootstrapApp()

	router := setupRouting(&app)

	port := os.Getenv("SERV_PORT")
	log.Debugf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
