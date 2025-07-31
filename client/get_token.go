package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/taadis/http2curl"
	"github.com/taadis/zhichi/cache"
	"github.com/taadis/zhichi/core"
	"github.com/taadis/zhichi/sign"
)

// AccessTokenHandler AccessToken 接口
type AccessTokenHandler interface {
	GetAccessToken(ctx context.Context) (accessToken string, err error)
}

var _ AccessTokenHandler = (*DefaultAccessToken)(nil)

type DefaultAccessToken struct {
	host string
	//
	appid string
	//
	app_key string
	//
	cache cache.Cache
	//
	sync.Mutex
	//
	httpClient *http.Client
}

func NewDefaultAccessToken(host string, appid string, app_key string, httpClient *http.Client, cache cache.Cache) *DefaultAccessToken {
	t := new(DefaultAccessToken)
	t.host = host
	t.appid = appid
	t.app_key = app_key
	t.httpClient = httpClient
	t.cache = cache
	return t
}

func (t *DefaultAccessToken) GetAccessToken(ctx context.Context) (string, error) {
	if t.cache == nil {
		return "", fmt.Errorf("[zhichi]invalid cache is nil")
	}

	// 先从cache中取
	key := fmt.Sprintf("zhichi:token:%s", t.appid)

	val := t.cache.Get(ctx, key)
	if val != nil {
		return val.(string), nil
	}

	// 加锁,防止并发获取token,缓存失效问题,从服务器获取到不同的token
	t.Lock()
	defer t.Unlock()

	// 双检,防止并发请求时,锁释放后,其他并发重复请求服务器
	val = t.cache.Get(ctx, key)
	if val != nil {
		return val.(string), nil
	}

	// 无缓存时,请求服务器获取
	newToken, err := t.GetTokenFromServer(ctx)
	if err != nil {
		return "", err
	}

	// todo:
	// 默认24小,但这里先改为2分钟,以便快速验证整体完整性
	// 后续需要调整为可通过选项参数控制
	expire := time.Now().Add(2 * time.Minute).Second()
	err = t.cache.Set(ctx, key, newToken, time.Duration(expire))
	if err != nil {
		return newToken, err
	}

	return newToken, nil
}

// see more: https://developer.zhichi.com/pages/950d89/
type GetTokenRequest struct {
	// 第三方用户接口调用唯一凭证id
	AppId string `json:"-" query:"appid"`
	// 时间戳（秒），例如：2019-09-25 15:49:33 的时间戳1569397773
	CreateTime int64 `json:"-" query:"create_time"`
	// 签名md5(appid+create_time+app_key) sign签名,app_key为密钥
	Sign string `json:"-" query:"sign"`
}

type GetTokenResponse struct {
	*core.BaseResponse
	// 返回数据
	Item *TokenItem `json:"item,omitempty"`
}

type TokenItem struct {
	// 访问令牌
	Token string `json:"token"`
	// 访问令牌有效期，单位秒
	Expire int `json:"expire"`
}

func (r *GetTokenResponse) String() string {
	if r == nil {
		return ""
	}
	bs, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(bs)
}

func (r *GetTokenResponse) MarshalIndent() string {
	if r == nil {
		return ""
	}
	bs, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return ""
	}
	return string(bs)
}

func (t *DefaultAccessToken) GetTokenFromServer(ctx context.Context) (string, error) {
	uri := core.GetUri(t.host, "/api/get_token")
	fmt.Printf("uri=%s\n", uri.String())
	req, err := http.NewRequest(http.MethodGet, uri.String(), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create base request: %w", err)
	}
	create_time := fmt.Sprintf("%d", time.Now().Unix())
	sign := sign.GenSign(t.appid, create_time, t.app_key)

	query := req.URL.Query()
	if t.appid != "" {
		query.Set("appid", t.appid)
	}
	if create_time != "" {
		query.Set("create_time", create_time)
	}
	if sign != "" {
		query.Set("sign", sign)
	}
	req.URL.RawQuery = query.Encode()

	// todo:这里的http请求应该和其他接口的复用,http应该共享一个client
	curlcmd, err := http2curl.GetCurlCommand(req)
	if err != nil {
		return "", err
	}
	fmt.Println(curlcmd.String())

	rsp, err := t.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	buf, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	fmt.Printf("[debug]response body statusCode=%d,body=%s\n", rsp.StatusCode, string(buf))

	// if rsp.StatusCode != http.StatusOK && rsp.StatusCode != http.StatusUnauthorized {
	// 	return "", fmt.Errorf("[zhichi]%d - %s", rsp.StatusCode, http.StatusText(rsp.StatusCode))
	// }

	var ret GetTokenResponse
	if err := json.Unmarshal(buf, &ret); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if ret.RetCode != "0" {
		return "", fmt.Errorf("%s - %s", ret.RetCode, ret.RetMsg)
	}

	return ret.Item.Token, nil
}
