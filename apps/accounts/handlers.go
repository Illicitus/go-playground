package accounts

import (
	"encoding/json"
	"errors"
	"go-playground/core"
	"net/http"
)

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	// Decode json and validate it
	var user User

	err := user.decodeAndValidate(w, r)
	if err != nil {
		return
	}

	// Insert new user
	err = user.createNewUser(w)
	if err != nil {
		return
	}

	// Return user object as response and add jwt token
	js, err := serializeUserProfileWithTokenSchema(user)
	if core.JsonErrorHandler500(w, err) {
		return
	}

	core.JsonResponse201(w, js)
}

func signInHandler(w http.ResponseWriter, r *http.Request) {
	// Decode json and create user object
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if core.JsonErrorHandler400(w, err) {
		return
	}

	// Get db connection and get new user if it exist
	db := core.GetDb()

	// Check if exist
	status, err := db.Model(&User{}).Where("email = ?", u.Email).Exists()
	if core.JsonErrorHandler500(w, err) {
		return
	}

	if !status {
		core.JsonErrorHandler400(w, errors.New("invalid email"))
		return
	}

	// Select exists user object
	var user User
	err = db.Model(&user).Where("email = ?", u.Email).Select()
	if err := core.JsonErrorHandler500(w, err); err {
		return
	}

	// Check password
	if user.checkPassword(u.Password) != true {
		core.JsonErrorHandler400(w, errors.New("invalid password"))
		return
	}

	// Return user object as response and add jwt token
	js, err := serializeUserProfileWithTokenSchema(user)
	if core.JsonErrorHandler500(w, err) {
		return
	}

	core.JsonResponse200(w, js)

}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	// Get user object and check permissions
	var user User
	if !core.PermissionsCheck("isAuthenticated", &user, w, r) {
		return
	}

	// Get db connection and get new user if it exist
	db := core.GetDb()

	switch method := r.Method; method {
	default: // GET

		// Return user object as response
		js, err := serializeUserProfileSchema(user)
		if core.JsonErrorHandler500(w, err) {
			return
		}

		core.JsonResponse200(w, js)

	case "PUT":
		// Decode json and get user data
		var data User
		err := json.NewDecoder(r.Body).Decode(&data)
		if core.JsonErrorHandler400(w, err) {
			return
		}

		// Update user name, email object and push changes to db
		user.Email = data.Email
		user.Name = data.Name
		err = db.Update(&user)
		if err := core.JsonErrorHandler500(w, err); err {
			return
		}

		// Return user object as response
		js, err := serializeUserProfileSchema(user)
		if core.JsonErrorHandler500(w, err) {
			return
		}

		core.JsonResponse200(w, js)

	case "DELETE":
		// Delete user object
		err := db.Delete(&user)
		if err := core.JsonErrorHandler500(w, err); err {
			return
		}

		core.JsonResponse204(w)
	}
}
