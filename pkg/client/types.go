package client

import "net"

type Info struct {
	Message string `json:"message"`
}

type User struct {
	Subdomain string `json:"subdomain"`
	NoVncPass string `json:"novncPass"`
}

type LabStart struct {
	Subdomain  string `json:"subdomain"`
	TemplateId int    `json:"templateId"`
	IP         net.IP `json:"ip"`
}

type LabStatus struct {
	State  string `json:"state"`
	Status int    `json:"status"`
}
