package ticket

import (
	"context"
	"testing"

	"github.com/taadis/zhichi/cache"
	"github.com/taadis/zhichi/client"
)

func TestSaveUserTicket(t *testing.T) {
	ctx := context.Background()

	req := &SaveUserTicketRequest{}
	client := client.NewHttpClient(
		test_host,
		test_appid,
		client.WithAppKey(test_app_key),
		client.WithCache(cache.NewMemoryCache()),
	)
	ticket := NewTicket(client)
	rsp, err := ticket.SaveUserTicket(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp.String())
}
