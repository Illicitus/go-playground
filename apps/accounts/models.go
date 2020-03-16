package accounts

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"reflect"
	"strings"
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (u *User) str() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Email)
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

func (u *User) GetTableName() string {
	modelName := strings.Split(reflect.TypeOf(u).String(), ".")
	return fmt.Sprintf("%vs", strings.ToLower(modelName[1]))
}
