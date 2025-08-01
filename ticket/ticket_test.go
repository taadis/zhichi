package ticket

import (
	"os"
	"testing"

	"github.com/taadis/zhichi/core"
	"github.com/taadis/zhichi/env"
)

var test_appid = ""
var test_app_key = ""
var test_host string

func TestMain(m *testing.M) {
	test_appid = env.GetAppid()
	test_app_key = env.GetAppKey()
	test_host = core.SOBOT_ALIYUN
	os.Exit(m.Run())
}
