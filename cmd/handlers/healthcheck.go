package handlers

import (
	"fmt"
	"log"
	"net/http"
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
	fmt.Fprintln(w, "Status available: ")
	fmt.Fprintf(w, "enviroment: %s\n", app.Config.Env)
	fmt.Fprintf(w, "version: %s\n", Version)

}
