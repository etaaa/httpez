package httpez

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestBuilder struct {
	client *Client
	req    *http.Request
	err    error
}

func (rb *RequestBuilder) WithHeader(key, value string) *RequestBuilder {
	if rb.err != nil {
		return rb
	}

	rb.req.Header.Add(key, value)
	return rb
}

func (rb *RequestBuilder) WithQuery(key, value string) *RequestBuilder {
	if rb.err != nil {
		return rb
	}

	q := rb.req.URL.Query()
	q.Add(key, value)
	rb.req.URL.RawQuery = q.Encode()
	return rb
}

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

func (rb *RequestBuilder) WithForm(data url.Values) *RequestBuilder {
	if rb.err != nil {
		return rb
	}

	rb.req.Body = io.NopCloser(strings.NewReader(data.Encode()))
	rb.req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return rb
}

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
