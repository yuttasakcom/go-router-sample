package middleware

import (
	"log"
	"net/http"
)

// Auth Middleware
func Auth(token string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			log.Println("Bearer " + token)

			next.ServeHTTP(w, r)
		})
	}
}
