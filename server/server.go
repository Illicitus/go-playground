package server

import (
	"github.com/gorilla/mux"
	"go-playground/core"
	"net/http"
	"time"
)

func Serve(r *mux.Router) {
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	core.ErrorHandler(srv.ListenAndServe())
}
