package books

import "github.com/gorilla/mux"

func AddAppRouter(router *mux.Router) {
	router.HandleFunc("/", listCreateBooksHandler).Methods("GET", "POST")
	//router.HandleFunc("/books/{id:[0-9]+}", retrieveUpdateDeleteBooksHandler).Methods("POST")
}
