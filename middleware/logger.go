package middleware

import (
	"log"
	"net/http"
)

// loggingResponseWriter struct
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewLoggingResponseWriter returns responseWriter with statusCode
func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

// WriteHeader writes response status code
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
}

// Logger logs requests and responses
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		au := r.Header.Get("Authorization")
		log.Println("-->", "header:", au, "URL:", r.URL, "method:", r.Method)
		lrw := newLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)
		statusCode := lrw.statusCode
		log.Println("<--", statusCode, http.StatusText(statusCode))
	})
}

//
