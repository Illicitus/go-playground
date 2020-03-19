package accounts

import (
	"encoding/json"
	"errors"
	"go-playground/core"
	"net/http"
)

type UserProfileWithTokenSchema struct {
	Id    int64             `json:"id"`
	Name  string            `json:"name"`
	Email string            `json:"email"`
	Token core.UserJwtToken `json:"tokens"`
}

func serializeUserProfileWithTokenSchema(u User) ([]byte, error) {
	return json.Marshal(UserProfileWithTokenSchema{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
		Token: core.CreateUserNewJwtToken(u.Id),
	})
}

func (u *User) decodeRequestData(w http.ResponseWriter, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(u)
	if core.JsonErrorHandler400(w, err) {
		return err
	}
	return nil
}

func (u *User) validate(w http.ResponseWriter) error {
	// Validate email
	match, err := core.IsEmailValid(u.Email)
	if err != nil {
		core.JsonErrorHandler400(w, err)
		return err
	}
	if !match {
		core.JsonErrorHandler400(w, errors.New("invalid email"))
		return err
	}

	// Password should be greater then 8 elements and with 1 digit
	match, err = core.IsPasswordValid(u.Password)
	if err != nil {
		core.JsonErrorHandler400(w, err)
		return err
	}
	if !match {
		core.JsonErrorHandler400(w, errors.New("invalid password"))
		return errors.New("")
	}

	// Check if email doesn't exist
	db := core.GetDb()
	status, err := db.Model(&User{}).Where("email = ?", u.Email).Exists()
	if core.JsonErrorHandler500(w, err) {
		return err
	}

	if !status {
		core.JsonErrorHandler400(w, errors.New("email already exists"))
		return errors.New("")
	}
	return nil
}

func (u *User) decodeAndValidate(w http.ResponseWriter, r *http.Request) error {
	err := u.decodeRequestData(w, r)
	if err != nil {
		return err
	}
	err = u.validate(w)
	if err != nil {
		return err
	}
	return nil
}

type UserProfileSchema struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func serializeUserProfileSchema(u User) ([]byte, error) {
	return json.Marshal(UserProfileSchema{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
	})
}
