package accounts

import "go-playground/core"

type UserProfileWithTokenSchema struct {
	Id    int64             `json:"id"`
	Name  string            `json:"name"`
	Email string            `json:"email"`
	Token core.UserJwtToken `json:"tokens"`
}

type UserProfile struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
