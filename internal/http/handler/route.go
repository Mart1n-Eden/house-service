package handler

import (
	"net/http"
)

func (h *Handler) Route() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /register", h.Registration)
	router.HandleFunc("POST /login", h.Login)
	router.HandleFunc("GET /dummyLogin", h.DummyLogin)

	router.Handle("POST /house/{id}/subscribe", h.jwtMiddleware(http.HandlerFunc(h.NewSubscription), []string{"client", "moderator"}))
	router.Handle("POST /house/create", h.jwtMiddleware(http.HandlerFunc(h.CreateHouse), []string{"moderator"}))
	router.Handle("GET /house/{id}", h.jwtMiddleware(http.HandlerFunc(h.GetHouse), []string{"client", "moderator"}))

	router.Handle("POST /flat/create", h.jwtMiddleware(http.HandlerFunc(h.CreateFlat), []string{"client", "moderator"}))
	router.Handle("POST /flat/update", h.jwtMiddleware(http.HandlerFunc(h.UpdateFlat), []string{"moderator"}))

	return router
}
