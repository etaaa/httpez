package httpez

import (
	"net/http"
)

// A Middleware is a function that wraps an http.RoundTripper and returns a new one. This
// pattern allows for the interception and modification of HTTP requests and responses.
type Middleware func(http.RoundTripper) http.RoundTripper

// RoundTripperFunc is an adapter that allows a function with the correct signature to be
// used as an http.RoundTripper.
type RoundTripperFunc func(*http.Request) (*http.Response, error)

// RoundTrip allows RoundTripperFunc to implement the http.RoundTripper interface. It simply
// calls the wrapped function.
func (f RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func (c *Client) headersMiddleware() Middleware {
	return func(next http.RoundTripper) http.RoundTripper {
		return RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
			c.headers.mu.RLock()
			for k, v := range c.headers.data {
				for _, value := range v {
					req.Header.Add(k, value)
				}
			}
			c.headers.mu.RUnlock()
			return next.RoundTrip(req)
		})
	}
}
