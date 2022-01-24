package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SolBaa/Greenlight/internal/data"
	"github.com/SolBaa/Greenlight/internal/validator"
	cerror "github.com/SolBaa/Greenlight/pkg/error"
	"github.com/SolBaa/Greenlight/pkg/utils"
)

func (app *Application) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := utils.ReadJSON(w, r, &input)
	if err != nil {
		cerror.BadRequestResponse(w, r, err)
		return
	}
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}
	// Initialize a new Validator.
	v := validator.New()
	// Call the ValidateMovie() function and return a response containing the errors if
	// any of the checks fail.
	if data.ValidateMovie(v, movie); !v.Valid() {
		cerror.FailedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *Application) ShowMovieHandler(w http.ResponseWriter, r *http.Request) {
	// params := httprouter.ParamsFromContext(r.Context())
	id, err := utils.ReadIDParam(r)
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
	err = utils.WriteJSON(w, http.StatusOK, utils.Envelope{"movie": movie}, nil)
	if err != nil {
		app.Logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
