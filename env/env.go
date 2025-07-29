package env

import "os"

func GetHost() string {
	return os.Getenv("ZHICHI_HOST")
}

func GetAppid() string {
	return os.Getenv("ZHICHI_APPID")
}
