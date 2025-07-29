package ticket

import (
	"context"
	"testing"

	"github.com/taadis/zhichi/client"
	"github.com/taadis/zhichi/env"
)

func TestGetTimezone(t *testing.T) {
	ctx := context.Background()

	req := &GetTimezoneRequest{}
	host := env.GetHost()
	appid := env.GetAppid()
	client := client.NewHttpClient(host, appid)
	ticket := NewTicket(client)
	rsp, err := ticket.GetTimezone(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp.String())
}
