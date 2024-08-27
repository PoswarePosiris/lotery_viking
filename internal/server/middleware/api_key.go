package middleware

import (
	"net/http"
	"os"
)

// CheckAPIKey checks if the request contains a valid API key
func CheckAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("api-key")
		expectedAPIKey := os.Getenv("API_KEY")

		if apiKey != expectedAPIKey {
			http.Error(w, "Forbidden access", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
