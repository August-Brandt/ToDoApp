package todoServer

import (
	"log"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func ApplyMiddlewareChain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := 0; i < len(middlewares); i++ {
		handler = middlewares[i](handler)
	}
	return handler
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %s | URL: %s\n", r.Method, r.URL)
		next(w, r)
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}