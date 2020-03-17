package core

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
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
	// Create the JWT access & refresh claims, which includes the user email and expiry time
	accessClaims := &UserJwtClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	refreshClaims := &UserJwtClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{

			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
		},
	}

	// Declare access and refresh token with the algorithm used for signing, and the claims
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// Create the access & refresh JWT strings
	accessTokenString, err := accessToken.SignedString(GetJwtSecretKey())
	if err != nil {
		panic(err)
	}

	refreshTokenString, err := refreshToken.SignedString(GetJwtSecretKey())
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
	tkn, err := jwt.ParseWithClaims(CleanJwtToken(jwtToken), &UserJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecretKey(), nil
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
