package httpez

import (
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	*http.Client
	baseURL     string
	headers     *Headers
	middlewares []Middleware
}

func NewClient() *Client {
	c := &Client{
		Client: &http.Client{},
	}

	c.headers = NewHeaders(c)

	c.WithMiddleware(c.headersMiddleware())

	return c
}

func (c *Client) BaseURL() string {
	return c.baseURL
}

func (c *Client) WithBaseURL(baseURL string) *Client {
	c.baseURL = baseURL
	return c
}

func (c *Client) Headers() *Headers {
	return c.headers
}

func (c *Client) WithHeader(key, value string) *Client {
	c.headers.Add(key, value)
	return c
}

func (c *Client) WithMiddleware(m Middleware) *Client {
	c.middlewares = append(c.middlewares, m)
	return c
}

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

	req, err := http.NewRequest(method, finalURL, body)

	return &RequestBuilder{
		client: c,
		req:    req,
		err:    err,
	}
}

func (c *Client) Get(url string) *RequestBuilder {
	return c.Request("GET", url, nil)
}

func (c *Client) Post(url string, body io.Reader) *RequestBuilder {
	return c.Request("POST", url, body)
}

func (c *Client) Put(url string, body io.Reader) *RequestBuilder {
	return c.Request("PUT", url, body)
}

func (c *Client) Patch(url string, body io.Reader) *RequestBuilder {
	return c.Request("PATCH", url, body)
}

func (c *Client) Delete(url string) *RequestBuilder {
	return c.Request("DELETE", url, nil)
}

func (c *Client) Head(url string) *RequestBuilder {
	return c.Request("HEAD", url, nil)
}
