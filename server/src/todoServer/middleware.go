package todoServer

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s | URL: %s\n", r.Method, r.URL)
		next(w, r)
	}
}
