package api

import (
	"io/ioutil"
	"os"
)

func GetApiKey() string {
	key := os.Getenv("DEV_API_KEY")
	if len(key) == 0 {
		//ignoring error
		b, _ := ioutil.ReadFile(".devto")
		key = string(b[:])
	}
	return key
}
