package httpez

import (
	"io"
	"net/http"
	"net/url"
	"strings"
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

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	transport := c.Client.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	for i := len(c.middlewares) - 1; i >= 0; i-- {
		transport = c.middlewares[i](transport)
	}

	return transport.RoundTrip(req)
}

func (c *Client) DoAndReadBody(req *http.Request) ([]byte, *http.Response, error) {
	resp, err := c.Do(req)
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

func (c *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) GetAndReadBody(url string) ([]byte, *http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	return c.DoAndReadBody(req)
}

func (c *Client) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req)
}

func (c *Client) PostAndReadBody(url, contentType string, body io.Reader) ([]byte, *http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.DoAndReadBody(req)
}

func (c *Client) PostForm(url string, data url.Values) (*http.Response, error) {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func (c *Client) PostFormAndReadBody(url string, data url.Values) ([]byte, *http.Response, error) {
	return c.PostAndReadBody(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func (c *Client) Head(url string) (*http.Response, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) HeadAndReadBody(url string) ([]byte, *http.Response, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, nil, err
	}
	return c.DoAndReadBody(req)
}
