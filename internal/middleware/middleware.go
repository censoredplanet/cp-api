package middleware

import (
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	// TODO: Add logging & throttling functionality here
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
