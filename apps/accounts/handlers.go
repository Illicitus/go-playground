package accounts

import (
	"encoding/json"
	"go-playground/core"
	"net/http"
)

type UserProfile struct {
	Id    int64             `json:"id"`
	Name  string            `json:"name"`
	Email string            `json:"email"`
	Token core.UserJwtToken `json:"tokens"`
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {

	// Decode json and create user object
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	core.JsonBadRequestErrorHandler(w, err)

	// Setup user hashed password
	u.setPassword()

	// Get db connection and insert new user
	db := core.GetDb()

	err = db.Insert(&u)
	core.JsonBadRequestErrorHandler(w, err)

	// Return user object as response and add jwt token
	js, err := json.Marshal(UserProfile{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
		Token: core.CreateNewUserJwtToken(u.Email),
	})
	core.JsonInternalServerErrorHandler(w, err)
	core.JsonResponce(w, js, http.StatusCreated)
}

//switch method := r.Method; method {
//case "GET":
//	fmt.Printf("Get")
//case "POST":
//	fmt.Printf("Post")
//default:
//	fmt.Printf("Not allowed method.")
//}
