package httpez

import (
	"net/http"
)

type Middleware func(http.RoundTripper) http.RoundTripper

type RoundTripperFunc func(*http.Request) (*http.Response, error)

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
