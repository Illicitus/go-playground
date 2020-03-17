package core

import (
	"errors"
	"net/http"
)

func PermissionsCheck(permission string, u interface{}, w http.ResponseWriter, r *http.Request) bool {
	db := GetDb()

	jwtToken := r.Context().Value("userJwtToken")

	switch permission {
	default:
		return JsonUnauthorizedErrorHandler(w, errors.New("unauthorized"))

	case "isAuthenticated":
		if jwtToken != nil {
			userId, err := CheckUserJwtToken(jwtToken.(string))
			if err != nil {
				if JsonInternalServerErrorHandler(w, err) {
					return false
				}
			}
			err = db.Model(u).Where("id = ?", userId).Select()
			if err := JsonInternalServerErrorHandler(w, err); err {
				return false
			}
			return true
		}
		JsonUnauthorizedErrorHandler(w, errors.New("unauthorized"))
		return false
	}
}
