package main

import (
	"github.com/gorilla/mux"
	"go-playground/apps/accounts"
	"go-playground/core"
	"net/http"
)

//func home(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//}

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

	//// Add db to request context
	//r.Use(core.DatabaseMiddleware(db))

	// Setup jwt
	r.Use(core.JwtAuthMiddleware(db))

	//defer core.ErrorHandler(db.Close())

	http.ListenAndServe(":5000", r)
	defer core.ErrorHandler(core.CloseDatabase())
}
