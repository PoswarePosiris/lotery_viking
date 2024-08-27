package middleware_test

import (
	"io"
	"lotery_viking/internal"
	"lotery_viking/internal/server/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckKiosk(t *testing.T) {
	t.Run("valid kiosk ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer valid-kiosk-id")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			macKiosk := r.Context().Value("macKiosk").(string)
			if macKiosk != "valid-kiosk-id" {
				t.Errorf("expected macKiosk to be 'valid-kiosk-id', got '%s'", macKiosk)
			}
		})
		middleware.CheckKiosk(handler).ServeHTTP(resp, req)
		internal.AssertStatusCode(t, resp.Code, http.StatusOK)
	})

	t.Run("invalid kiosk ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("handler should not be called")
		})
		middleware.CheckKiosk(handler).ServeHTTP(resp, req)
		internal.AssertStatusCode(t, resp.Code, http.StatusForbidden)
		internal.AssertResponseBody(t, io.NopCloser(resp.Body), "Invalid Kiosk ID\n")
	})
}
