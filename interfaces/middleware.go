package interfaces

import (
    "log"
    "net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
        next.ServeHTTP(w, r)
    })
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Implement authentication check here
        // For example, checking a JWT token in the request header
        next.ServeHTTP(w, r)
    })
}
