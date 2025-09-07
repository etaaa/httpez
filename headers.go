package httpez

import (
	"net/http"
	"sync"
)

// A Headers instance holds a collection of HTTP headers. It is safe for concurrent use.
type Headers struct {
	data http.Header
	mu   sync.RWMutex
}

// NewHeaders creates a new Headers instance.
func NewHeaders() *Headers {
	return &Headers{
		data: make(http.Header),
	}
}

// Get returns the value of the first header for the given key. It is safe for concurrent
// reads.
func (h *Headers) Get(key string) string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.data.Get(key)
}

// Add adds a new value for a header key. It is safe for concurrent writes.
func (h *Headers) Add(key, value string) *Headers {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data.Add(key, value)
	return h
}

// Set replaces all existing values for a header key with a new value. It is safe for
// concurrent writes.
func (h *Headers) Set(key, value string) *Headers {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data.Set(key, value)
	return h
}

// Del deletes all values for a given header key. It is safe for concurrent writes.
func (h *Headers) Del(key string) *Headers {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data.Del(key)
	return h
}

// Clear removes all headers from the collection. It is safe for concurrent writes.
func (h *Headers) Clear() *Headers {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data = make(http.Header)
	return h
}
