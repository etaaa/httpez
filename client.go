package httpez

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// A Client manages the configuration for an HTTP client, including a base URL,
// default headers, and a chain of middlewares. It wraps a standard http.Client.
type Client struct {
	*http.Client
	baseURL     string
	headers     *Headers
	middlewares []Middleware
}

// NewClient creates and returns a new Client instance with a default http.Client.
// It automatically adds a middleware to apply the client's global headers to all
// requests.
func NewClient() *Client {
	c := &Client{
		Client:  &http.Client{},
		headers: NewHeaders(),
	}

	c.WithMiddleware(c.headersMiddleware())

	return c
}

// BaseURL returns the base URL configured for the client.
func (c *Client) BaseURL() string {
	return c.baseURL
}

// WithBaseURL sets a base URL for all requests made by the client. Relative URLs
// provided to request methods will be resolved against this base URL.
func (c *Client) WithBaseURL(baseURL string) *Client {
	c.baseURL = baseURL
	return c
}

// Headers returns a pointer to the Headers instance, allowing for global header
// configuration on the client.
func (c *Client) Headers() *Headers {
	return c.headers
}

// WithHeader is a convenience method that adds a header to the client's global
// headers. This header will be applied to all subsequent requests.
func (c *Client) WithHeader(key, value string) *Client {
	c.headers.Add(key, value)
	return c
}

// WithMiddleware adds a new Middleware to the client's middleware chain. Middlewares
// are executed in the order they are added.
func (c *Client) WithMiddleware(m Middleware) *Client {
	c.middlewares = append(c.middlewares, m)
	return c
}

// Request creates a new RequestBuilder for a given HTTP method and URL. It resolves
// urlStr against the client's base URL if one is set.
func (c *Client) Request(method, urlStr string, body io.Reader) *RequestBuilder {
	finalURL := urlStr

	if c.baseURL != "" {
		reqURL, err := url.Parse(urlStr)
		if err != nil {
			return &RequestBuilder{client: c, err: err}
		}

		if !reqURL.IsAbs() {
			baseURL, err := url.Parse(c.baseURL)
			if err != nil {
				return &RequestBuilder{client: c, err: err}
			}
			finalURL = baseURL.ResolveReference(reqURL).String()
		}
	}

	req, err := http.NewRequestWithContext(context.Background(), method, finalURL, body)

	return &RequestBuilder{
		client: c,
		req:    req,
		err:    err,
	}
}

// Get creates a new RequestBuilder for a GET request.
func (c *Client) Get(url string) *RequestBuilder {
	return c.Request("GET", url, nil)
}

// Post creates a new RequestBuilder for a POST request with the specified body.
func (c *Client) Post(url string, body io.Reader) *RequestBuilder {
	return c.Request("POST", url, body)
}

// Put creates a new RequestBuilder for a PUT request with the specified body.
func (c *Client) Put(url string, body io.Reader) *RequestBuilder {
	return c.Request("PUT", url, body)
}

// Patch creates a new RequestBuilder for a PATCH request with the specified body.
func (c *Client) Patch(url string, body io.Reader) *RequestBuilder {
	return c.Request("PATCH", url, body)
}

// Delete creates a new RequestBuilder for a DELETE request.
func (c *Client) Delete(url string) *RequestBuilder {
	return c.Request("DELETE", url, nil)
}

// Head creates a new RequestBuilder for a HEAD request.
func (c *Client) Head(url string) *RequestBuilder {
	return c.Request("HEAD", url, nil)
}
