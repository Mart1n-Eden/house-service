package handler

import (
	"net/http"

	"house-service/internal/http/middleware"
)

var roleSet = []map[string]struct{}{ // TODO: do something with it
	{
		"moderator": {},
		"user":      {},
	},
	{
		"moderator": {},
	},
}

func (h *Handler) Route(secret string) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /register", h.registration)
	router.HandleFunc("POST /login", h.login)

	router.Handle("POST /house/create", middleware.JWTMiddleware(http.HandlerFunc(h.createHouse), secret, roleSet[1]))
	router.Handle("GET /house/{id}", middleware.JWTMiddleware(http.HandlerFunc(h.getHouse), secret, roleSet[0]))

	router.Handle("POST /flat/create", middleware.JWTMiddleware(http.HandlerFunc(h.createFlat), secret, roleSet[0]))
	router.Handle("POST /flat/update", middleware.JWTMiddleware(http.HandlerFunc(h.updateFlat), secret, roleSet[1]))

	return router
}
