package handler

import (
	"context"
	"net/http"

	"house-service/internal/http/handler/tools"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) jwtMiddleware(next http.Handler, role []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authHeader)
		if header == "" {
			tools.SendClientError(w, "no auth header", http.StatusUnauthorized)
			return
		}

		id, userRole, err := h.authService.ParseToken(header)
		if err != nil {
			tools.SendClientError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var ok bool
		for i := range role {
			if role[i] == userRole {
				ok = true
			}
		}

		if !ok {
			tools.SendClientError(w, "invalid role", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "role", userRole)
		ctx = context.WithValue(ctx, "id", id)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
