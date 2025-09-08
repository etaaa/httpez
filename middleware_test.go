package httpez

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_WithMiddleware(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Test-Middleware") != "true" {
			t.Errorf("expected header X-Test-Middleware to be true")
		}
		fmt.Fprintln(w, "ok")
	}))
	defer server.Close()

	testMiddleware := func(next http.RoundTripper) http.RoundTripper {
		return RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
			req.Header.Set("X-Test-Middleware", "true")
			return next.RoundTrip(req)
		})
	}

	client := NewClient().
		WithBaseURL(server.URL).
		WithMiddleware(testMiddleware)

	_, _, err := client.Get("/").AsBytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
