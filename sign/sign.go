package sign

import (
	"crypto/md5"
	"encoding/hex"
)

func GenSign(appid string, create_time string, app_key string) string {
	// first concat the strings
	s := appid + create_time + app_key

	// then calculate the md5 hash
	h := md5.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	sign := hex.EncodeToString(bs)
	return sign
}
