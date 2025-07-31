package client

import (
	"net/http"
	"time"

	"github.com/taadis/zhichi/cache"
)

type Options struct {
	server string
	// 第三方用户接口调用唯一凭证id
	appid string
	//
	app_key string
	// 时间戳（秒）
	createTime int64
	// 签名
	sign string
	// token
	token string
	// cache
	cache cache.Cache

	// transport
	transport *http.Transport
	// timeout
	timeout *time.Duration
}

type Option func(*Options)

func WithAppid(appid string) Option {
	return func(o *Options) {
		o.appid = appid
	}
}

func WithAppKey(app_key string) Option {
	return func(o *Options) {
		o.app_key = app_key
	}
}

// see more: https://developer.zhichi.com/pages/6cc489/
const (
	// 阿里云服务
	SOBOT_ALIYUN = "www.sobot.com"
	// 腾讯云环境
	SOBOT_TENCENT = "www.soboten.com"
)

func WithServer(server string) Option {
	return func(o *Options) {
		o.server = server
	}
}

func WithTransport(tr *http.Transport) Option {
	return func(o *Options) {
		o.transport = tr
	}
}

func WithTimeout(t *time.Duration) Option {
	return func(o *Options) {
		o.timeout = t
	}
}

func WithCache(cache cache.Cache) Option {
	return func(o *Options) {
		o.cache = cache
	}
}
