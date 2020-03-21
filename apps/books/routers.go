package books

import "github.com/gorilla/mux"

func AddAppRouter(router *mux.Router) {
	router.HandleFunc("/", listCreateBooksHandler).Methods("GET", "POST")
	router.HandleFunc("/{id:[0-9]+}/", retrieveUpdateDeleteBooksHandler).Methods("GET", "PUT", "DELETE")
	router.HandleFunc("/title-image/", createBookTitleImage).Methods("POST")
}
