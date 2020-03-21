package books

import "github.com/gorilla/mux"

func AddAppRouter(router *mux.Router) {
	router.HandleFunc("/", listCreateBooksHandler).Methods("GET", "POST")
	router.HandleFunc("/{id:[0-9]+}/", retrieveUpdateDeleteBookHandler).Methods("GET", "PUT", "DELETE")
	router.HandleFunc("/title-image/", createBookTitleImage).Methods("POST")
	router.HandleFunc("/{id:[0-9]+}/comments/", listCreateBookComments).Methods("GET", "POST")
}
