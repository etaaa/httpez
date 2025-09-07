package httpez

import (
	"io"
	"net/http"
)

type Client struct {
	*http.Client
	headers     *Headers
	middlewares []Middleware
}

func NewClient() *Client {
	c := &Client{
		Client:  &http.Client{},
		headers: NewHeaders(),
	}

	c.Use(c.headersMiddleware())

	return c
}

func (c *Client) Headers() *Headers {
	return c.headers
}

func (c *Client) Use(m Middleware) *Client {
	c.middlewares = append(c.middlewares, m)
	return c
}

func (c *Client) Request(method, url string, body io.Reader) *RequestBuilder {
	req, err := http.NewRequest(method, url, body)
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
