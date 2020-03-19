package accounts

import (
	"errors"
	"go-playground/core"
	"gopkg.in/validator.v2"
	"net/http"
)

func (u *User) validate(w http.ResponseWriter) error {
	if err := validator.Validate(u); err != nil {
		core.JsonErrorHandler400(w, err)
		return err
	}

	// Check if email doesn't exist
	db := core.GetDb()
	status, err := db.Model(u).Where("email = ?", u.Email).Exists()
	if core.JsonErrorHandler500(w, err) {
		return err
	}

	if status {
		core.JsonErrorHandler400(w, errors.New("email already exists"))
		return errors.New("")
	}
	return nil
}

func (u *User) decodeAndValidate(w http.ResponseWriter, r *http.Request) error {
	if err := u.decodeRequestData(w, r); err != nil {
		return err
	}
	if err := u.validate(w); err != nil {
		return err
	}
	return nil
}
