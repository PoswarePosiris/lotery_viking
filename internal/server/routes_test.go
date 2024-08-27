package server

import (
	"lotery_viking/internal"
	"net/http"
	"net/http/httptest"
	"testing"
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
		handler    http.HandlerFunc
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
			server := httptest.NewServer(http.HandlerFunc(tt.handler))
			defer server.Close()

			req, err := http.NewRequest(http.MethodGet, server.URL+tt.path, nil)
			if err != nil {
				t.Fatalf("error creating request. Err: %v", err)
			}

			for key, value := range tt.headers {
				req.Header.Set(key, value)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("error making request to server. Err: %v", err)
			}
			defer resp.Body.Close()

			// Assertions
			internal.AssertStatusCode(t, resp.StatusCode, tt.statusCode)
			internal.AssertResponseBody(t, resp.Body, tt.expected)
			internal.AssertContentType(t, resp, jsonContentType)
		})
	}
}
