package token

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	secret string
}

func New(secret string) *JWTToken {
	return &JWTToken{
		secret: secret,
	}
}

func (t *JWTToken) CreateToken(role string) (string, error) {
	claims := jwt.MapClaims{
		"aud": role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.secret))
}

func (t *JWTToken) ParseToken(header string) (string, error) {
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid token format")
	}

	token, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("unexpected signing method")
		}
		return []byte(t.secret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	audience, err := token.Claims.GetAudience()
	if err != nil || len(audience) != 1 {
		return "", errors.New("invalid token claims")
	}

	return audience[0], nil
}
