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

//func GetDb(r *http.Request) *pg.DB {
//	return r.Context().Value("db").(*pg.DB)
//}

func JsonResponce(w http.ResponseWriter, js []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
}

func JsonBadRequestErrorHandler(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func JsonInternalServerErrorHandler(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetJwtSecretKey() []byte {
	return []byte("SECRET")
}
