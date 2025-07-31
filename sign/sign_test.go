package sign

import "testing"

func TestGenSign(t *testing.T) {
	appid := "1"
	create_time := "1569397773"
	app_key := "2"
	sign := GenSign(appid, create_time, app_key)

	want := "258eec3118705112b2c53dc8043d4d34"
	if sign != want {
		t.Logf("want %s but got %s", want, sign)
	}
}
