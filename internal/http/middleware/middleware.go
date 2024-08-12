package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"house-service/internal/http/model/response"
)

const (
	CookieName = "token"
)

//type Claims struct {
//	Role string `json:"role"`
//}

func JWTMiddleware(next http.Handler, secret string, role map[string]struct{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			response.SendStarus(w, http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			response.SendStarus(w, http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			response.SendStarus(w, http.StatusUnauthorized)
			return
		}

		audience, err := token.Claims.GetAudience()
		if err != nil {
			response.SendStarus(w, http.StatusUnauthorized)
			return
		}

		if _, ok := role[audience[0]]; !ok {
			response.SendStarus(w, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "role", audience[0])
		r = r.WithContext(ctx)

		//if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//	ctx := context.WithValue(r.Context(), "role", claims) // TODO: ???
		//	r = r.WithContext(ctx)
		//} else {
		//	http.Error(w, "invalid token claims", http.StatusUnauthorized)
		//	return
		//}

		next.ServeHTTP(w, r)
	})
}
