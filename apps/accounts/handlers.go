package accounts

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-playground/core"
	"net/http"
)

func signUpHandler(w http.ResponseWriter, r *http.Request) {

	// Decode json and create user object
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if core.JsonBadRequestErrorHandler(w, err) {
		return
	}

	// Setup user hashed password
	u.setPassword()

	// Get db connection and insert new user
	db := core.GetDb()

	err = db.Insert(&u)
	if core.JsonBadRequestErrorHandler(w, err) {
		return
	}

	// Return user object as response and add jwt token
	js, err := json.Marshal(UserProfileWithTokenSchema{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
		Token: core.CreateUserNewJwtToken(u.Email),
	})
	if core.JsonInternalServerErrorHandler(w, err) {
		return
	}

	core.JsonResponce(w, js, http.StatusCreated)
}

func signInHandler(w http.ResponseWriter, r *http.Request) {

	// Decode json and create user object
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if core.JsonBadRequestErrorHandler(w, err) {
		return
	}

	// Get db connection and get new user if it exists
	db := core.GetDb()

	// Check if exists
	status, err := db.Model(&User{}).Where("email = ?", u.Email).Exists()
	if core.JsonInternalServerErrorHandler(w, err) {
		return
	}

	if !status {
		core.JsonBadRequestErrorHandler(w, errors.New("invalid email"))
		return
	}

	// Select exists user object
	var user User
	err = db.Model(&user).Where("email = ?", u.Email).Select()
	if err := core.JsonInternalServerErrorHandler(w, err); err {
		return
	}

	// Check password
	if user.checkPassword(u.Password) != true {
		core.JsonBadRequestErrorHandler(w, errors.New("invalid password"))
		return
	}

	// Return user object as response and add jwt token
	js, err := json.Marshal(UserProfileWithTokenSchema{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: core.CreateUserNewJwtToken(user.Email),
	})
	if core.JsonInternalServerErrorHandler(w, err) {
		return
	}

	core.JsonResponce(w, js, http.StatusOK)

}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if !core.PermissionsCheck("isAuthenticated", &user, w, r) {
		return
	}

	switch method := r.Method; method {
	default: // GET

		// Return user object as response
		js, err := json.Marshal(UserProfile{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		})
		if core.JsonInternalServerErrorHandler(w, err) {
			return
		}

		core.JsonResponce(w, js, http.StatusOK)

	case "POST":
		fmt.Printf("PUT")
	case "DELETE":
		fmt.Printf("DELETE")
	}
}
