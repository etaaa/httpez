package httpez

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	*http.Client
	headers *Headers
}

func NewClient() *Client {
	return &Client{
		Client:  &http.Client{},
		headers: NewHeaders(),
	}
}

func (c *Client) Headers() *Headers {
	return c.headers
}

func (c *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	c.headers.mu.RLock()
	for k, v := range c.headers.data {
		for _, value := range v {
			req.Header.Add(k, value)
		}
	}
	c.headers.mu.RUnlock()

	return c.Client.Do(req)
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

func (c *Client) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req)
}

func (c *Client) PostForm(url string, data url.Values) (*http.Response, error) {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func (c *Client) Head(url string) (*http.Response, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
