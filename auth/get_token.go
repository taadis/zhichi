package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/taadis/zhichi/core"
)

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

func (c *Auth) GetToken(ctx context.Context, req *GetTokenRequest) (*GetTokenResponse, error) {
	return c.getToken(ctx, req)
}

func (c *Auth) getToken(ctx context.Context, req *GetTokenRequest) (*GetTokenResponse, error) {
	// if req.AppId == "" {
	// 	return nil, fmt.Errorf("missing required app_id")
	// }
	// if req.CreateTime == 0 {
	// 	return nil, fmt.Errorf("missing required create_time")
	// }
	// if req.Sign == "" {
	// 	return nil, fmt.Errorf("missing required sign")
	// }
	// %s={task_id}
	// apiUrl := fmt.Sprintf("/completions-messages/%s/stop", req.TaskId)
	endpoint := "/api/get_token"
	r, err := c.httpClient.NewRawRequest(ctx, http.MethodGet, endpoint, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create base request: %w", err)
	}

	rsp, err := c.httpClient.SendRequest(r)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %s: %s", rsp.Status, c.httpClient.ReadResponseBody(rsp.Body))
	}

	var ret GetTokenResponse
	if err := json.NewDecoder(rsp.Body).Decode(&ret); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &ret, nil
}
