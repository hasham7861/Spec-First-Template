package middlewares

import (
	"log"
	"net/http"
	"time"
)

// Simple request logging middleware
func applyLogging(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	}
}