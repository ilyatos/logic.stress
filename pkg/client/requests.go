package client

import (
	"net/http"
)

func (c *Client) Index() (*Message, error) {
	req, err := c.newRequest(http.MethodGet, "/", nil)
	if err != nil {
		return nil, err
	}

	var message *Message
	_, err = c.do(req, message)
	return message, err
}