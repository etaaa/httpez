package httpez

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRequestBuilder_WithQuery(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("drink") != "coffee" {
			t.Errorf("expected query param drink=coffee, got %s", r.URL.RawQuery)
		}
		fmt.Fprintln(w, "ok")
	}))
	defer server.Close()

	client := NewClient().WithBaseURL(server.URL)
	_, _, err := client.Get("/").WithQuery("drink", "coffee").AsBytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRequestBuilder_WithJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		var payload map[string]string
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("failed to decode json: %v", err)
		}
		if payload["message"] != "hello" {
			t.Errorf("expected message 'hello', got %s", payload["message"])
		}

		fmt.Fprintln(w, "ok")
	}))
	defer server.Close()

	client := NewClient().WithBaseURL(server.URL)
	payload := map[string]string{"message": "hello"}
	_, _, err := client.Post("/", nil).WithJSON(payload).AsBytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRequestBuilder_WithForm(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			t.Errorf("expected Content-Type application/x-www-form-urlencoded, got %s", r.Header.Get("Content-Type"))
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read body: %v", err)
		}

		if string(body) != "drink=coffee" {
			t.Errorf("expected form data drink=coffee, got %s", string(body))
		}

		fmt.Fprintln(w, "ok")
	}))
	defer server.Close()

	client := NewClient().WithBaseURL(server.URL)
	formData := url.Values{}
	formData.Set("drink", "coffee")

	_, _, err := client.Post("/", nil).WithForm(formData).AsBytes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRequestBuilder_AsJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"drink": "coffee"}`)
	}))
	defer server.Close()

	client := NewClient().WithBaseURL(server.URL)

	var result struct {
		Drink string `json:"drink"`
	}

	_, err := client.Get("/").AsJSON(&result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Drink != "coffee" {
		t.Errorf("expected name coffee, got %s", result.Drink)
	}
}
