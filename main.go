package main

import (
	"github.com/gorilla/mux"
	"go-playground/apps/accounts"
	"go-playground/core"
	"net/http"
)

func main() {

	// Collect routers
	r := mux.NewRouter()
	accounts.AddAppRouter(r.PathPrefix("/accounts").Subrouter())

	// Db connection
	db := core.DbConnect()

	// Setup db tables
	apps := []interface{}{
		(*accounts.User)(nil),
	}

	core.ErrorHandler(core.CreateSchema(db, apps))

	// Setup jwt
	r.Use(core.JwtAuthMiddleware(db))

	http.ListenAndServe(":5000", r)
	defer core.ErrorHandler(core.CloseDatabase())
}
