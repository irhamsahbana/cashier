package jwthandler

import (
	"errors"
	"lucy/cashier/bootstrap"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	UserUUID string `json:"user_uuid"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var mySigningKey = []byte(bootstrap.App.Config.GetString("jwt.secret"))
var lifetime = bootstrap.App.Config.GetInt("jwt.lifetime")

func GenerateAllTokens(userId, role string) (accessT, refreshT string, err error) {
	tokenLifetime := time.Duration(lifetime)

	// Create the claims
	claims := &MyCustomClaims{
		UserUUID: userId,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			Issuer:    "Authenticator",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(tokenLifetime * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	// create refresh token
	refreshClaims := &MyCustomClaims{
		UserUUID: userId,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Authenticator",
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().AddDate(1, 0, 0)),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(mySigningKey)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("token is invalid")
	}

	return claims, nil
}
