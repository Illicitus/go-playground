package core

import (
	"net/http"
)

func PermissionsCheck(permission string, u interface{}, w http.ResponseWriter, r *http.Request) bool {
	db := GetDb()

	jwtToken := r.Context().Value("userJwtToken")

	switch permission {
	default:
		return JsonErrorHandler401(w)

	case "isAuthenticated":
		if jwtToken != nil {
			userId, err := CheckUserJwtToken(jwtToken.(string))
			if err != nil {
				if JsonErrorHandler500(w, err) {
					return false
				}
			}
			err = db.Model(u).Where("id = ?", userId).Select()
			if err := JsonErrorHandler500(w, err); err {
				return false
			}
			return true
		}
		JsonErrorHandler401(w)
		return false
	}
}
