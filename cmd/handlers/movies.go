package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SolBaa/Greenlight/internal/data"
	"github.com/SolBaa/Greenlight/pkg/validations"
)

func (app *Application) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new movie")
}

func (app *Application) ShowMovieHandler(w http.ResponseWriter, r *http.Request) {
	// params := httprouter.ParamsFromContext(r.Context())
	id, err := validations.ReadIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}
	err = validations.WriteJSON(w, http.StatusOK, validations.Envelope{"movie": movie}, nil)
	if err != nil {
		app.Logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
