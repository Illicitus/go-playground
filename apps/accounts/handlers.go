package accounts

import (
	"encoding/json"
	"errors"
	"go-playground/core"
	"net/http"
)

func signUpHandler(w http.ResponseWriter, r *http.Request) {

	// Decode json and create user object
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if core.JsonBadRequestErrorHandler(w, err) {
		return
	}

	// Setup user hashed password
	user.setPassword()

	// Get db connection and insert new user
	db := core.GetDb()

	err = db.Insert(&user)
	if core.JsonBadRequestErrorHandler(w, err) {
		return
	}

	// Return user object as response and add jwt token
	js, err := serializeUserProfileWithTokenSchema(user)
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
	js, err := serializeUserProfileWithTokenSchema(user)
	if core.JsonInternalServerErrorHandler(w, err) {
		return
	}

	core.JsonResponce(w, js, http.StatusOK)

}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	// Get user object and check permissions
	var user User
	if !core.PermissionsCheck("isAuthenticated", &user, w, r) {
		return
	}

	// Get db connection and get new user if it exists
	db := core.GetDb()

	switch method := r.Method; method {
	default: // GET

		// Return user object as response
		js, err := serializeUserProfileSchema(user)
		if core.JsonInternalServerErrorHandler(w, err) {
			return
		}

		core.JsonResponce(w, js, http.StatusOK)

	case "PUT":
		// Decode json and get user data
		var data User
		err := json.NewDecoder(r.Body).Decode(&data)
		if core.JsonBadRequestErrorHandler(w, err) {
			return
		}

		// Update user name, email object and push changes to db
		user.Email = data.Email
		user.Name = data.Name
		err = db.Update(&user)
		if err := core.JsonInternalServerErrorHandler(w, err); err {
			return
		}

		// Return user object as response
		js, err := serializeUserProfileSchema(user)
		if core.JsonInternalServerErrorHandler(w, err) {
			return
		}

		core.JsonResponce(w, js, http.StatusOK)

	case "DELETE":
		// Delete user object
		err := db.Delete(&user)
		if err := core.JsonInternalServerErrorHandler(w, err); err {
			return
		}

		core.JsonStatusNoContentResponse(w)
	}
}
