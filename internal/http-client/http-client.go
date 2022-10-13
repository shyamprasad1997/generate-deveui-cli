package httpclient

import (
	"net/http"
)

type Client interface {
	Do(*http.Request) (*http.Response, error)
}
type HttpClient struct {
	client http.Client
}

func (c *HttpClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

func NewHttpClient(client http.Client) *HttpClient {
	return &HttpClient{client}
}
