package books

import (
	"github.com/gorilla/mux"
	"go-playground/apps/accounts"
	"go-playground/core"
	"net/http"
	"strconv"
)

func createBookTitleImage(w http.ResponseWriter, r *http.Request) {
	// Get user object and check permissions
	var user accounts.User
	if !core.PermissionsCheck("isAuthenticated", &user, w, r) {
		return
	}

	// Parse form data
	paths, err := core.ParseValidateAndCopyFile(w, r, []string{"image"})
	if err != nil {
		return
	}

	var bookTitleImage BookTitleImage
	bookTitleImage.Source = paths["image"]
	bookTitleImage.Thumbnail = paths["image"]

	// Insert new book title image
	if err := bookTitleImage.createNewBookTitleImage(w); err != nil {
		return
	}

	// Return book object as response
	js, err := serializeBookTitleImageSchema(bookTitleImage)
	if core.JsonErrorHandler500(w, err) {
		return
	}

	core.JsonResponse201(w, js)
}

func listCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Get user object and check permissions
	var user accounts.User
	if !core.PermissionsCheck("isAuthenticated", &user, w, r) {
		return
	}

	// Get db connection and get list of books if they exist
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
		if err := core.JsonErrorHandler500(w, err); err {
			return
		}

		// Return book objects as response
		js, err := serializeManyBookSchema(books)
		if core.JsonErrorHandler500(w, err) {
			return
		}

		core.JsonResponse200(w, js)

	case "POST":
		// Decode json and get book data
		var book Book
		book.AuthorId = user.Id

		if err := book.decodeAndValidate(w, r); err != nil {
			return
		}

		// Insert new book
		if err := book.createNewBook(w); err != nil {
			return
		}

		// Return book object as response
		if err := book.getBookById(w, book.Id); err != nil {
			return
		}

		js, err := serializeBookSchema(book)
		if core.JsonErrorHandler500(w, err) {
			return
		}

		core.JsonResponse201(w, js)
	}
}

func retrieveUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Get user object and check permissions
	var user accounts.User
	if !core.PermissionsCheck("isAuthenticated", &user, w, r) {
		return
	}

	// Get db connection and get book id and check if it exist
	db := core.GetDb()

	// Get book id
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if core.JsonErrorHandler400(w, err) {
		return
	}

	// Check if exist
	status, err := db.Model(&Book{}).Where("id = ?", id).Exists()
	if core.JsonErrorHandler500(w, err) {
		return
	}

	if !status {
		core.JsonErrorHandler404(w)
		return
	}

	switch method := r.Method; method {
	default: // GET
		var book Book
		if err := book.getBookById(w, id); err != nil {
			return
		}

		// Return user object as response
		js, err := serializeBookSchema(book)
		if core.JsonErrorHandler500(w, err) {
			return
		}

		core.JsonResponse200(w, js)

	case "PUT":
		// Check if book with selected id exist
		status, err := db.Model(&Book{}).Where("id = ?", id).Exists()
		if core.JsonErrorHandler500(w, err) {
			return
		}

		if !status {
			core.JsonErrorHandler404(w)
			return
		}

		// Decode json and get book data
		var data Book
		data.AuthorId = user.Id

		if err := data.decodeAndValidate(w, r); err != nil {
			return
		}

		err = data.updateBook(w, id)
		if err != nil {
			return
		}

		var book Book
		err = book.getBookById(w, id)
		if err != nil {
			return
		}

		// Return book object as response
		js, err := serializeBookSchema(book)
		if core.JsonErrorHandler500(w, err) {
			return
		}

		core.JsonResponse200(w, js)

	case "DELETE":
		// Check if book with selected id exist
		status, err := db.Model(&Book{}).Where("id = ?", id).Exists()
		if core.JsonErrorHandler500(w, err) {
			return
		}

		if !status {
			core.JsonErrorHandler404(w)
			return
		}

		err = db.Delete(&Book{Id: id})
		if err := core.JsonErrorHandler500(w, err); err {
			return
		}

		// Return empty response
		core.JsonResponse204(w)
	}
}
