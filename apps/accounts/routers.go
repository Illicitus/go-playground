package accounts

import (
	"github.com/gorilla/mux"
)

func AddAppRouter(router *mux.Router) {
	router.HandleFunc("/sign-up/", signUpHandler).Methods("POST")
	router.HandleFunc("/sign-in/", signInHandler).Methods("POST")
	router.HandleFunc("/profile/", profileHandler).Methods("GET", "PUT", "DELETE")
}
