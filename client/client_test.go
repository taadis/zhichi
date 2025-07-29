package client

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	server := ""
	appid := ""
	client := NewHttpClient(server, appid)
	if client == nil {
		t.Fatal("failed to create a new client")
	}
}
