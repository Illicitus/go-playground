package core

import "net/http"

func JsonResponce(w http.ResponseWriter, js []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(js)
	ErrorHandler(err)
}

func JsonResponse204(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
