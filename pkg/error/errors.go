package cerror

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SolBaa/Greenlight/pkg/utils"
)

type Config struct {
	Port int
	Env  string
}

type app struct {
	Config Config
	Logger *log.Logger
}

// The logError() method is a generic utils for logging an error message. Later in the
// book we'll upgrade this to use structured logging, and record additional information
// about the request including the HTTP method and URL.
func LogError(r *http.Request, err error) {
	// app.Logger.Println(err)

}

// The errorResponse() method is a generic utils for sending JSON-formatted error
// messages to the client with a given status code. Note that we're using an interface{}
// type for the message parameter, rather than just a string type, as this gives us
// more flexibility over the values that we can include in the response.
func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := utils.Envelope{"error": message}
	// Write the response using the writeJSON() utils. If this happens to return an
	// error then log it, and fall back to sending the client an empty response with a
	// 500 Internal Server Error status code.

	err := utils.WriteJSON(w, status, env, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

// The serverErrorResponse() method will be used when our Application encounters an
// unexpected problem at runtime. It logs the detailed error message, then uses the
// errorResponse() utils to send a 500 Internal Server Error status code and JSON
// response (containing a generic error message) to the client.
func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	message := "the server encountered a problem and could not process your request"
	ErrorResponse(w, r, http.StatusInternalServerError, message)
}

// The notFoundResponse() method will be used to send a 404 Not Found status code and
// JSON response to the client.
func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	ErrorResponse(w, r, http.StatusNotFound, message)
}

// The methodNotAllowedResponse() method will be used to send a 405 Method Not Allowed
// status code and JSON response to the client.
func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	ErrorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
