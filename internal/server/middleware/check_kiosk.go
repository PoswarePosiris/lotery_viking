package middleware

import (
	"context"
	"net/http"
	"strings"
)

func CheckKiosk(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		macKiosk := r.Header.Get("Authorization")
		if macKiosk != "" {
			macKiosk = strings.TrimPrefix(macKiosk, "Bearer ")
			ctx := context.WithValue(r.Context(), "macKiosk", macKiosk)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid Kiosk ID", http.StatusForbidden)
		}
	})
}
