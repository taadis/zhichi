package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/taadis/http2curl"
)

type HttpClient struct {
	//
	options *Options
	//
	*http.Client
	//
	tokenHandler AccessTokenHandler
}

func NewHttpClient(server string, appid string, opts ...Option) *HttpClient {
	options := &Options{}
	options.server = server
	options.appid = appid
	for _, opt := range opts {
		opt(options)
	}

	tokenHandler := NewDefaultAccessToken(options.server, appid, options.app_key, http.DefaultClient, options.cache)
	c := &HttpClient{
		options:      options,
		Client:       http.DefaultClient,
		tokenHandler: tokenHandler,
	}
	return c
}

func (c *HttpClient) SendRequest(req *http.Request) (*http.Response, error) {
	return c.sendRequest(req)
}

func (c *HttpClient) sendRequest(req *http.Request) (*http.Response, error) {
	curlcmd, err := http2curl.GetCurlCommand(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(curlcmd.String())

	return c.Do(req)
}

func (c *HttpClient) SendJSONRequest(req *http.Request, res interface{}) error {
	return c.sendJSONRequest(req, res)
}

func (c *HttpClient) sendJSONRequest(req *http.Request, res interface{}) error {
	resp, err := c.sendRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read response error: %v", err)
		}
		// var errBody BaseResponse
		// err = json.NewDecoder(resp.Body).Decode(&errBody)
		// if err != nil {
		// 	return err
		// }
		return fmt.Errorf("HTTP response error: %v - %v", resp.StatusCode, string(bs))
	}

	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return err
	}
	return nil
}

func (c *HttpClient) GetServer() string {
	// var server = strings.TrimSuffix(c.options.server, "/")
	return c.options.server
}

func (c *HttpClient) GetAppid() string {
	return c.options.appid
}

func (c *HttpClient) SetAppid(appid string) *HttpClient {
	c.options.appid = appid
	return c
}

func (c *HttpClient) NewRawRequest(ctx context.Context, method string, endpoint string, body interface{}) (*http.Request, error) {
	var b io.Reader
	if body != nil {
		reqBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		b = bytes.NewBuffer(reqBytes)
	} else {
		b = http.NoBody
	}

	uri := &url.URL{
		Scheme: "https",
		Host:   c.GetServer(),
		Path:   path.Join(endpoint),
	}
	req, err := http.NewRequestWithContext(ctx, method, uri.String(), b)
	if err != nil {
		return nil, err
	}
	fmt.Println("[debug] got appid=", c.GetAppid())

	if c.options.token == "" {
		// 无token选项参数时, 从内置逻辑中自动获取token
		token, err := c.tokenHandler.GetAccessToken(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get access token: %w", err)
		}
		req.Header.Set("token", token)
	} else {
		// 有token选项参数时,使用用户提供的token选项参数(需要用户自己处理token存储和刷新逻辑)
		req.Header.Set("token", c.options.token)
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return req, nil
}

func (c *HttpClient) ReadResponseBody(body io.Reader) string {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return fmt.Sprintf("failed to read response body: %v", err)
	}
	return string(bodyBytes)
}
