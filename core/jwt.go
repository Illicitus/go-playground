package core

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type UserJwtClaims struct {
	UserEmail          string `json:"userEmail"`
	jwt.StandardClaims `json:"token"`
}

type UserJwtToken struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func CreateUserNewJwtToken(userEmail string) UserJwtToken {
	// Create the JWT access & refresh claims, which includes the user email and expiry time
	accessClaims := &UserJwtClaims{
		UserEmail: userEmail,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	refreshClaims := &UserJwtClaims{
		UserEmail: userEmail,
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
	jwtOnly := strings.Split(jwtToken, " ")
	if l := len(jwtOnly); l > 1 {
		return jwtOnly[1]
	}
	return jwtOnly[0]
}

func CheckUserJwtToken(jwtToken string) (string, error) {
	tkn, err := jwt.ParseWithClaims(CleanJwtToken(jwtToken), &UserJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJwtSecretKey(), nil
	})

	// Initialize a new instance of UserJwtClaims
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}
	if !tkn.Valid {
		return "", errors.New("invalid jwt token")
	}

	return tkn.Claims.(*UserJwtClaims).UserEmail, nil
}
