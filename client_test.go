package httpez

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_WithBaseURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/test" {
			t.Errorf("expected path /test, got %s", r.URL.Path)
		}
		fmt.Fprintln(w, "ok")
	}))
	defer server.Close()

	client := NewClient().WithBaseURL(server.URL)
	_, _, err := client.Get("/test").AsBytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestClient_WithHeader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") != "httpez" {
			t.Errorf("expected User-Agent httpez, got %s", r.Header.Get("User-Agent"))
		}
		fmt.Fprintln(w, "ok")
	}))
	defer server.Close()

	client := NewClient().
		WithBaseURL(server.URL).
		WithHeader("User-Agent", "httpez")

	_, _, err := client.Get("/").AsBytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
