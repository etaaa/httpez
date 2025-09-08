package httpez

import (
	"testing"
)

func TestHeaders_Add_Get(t *testing.T) {
	h := NewHeaders()
	h.Add("Content-Type", "application/json")

	if h.Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", h.Get("Content-Type"))
	}
}

func TestHeaders_Set(t *testing.T) {
	h := NewHeaders()
	h.Add("X-Request-Id", "123")
	h.Set("X-Request-Id", "456")

	if h.Get("X-Request-Id") != "456" {
		t.Errorf("expected X-Request-Id to be 456, got %s", h.Get("X-Request-Id"))
	}
}

func TestHeaders_Del(t *testing.T) {
	h := NewHeaders()
	h.Add("Content-Type", "application/json")
	h.Del("Content-Type")

	if h.Get("Content-Type") != "" {
		t.Errorf("expected Content-Type to be empty, got %s", h.Get("Content-Type"))
	}
}

func TestHeaders_Clear(t *testing.T) {
	h := NewHeaders()
	h.Add("User-Agent", "httpez")
	h.Add("Accept", "application/json")
	h.Clear()

	if h.Get("User-Agent") != "" || h.Get("Accept") != "" {
		t.Errorf("expected all headers to be cleared")
	}
}
