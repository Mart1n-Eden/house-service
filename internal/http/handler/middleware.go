package handler

import (
	"context"
	"net/http"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) jwtMiddleware(next http.Handler, role []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authHeader)
		if header == "" {
			http.Error(w, "no auth header", http.StatusUnauthorized)
			return
		}

		audience, err := h.authService.ParseToken(header)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var ok bool
		for i := range role {
			if role[i] == audience {
				ok = true
			}
		}

		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "role", audience[0])
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
