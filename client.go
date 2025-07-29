package zhichi

import (
	"github.com/taadis/zhichi/auth"
	"github.com/taadis/zhichi/client"
	"github.com/taadis/zhichi/ticket"
)

type Client struct {
	// Auth APIs
	Auth *auth.Auth

	// 工单APIs
	Ticket *ticket.Ticket
}

func NewClient(server string, appid string, opts ...client.Option) *Client {
	httpClient := client.NewHttpClient(server, appid, opts...)

	c := new(Client)
	c.Auth = auth.NewAuth(httpClient)
	c.Ticket = ticket.NewTicket(httpClient)
	return c
}
