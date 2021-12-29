package handlers

import (
	"fmt"
	"net/http"

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
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}
