package internal

import (
	"bytes"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func AssertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected status %v; got %v", want, got)
	}
}

func AssertResponseBody(t *testing.T, body *bytes.Buffer, expected string) {
	t.Helper()

	b, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(b) {
		t.Errorf("expected response body to be %v; got %v", expected, string(b))
	}
}

// assertContentType checks if the response content-type matches the expected value
func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	contentType := response.Header().Get("content-type")
	if !strings.HasPrefix(contentType, want) {
		t.Errorf("response did not have content-type of %s, got %s", want, contentType)
	}
}

func AssertContextValue(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("expected context value %v; got %v", want, got)
	}
}
