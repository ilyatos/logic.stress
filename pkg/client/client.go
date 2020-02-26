package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	BaseUrl    *url.URL
	Bearer     string
	HttpClient *http.Client
}

func NewClient(baseUrl *url.URL, bearer string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		BaseUrl:    baseUrl,
		Bearer:     bearer,
		HttpClient: httpClient,
	}
}

func (c *Client) newRequest(method, path string, body interface{}, query map[string]string) (*http.Request, error) {
	ref := &url.URL{Path: path}
	u := c.BaseUrl.ResolveReference(ref)

	if query != nil {
		q := u.Query()
		for name, value := range query {
			q.Add(name, value)
		}
		u.RawQuery = q.Encode()
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, fmt.Errorf("encoding body error: %w", err)
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, fmt.Errorf("creating request error: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Bearer)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("doing request error: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, err
}
