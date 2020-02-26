package client

import (
	"net/http"
)

func (c *Client) Index() (*Info, error) {
	req, err := c.newRequest(http.MethodGet, "/", nil, nil)
	if err != nil {
		return nil, err
	}

	var message *Info
	_, err = c.do(req, &message)
	return message, err
}

func (c *Client) GetUser(subdomain string) (*User, error) {
	req, err := c.newRequest(http.MethodGet, "/users/"+subdomain, nil, nil)
	if err != nil {
		return nil, err
	}

	user := &User{
		Subdomain: subdomain,
	}
	_, err = c.do(req, &user)
	return user, err
}

func (c *Client) StartLab(startLab *LabStart) error {
	req, err := c.newRequest(http.MethodPost, "/labs/start", startLab, nil)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	return err
}

func (c *Client) StopLab(user *User) error {
	req, err := c.newRequest(http.MethodPost, "/labs/stop", user, nil)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)
	return err
}

func (c *Client) GetLabStatus(user *User) (*LabStatus, error) {
	req, err := c.newRequest(http.MethodGet, "/labs/status", nil, map[string]string{"subdomain": user.Subdomain})
	if err != nil {
		return nil, err
	}

	var labStatus *LabStatus
	_, err = c.do(req, &labStatus)
	return labStatus, err
}
