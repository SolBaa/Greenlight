package handlers

import (
	"log"
	"net/http"

	"github.com/SolBaa/Greenlight/pkg/validations"
)

const Version = "1.0.0"

type Config struct {
	Port int
	Env  string
}

type Application struct {
	Config Config
	Logger *log.Logger
}

func (app *Application) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.Config.Env,
		"version":     Version,
	}
	err := validations.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.Logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
