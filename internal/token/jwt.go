package token

import (
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(role string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"aud": role})

	return token.SignedString([]byte(secret))
}
