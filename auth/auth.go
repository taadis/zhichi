package auth

import "github.com/taadis/zhichi/client"

type Auth struct {
	httpClient *client.HttpClient
}

func NewAuth(httpClient *client.HttpClient) *Auth {
	return newAuth(httpClient)
}

func newAuth(httpClient *client.HttpClient) *Auth {
	t := new(Auth)
	t.httpClient = httpClient
	return t
}
