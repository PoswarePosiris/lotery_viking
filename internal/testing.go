package internal

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func AssertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected status %v; got %v", want, got)
	}
}

func AssertResponseBody(t *testing.T, body io.ReadCloser, expected string) {
	t.Helper()
	defer body.Close()

	b, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(b) {
		t.Errorf("expected response body to be %v; got %v", expected, string(b))
	}
}

// assertContentType checks if the response content-type matches the expected value
func AssertContentType(t testing.TB, response *http.Response, want string) {
	t.Helper()
	if response.Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Header)
	}
}

// newRequestPath creates a new HTTP GET request for the given path
func NewRequestPath(url string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", url), nil)
	return req
}
