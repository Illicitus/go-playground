package accounts

import (
	"fmt"
	"go-playground/core"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `pg:",unique",json:"email"`
	Password string `json:"password"`
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
	err := db.Insert(&u)
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
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
