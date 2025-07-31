package client

import (
	"context"
	"net/http"
	"testing"

	"github.com/taadis/zhichi/cache"
	"github.com/taadis/zhichi/core"
)

func TestGetToken(t *testing.T) {
	appid := "123"
	app_key := "your_app_key"
	ctx := context.Background()
	at := NewDefaultAccessToken(core.SOBOT_ALIYUN, appid, app_key, http.DefaultClient, cache.NewMemoryCache())
	token, err := at.GetTokenFromServer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("got token=%s", token)
}
