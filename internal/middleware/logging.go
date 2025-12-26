package middleware

import (
    "log"
    "net/http"
    "time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Wrap response writer to capture status code
        wrappedWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
        
        log.Printf("[%s] %s %s", r.Method, r.RequestURI, r.RemoteAddr)
        
        next.ServeHTTP(wrappedWriter, r)
        
        duration := time.Since(start)
        log.Printf("[%d] %s %s took %v", wrappedWriter.statusCode, r.Method, r.RequestURI, duration)
    })
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (w *responseWriter) WriteHeader(code int) {
    w.statusCode = code
    w.ResponseWriter.WriteHeader(code)
}