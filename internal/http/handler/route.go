package handler

import (
	"net/http"
)

func (h *Handler) Route(secret string) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /register", h.registration)
	router.HandleFunc("POST /login", h.login)
	router.HandleFunc("GET /dummyLogin", h.dummyLogin)

	router.Handle("POST /house/create", h.jwtMiddleware(http.HandlerFunc(h.createHouse), []string{"moderator"}))
	router.Handle("GET /house/{id}", h.jwtMiddleware(http.HandlerFunc(h.getHouse), []string{"client", "moderator"}))

	router.Handle("POST /flat/create", h.jwtMiddleware(http.HandlerFunc(h.createFlat), []string{"client", "moderator"}))
	router.Handle("POST /flat/update", h.jwtMiddleware(http.HandlerFunc(h.updateFlat), []string{"moderator"}))

	return router
}
