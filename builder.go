package httpez

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// A RequestBuilder is a fluent API for constructing and executing a single HTTP
// request. It allows for method chaining to easily configure request headers,
// query parameters, and the request body.
type RequestBuilder struct {
	client *Client
	req    *http.Request
	err    error
}

// WithHeader adds a header to the request. This header is specific to this single
// request and will not be carried over to other requests made by the client.
func (rb *RequestBuilder) WithHeader(key, value string) *RequestBuilder {
	if rb.err != nil {
		return rb
	}

	rb.req.Header.Add(key, value)
	return rb
}

// WithContext replaces the request's default context with the provided ctx.
// This is used to control cancellation and deadlines for a single HTTP request.
func (rb *RequestBuilder) WithContext(ctx context.Context) *RequestBuilder {
	if rb.err != nil {
		return rb
	}
	rb.req = rb.req.WithContext(ctx)
	return rb
}

// WithQuery adds a URL query parameter to the request.
func (rb *RequestBuilder) WithQuery(key, value string) *RequestBuilder {
	if rb.err != nil {
		return rb
	}

	q := rb.req.URL.Query()
	q.Add(key, value)
	rb.req.URL.RawQuery = q.Encode()
	return rb
}

// WithJSON marshals the given value v into a JSON body and sets the Content-Type header
// to application/json.
func (rb *RequestBuilder) WithJSON(v interface{}) *RequestBuilder {
	if rb.err != nil {
		return rb
	}

	b, err := json.Marshal(v)
	if err != nil {
		rb.err = err
		return rb
	}

	rb.req.Body = io.NopCloser(bytes.NewReader(b))
	rb.req.Header.Set("Content-Type", "application/json")
	return rb
}

// WithForm sets the request body to the URL-encoded form data and sets the Content-Type
// header to application/x-www-form-urlencoded.
func (rb *RequestBuilder) WithForm(data url.Values) *RequestBuilder {
	if rb.err != nil {
		return rb
	}

	rb.req.Body = io.NopCloser(strings.NewReader(data.Encode()))
	rb.req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return rb
}

// Do executes the HTTP request. It applies all configured middlewares before performing
// the RoundTrip.
func (rb *RequestBuilder) Do() (*http.Response, error) {
	if rb.err != nil {
		return nil, rb.err
	}

	transport := rb.client.Client.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	for i := len(rb.client.middlewares) - 1; i >= 0; i-- {
		transport = rb.client.middlewares[i](transport)
	}

	return transport.RoundTrip(rb.req)
}

// AsBytes executes the request and reads the entire response body into a byte slice. It
// returns the body, the http.Response, and any error encountered. The response body is
// automatically closed.
func (rb *RequestBuilder) AsBytes() ([]byte, *http.Response, error) {
	resp, err := rb.Do()
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}
	return body, resp, nil
}

// AsJSON executes the request, reads the response body, and unmarshals the JSON content into
// the provided interface v. The response body is automatically closed.
func (rb *RequestBuilder) AsJSON(v interface{}) (*http.Response, error) {
	body, resp, err := rb.AsBytes()
	if err != nil {
		return resp, err
	}

	if err := json.Unmarshal(body, v); err != nil {
		return resp, err
	}
	return resp, nil
}
