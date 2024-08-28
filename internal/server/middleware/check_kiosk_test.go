package middleware_test

import (
	"lotery_viking/internal"
	"lotery_viking/internal/server/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

const InvalidKioskID = "{\"error\":\"Invalid Kiosk ID\"}"

func TestCheckKiosk(t *testing.T) {
	t.Run("valid kiosk ID", func(t *testing.T) {
		g := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(g)
		c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
		c.Request.Header.Set("Authorization", "Bearer valid-kiosk-id")
		middleware.CheckKiosk()(c)
		internal.AssertStatusCode(t, g.Code, http.StatusOK)
		macKiosk := c.Request.Context().Value("macKiosk").(string)
		internal.AssertContextValue(t, macKiosk, "valid-kiosk-id")
	})

	t.Run("invalid kiosk ID", func(t *testing.T) {
		g := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(g)
		c.Request = httptest.NewRequest(http.MethodGet, "/", nil)
		middleware.CheckKiosk()(c)
		internal.AssertStatusCode(t, g.Code, http.StatusForbidden)
		internal.AssertResponseBody(t, g.Body, InvalidKioskID)
	})
}
