package middleware

import "net/http"

// CORSMiddleware adds CORS headers to allow cross-origin requests
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow these headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Allow these methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// If this is a preflight OPTIONS request, return 200 immediately
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Otherwise, continue to next handler
		next.ServeHTTP(w, r)
	})
}
