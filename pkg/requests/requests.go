package requests

import (
	"logic-stress/pkg"
	"logic-stress/pkg/responses"
	"net/http"
)

func (c *pkg.Client) Index() (*responses.Message, error) {
	req, err := c.newRequest(http.MethodGet, "/", nil)
	if err != nil {
		return nil, err
	}

	var message *responses.Message
	_, err = c.do(req, message)
	return message, err
}
