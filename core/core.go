package core

import "net/url"

// see more: https://developer.zhichi.com/pages/6cc489/
const (
	// 阿里云服务
	SOBOT_ALIYUN = "www.sobot.com"
	// 腾讯云环境
	SOBOT_TENCENT = "www.soboten.com"
)

type BaseResponse struct {
	//
	Errcode string `json:"errcode,omitempty"`
	//
	ErrorDesc string `json:"error_desc,omitempty"`
	// 返回编码
	RetCode string `json:"ret_code,omitempty"`
	// 返回信息
	RetMsg string `json:"ret_msg,omitempty"`
	// 返回数据
	Item interface{} `json:"item,omitempty"`
}

func GetUri(host string, path string) *url.URL {
	uri := &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}
	return uri
}
