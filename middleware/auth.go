package middleware

import (
	"log"
	"net/http"
)

// Auth Middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Bearer " + r.Header.Get("Authorization"))

		next.ServeHTTP(w, r)
	})
}
