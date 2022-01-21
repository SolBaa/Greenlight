package handlers

import (
	"log"
	"net/http"

	cerror "github.com/SolBaa/Greenlight/pkg/error"
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
	err := validations.WriteJSON(w, http.StatusOK, validations.Envelope{"data": data}, nil)
	if err != nil {
		cerror.ServerErrorResponse(w, r, err)
	}

}
