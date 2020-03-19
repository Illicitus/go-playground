package accounts

import (
	"encoding/json"
	"go-playground/core"
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
