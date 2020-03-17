package books

import (
	"encoding/json"
	"go-playground/apps/accounts"
	"go-playground/core"
	"net/http"
)

func listCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Get user object and check permissions
	var user accounts.User
	if !core.PermissionsCheck("isAuthenticated", &user, w, r) {
		return
	}

	// Get db connection and get list of books if they exists
	db := core.GetDb()

	switch method := r.Method; method {
	default: // GET
		var books []Book
		err := db.
			Model(&books).
			Where("author_id = ?", user.Id).
			ColumnExpr("book.*").
			ColumnExpr("u.id AS author__id, u.name AS author__name, u.email AS author__email").
			Join("JOIN users AS u ON u.id = book.author_id").
			Order("id ASC").
			Select()
		if err := core.JsonInternalServerErrorHandler(w, err); err {
			return
		}

		// Return user object as response
		js, err := serializeManyBookSchema(books)
		if core.JsonInternalServerErrorHandler(w, err) {
			return
		}

		core.JsonResponce(w, js, http.StatusOK)

	case "POST":
		// Decode json and get user data
		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if core.JsonBadRequestErrorHandler(w, err) {
			return
		}

		// Add author info
		book.Author = user
		book.AuthorId = user.Id

		// Insert new user
		err = db.Insert(&book)
		if core.JsonBadRequestErrorHandler(w, err) {
			return
		}

		// Return user object as response and add jwt token
		js, err := serializeBookSchema(book)
		if core.JsonInternalServerErrorHandler(w, err) {
			return
		}

		core.JsonResponce(w, js, http.StatusCreated)
	}
}
