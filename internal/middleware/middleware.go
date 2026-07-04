package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func Middleware(next http.Handler) http.Handler {
	// TODO: Add logging & throttling functionality here
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func ApiKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isIntrospection(r) {
			next.ServeHTTP(w, r)
			return
		}
		key := r.Header.Get("X-API-Key")
		if key == "" || !isValidKey(key) {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"error":"unauthorized – valid X-API-Key header required"}`, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isValidKey(key string) bool {
	allowed := os.Getenv("EXPERIMENTAL_API_KEY")
	return allowed != "" && key == allowed
}

func isIntrospection(r *http.Request) bool {
	if r.Body == nil {
		return false
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return false
	}
	r.Body = io.NopCloser(bytes.NewReader(body))

	var req struct {
		OperationName string `json:"operationName"`
	}
	if err := json.Unmarshal(body, &req); err != nil {
		return false
	}
	return req.OperationName == "IntrospectionQuery"
}
