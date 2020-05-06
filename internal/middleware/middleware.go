package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Headers(allowOrigin string) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("X-XSS-Protection", "1; mode=block")
			w.Header().Set("X-Frame-Options", "DENY")
			w.Header().Set("Strict-Transport-Security", "max-age=604800; includeSubDomains; preload")
			w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, GET, POST, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			handler.ServeHTTP(w, r)
		})
	}
}
