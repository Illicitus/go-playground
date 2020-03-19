package core

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonError struct {
	Err string `json:"error"`
}

// Return http response with json error message
func httpJsonError(w http.ResponseWriter, jsError []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	_, err := w.Write(jsError)
	ErrorHandler(err)
}

/* Handle error and if it not nil, write error message with status code 400
to response and return true */
func JsonErrorHandler400(w http.ResponseWriter, err error) bool {
	if err != nil {
		jsonError, err := json.Marshal(JsonError{Err: err.Error()})
		if err != nil {
			log.Panic(err)
		}

		httpJsonError(w, jsonError, http.StatusBadRequest)
		return true
	}
	return false
}

/* Handle error and if it not nil, write error message with status code 404
to response and return true */
func JsonErrorHandler404(w http.ResponseWriter) bool {
	jsonError, err := json.Marshal(JsonError{Err: "not found"})
	if err != nil {
		log.Panic(err)
	}

	httpJsonError(w, jsonError, http.StatusNotFound)
	return true
}

/* Handle error and if it not nil, write error message with status code 500
to response and return true */
func JsonErrorHandler500(w http.ResponseWriter, err error) bool {
	if err != nil {
		jsonError, err := json.Marshal(JsonError{Err: err.Error()})
		if err != nil {
			log.Panic(err)
		}

		httpJsonError(w, jsonError, http.StatusInternalServerError)
		return true
	}
	return false
}

/* Handle error and if it not nil, write error message with status code 401
to response and return true */
func JsonErrorHandler401(w http.ResponseWriter) bool {
	jsonError, err := json.Marshal(JsonError{Err: "unauthorized"})
	if err != nil {
		log.Panic(err)
	}

	httpJsonError(w, jsonError, http.StatusNotFound)
	return true
}
