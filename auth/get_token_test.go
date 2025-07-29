package auth

import (
	"context"
	"testing"

	"github.com/taadis/zhichi/client"
	"github.com/taadis/zhichi/env"
)

func TestGetToken(t *testing.T) {
	ctx := context.Background()

	req := &GetTokenRequest{}
	req.AppId = ""
	req.CreateTime = 0
	req.Sign = ""

	appid := "<your_appid>"
	client := NewAuth(client.NewHttpClient(env.GetHost(), appid))
	rsp, err := client.GetToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp.MarshalIndent())
}
