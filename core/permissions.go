package core

import (
	"errors"
	"go-playground/apps/accounts"
	"net/http"
)

func PermissionsCheck(permission string, u *accounts.User, w http.ResponseWriter, r *http.Request) bool {
	db := GetDb()

	jwtToken := r.Context().Value("userJwtToken")

	switch permission {
	default:
		return JsonUnauthorizedErrorHandler(w, errors.New("unauthorized"))

	case "isAuthenticated":
		if jwtToken != nil {
			userEmail, err := CheckUserJwtToken(jwtToken.(string))
			if err != nil {
				if JsonInternalServerErrorHandler(w, err) {
					return false
				}
			}
			err = db.Model(u).Where("email = ?", userEmail).Select()
			if err := JsonInternalServerErrorHandler(w, err); err {
				return false
			}
			return true
		}
		return false
	}
}
