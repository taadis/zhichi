package ticket

import "github.com/taadis/zhichi/client"

type Ticket struct {
	httpClient *client.HttpClient
}

func NewTicket(httpClient *client.HttpClient) *Ticket {
	return newTicket(httpClient)
}

func newTicket(httpClient *client.HttpClient) *Ticket {
	t := new(Ticket)
	t.httpClient = httpClient
	return t
}
