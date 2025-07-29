package zhichi

import (
	"context"
	"os"
	"testing"

	"github.com/taadis/zhichi/auth"
	"github.com/taadis/zhichi/env"
	"github.com/taadis/zhichi/ticket"
)

var (
	testBaseUrl = ""
	testAppid   = ""
)

func TestMain(m *testing.M) {
	testBaseUrl = env.GetHost()
	testBaseUrl = SOBOT_ALIYUN
	testAppid = env.GetAppid()
	os.Exit(m.Run())
}

func TestAuth(t *testing.T) {
	appid := testAppid
	client := NewClient(testBaseUrl, appid)
	ctx := context.Background()

	t.Run("auth:get_token", func(t *testing.T) {
		// t.Parallel()
		// Add tests for get_token here
		req := &auth.GetTokenRequest{}
		rsp, err := client.Auth.GetToken(ctx, req)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(rsp.MarshalIndent())
	})
}

func TestTicket(t *testing.T) {
	appid := testAppid
	client := NewClient(testBaseUrl, appid)
	ctx := context.Background()

	t.Run("ticket:get_timezone", func(t *testing.T) {
		// t.Parallel()
		// Add tests for get_token here
		req := &ticket.GetTimezoneRequest{}
		rsp, err := client.Ticket.GetTimezone(ctx, req)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(rsp.String())
	})
}

func TestA(t *testing.T) {
	t.Logf("A...")

	t.Run("B...", func(t *testing.T) {
		t.Logf("B...")

		t.Run("C...", func(t *testing.T) {
			t.Logf("C...")
		})
	})
}
