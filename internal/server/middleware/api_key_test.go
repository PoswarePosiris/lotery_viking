package middleware_test

import (
	"lotery_viking/internal"
	"lotery_viking/internal/server/middleware"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

const ForbiddenAccess = "{\"error\":\"Forbidden access\"}"

func TestCheckAPIKey(t *testing.T) {
	t.Run("valid API key", func(t *testing.T) {
		os.Setenv("API_KEY", "valid-api-key")
		g := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(g)
		c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
		c.Request.Header.Set("api-key", "valid-api-key")
		middleware.CheckAPIKey()(c)
		internal.AssertStatusCode(t, g.Code, http.StatusOK)
	})

	t.Run("invalid API key", func(t *testing.T) {
		os.Setenv("API_KEY", "valid-api-key")
		g := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(g)
		c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
		c.Request.Header.Set("api-key", "invalid-api-key")
		middleware.CheckAPIKey()(c)
		internal.AssertStatusCode(t, g.Code, http.StatusForbidden)
		internal.AssertResponseBody(t, g.Body, ForbiddenAccess)
	})

	t.Run("missing API key", func(t *testing.T) {
		os.Setenv("API_KEY", "valid-api-key")
		g := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(g)
		c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
		middleware.CheckAPIKey()(c)
		internal.AssertStatusCode(t, g.Code, http.StatusForbidden)
		internal.AssertResponseBody(t, g.Body, ForbiddenAccess)
	})
}
