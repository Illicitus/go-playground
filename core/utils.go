package core

import (
	"log"
	"net/http"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func JsonResponce(w http.ResponseWriter, js []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
}

func JsonBadRequestErrorHandler(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}
	return false
}

func JsonInternalServerErrorHandler(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func JsonUnauthorizedErrorHandler(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return true
	}
	return false
}

func GetJwtSecretKey() []byte {
	return []byte("SECRET")
}
