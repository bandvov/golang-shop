package interfaces

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next(w, r)
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Your authentication logic here, e.g., checking if the request has a valid token
		// If authentication fails, return a 401 Unauthorized response
		// Otherwise, call the next handler in the chain
		// For simplicity, let's assume authentication succeeds
		next(w, r)
	}
}
