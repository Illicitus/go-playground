package core

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go-playground/core/settings"
	"strings"
	"time"
)

type UserJwtClaims struct {
	UserId             int64 `json:"userId"`
	jwt.StandardClaims `json:"token"`
}

type UserJwtToken struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func CreateUserNewJwtToken(userId int64) UserJwtToken {
	s := settings.GetSettings()

	// Create the JWT access & refresh claims, which includes the user email and expiry time
	accessClaims := &UserJwtClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(time.Duration(s.GetInt("jwt.access_expire_time")) * time.Minute).Unix(),
		},
	}

	refreshClaims := &UserJwtClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{

			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(time.Duration(s.GetInt("jwt.refresh_expire_time")) * time.Minute).Unix(),
		},
	}

	// Declare access and refresh token with the algorithm used for signing, and the claims
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// Create the access & refresh JWT strings
	accessTokenString, err := accessToken.SignedString([]byte(s.GetString("jwt.secret_key")))
	if err != nil {
		panic(err)
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(s.GetString("jwt.secret_key")))
	if err != nil {
		panic(err)
	}

	return UserJwtToken{Access: accessTokenString, Refresh: refreshTokenString}
}

func CleanJwtToken(jwtToken string) string {
	// Remove Bearer key
	jwtOnly := strings.Split(jwtToken, " ")
	if l := len(jwtOnly); l > 1 {
		return jwtOnly[1]
	}
	return jwtOnly[0]
}

func CheckUserJwtToken(jwtToken string) (int64, error) {
	s := settings.GetSettings()

	tkn, err := jwt.ParseWithClaims(CleanJwtToken(jwtToken), &UserJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.GetString("jwt.secret_key")), nil
	})

	// Initialize a new instance of UserJwtClaims
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, err
		}
		return 0, err
	}
	if !tkn.Valid {
		return 0, errors.New("invalid jwt token")
	}

	return tkn.Claims.(*UserJwtClaims).UserId, nil
}
