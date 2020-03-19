package accounts

import (
	"errors"
	"fmt"
	"go-playground/core"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `validate:"nonzero,regexp=^[^@]+@[^\\.]+\\..+$",pg:",unique",json:"email"`
	Password string `validate:"min=8,regexp=(?:.*[0-9])"validate:"nonzero",json:"password"`
}

func (u *User) str() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Email)
}

func (u *User) getTableName() string {
	modelName := strings.Split(reflect.TypeOf(u).String(), ".")
	return fmt.Sprintf("%vs", strings.ToLower(modelName[1]))
}

func (u *User) createNewUser(w http.ResponseWriter) error {
	// Setup user hashed password
	u.setPassword()

	// Create new user object
	db := core.GetDb()
	err := db.Insert(u)
	if core.JsonErrorHandler500(w, err) {
		return err
	}
	return nil
}

func (u *User) setPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}
	u.Password = string(hash)
}

func (u *User) checkPassword(pass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (u *User) getUserByEmail(email string) error {
	db := core.GetDb()

	if err := db.Model(u).Where("email = ?", email).Select(); err != nil {
		return err
	}
	return nil
}

func (u *User) updateUser(w http.ResponseWriter, up User) error {
	db := core.GetDb()

	if u.Email != up.Email {
		status, err := db.Model(&User{}).Where("email = ?", up.Email).Exists()
		if core.JsonErrorHandler500(w, err) {
			return err
		}

		if status {
			core.JsonErrorHandler400(w, errors.New("email already exists"))
			return errors.New("")
		}

		u.Email = up.Email
	}

	if u.Name != up.Name {
		u.Name = up.Name
	}
	if err := validator.Validate(u); err != nil {
		core.JsonErrorHandler400(w, err)
		return err
	}

	if err := db.Update(u); core.JsonErrorHandler500(w, err) {
		return errors.New("")
	}

	return nil
}

func (u *User) deleteUser(w http.ResponseWriter) error {
	db := core.GetDb()

	if err := db.Delete(u); core.JsonErrorHandler500(w, err) {
		return errors.New("")
	}
	return nil
}
