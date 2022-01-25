package handlers

import (
	"log"
	"net/http"

	"github.com/SolBaa/Greenlight/internal/data"
	cerror "github.com/SolBaa/Greenlight/pkg/error"
	"github.com/SolBaa/Greenlight/pkg/utils"
)

const Version = "1.0.0"

type Config struct {
	Port int
	Env  string
	DB   struct {
		DSN string
	}
}
type Application struct {
	Config Config
	Logger *log.Logger
	Models data.Models
}

func (app *Application) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.Config.Env,
		"version":     Version,
	}
	err := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"data": data}, nil)
	if err != nil {
		cerror.ServerErrorResponse(w, r, err)
	}

}
