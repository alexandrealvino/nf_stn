package middleware

import (
	//"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hu:=r.Header.Get("username")
		hp:=r.Header.Get("password")
		log.Println("-->","header:",hu,",",hp,"URL:",r.URL,"method:", r.Method)
		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)
		statusCode := lrw.statusCode
		log.Println("<--", statusCode, http.StatusText(statusCode))
		})
}
