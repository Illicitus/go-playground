package core

import "net/http"

func JsonResponse(w http.ResponseWriter, js []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(js)
	ErrorHandler(err)
}

func JsonResponse200(w http.ResponseWriter, js []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(js)
	ErrorHandler(err)
}

func JsonResponse201(w http.ResponseWriter, js []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write(js)
	ErrorHandler(err)
}

func JsonResponse204(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
