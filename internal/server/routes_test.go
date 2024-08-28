package server

import (
	"fmt"
	"lotery_viking/internal"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler(t *testing.T) {
	// Ensure the Server is correctly initialized
	s := &Server{}
	if s == nil {
		t.Fatal("Server instance is nil")
	}

	tests := []struct {
		name       string
		path       string
		handler    gin.HandlerFunc
		expected   string
		statusCode int
		headers    map[string]string
	}{
		{
			name:       "HelloWorldHandler",
			path:       "/",
			handler:    s.HelloWorldHandler,
			expected:   "{\"message\":\"Hello World\"}",
			statusCode: http.StatusOK,
		},
		{
			name:       "ProtectedRoute",
			path:       "/test",
			handler:    s.HelloWorldHandler,
			expected:   "{\"message\":\"Hello World\"}",
			statusCode: http.StatusOK,
			headers:    map[string]string{"X-Api-Key": "secret"},
		},
		// Add more test cases here...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.New()
			r.GET(tt.path, tt.handler)

			req, err := http.NewRequest(http.MethodGet, tt.path, nil)
			if err != nil {
				t.Fatalf("error creating request. Err: %v", err)
			}

			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}

			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)

			fmt.Println("Response headers:", res.Header())

			// Assertions
			internal.AssertStatusCode(t, res.Code, tt.statusCode)
			internal.AssertResponseBody(t, res.Body, tt.expected)
			internal.AssertContentType(t, res, jsonContentType)
		})
	}
}
