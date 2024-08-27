package middleware_test

import (
	"io"
	"lotery_viking/internal"
	"lotery_viking/internal/server/middleware"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCheckAPIKey(t *testing.T) {
	t.Run("valid API key", func(t *testing.T) {
		os.Setenv("API_KEY", "valid-api-key")
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("api-key", "valid-api-key")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		middleware.CheckAPIKey(handler).ServeHTTP(resp, req)
		internal.AssertStatusCode(t, resp.Code, http.StatusOK)
	})

	t.Run("invalid API key", func(t *testing.T) {
		os.Setenv("API_KEY", "valid-api-key")
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("api-key", "invalid-api-key")
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("handler should not be called")
		})
		middleware.CheckAPIKey(handler).ServeHTTP(resp, req)
		internal.AssertStatusCode(t, resp.Code, http.StatusForbidden)
		internal.AssertResponseBody(t, io.NopCloser(resp.Body), "Forbidden access\n")
	})

	t.Run("missing API key", func(t *testing.T) {
		os.Setenv("API_KEY", "valid-api-key")
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("handler should not be called")
		})
		middleware.CheckAPIKey(handler).ServeHTTP(resp, req)
		internal.AssertStatusCode(t, resp.Code, http.StatusForbidden)
		internal.AssertResponseBody(t, io.NopCloser(resp.Body), "Forbidden access\n")
	})
}
