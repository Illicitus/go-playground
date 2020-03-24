package main

import (
	"github.com/gorilla/mux"
	"go-playground/apps/accounts"
	"go-playground/apps/books"
	"go-playground/core"
	"go-playground/core/settings"
	"go-playground/server"
)

func main() {
	// Init settings
	settings.Init("local")

	// Collect routers
	r := mux.NewRouter()
	accounts.AddAppRouter(r.PathPrefix("/accounts").Subrouter())
	books.AddAppRouter(r.PathPrefix("/books").Subrouter())

	// Db connection
	db := core.DbConnect()
	core.EnableDbQueryLogger()

	// Setup db tables
	apps := []interface{}{
		(*accounts.User)(nil),
		(*books.Book)(nil),
		(*books.BookTitleImage)(nil),
		(*books.BookComment)(nil),
		(*books.BookLikes)(nil),
	}

	core.ErrorHandler(core.CreateSchema(db, apps))

	// Setup jwt
	r.Use(core.JwtAuthMiddleware(db))

	server.Serve(r)
	defer core.ErrorHandler(core.CloseDatabase())
}
