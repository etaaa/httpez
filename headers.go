package httpez

import (
	"net/http"
	"sync"
)

type Headers struct {
	data http.Header
	mu   sync.RWMutex
}

func NewHeaders() *Headers {
	return &Headers{
		data: make(http.Header),
	}
}

func (h *Headers) Get(key string) string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.data.Get(key)
}

func (h *Headers) Add(key, value string) *Headers {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data.Add(key, value)
	return h
}

func (h *Headers) Set(key, value string) *Headers {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data.Set(key, value)
	return h
}

func (h *Headers) Del(key string) *Headers {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data.Del(key)
	return h
}
