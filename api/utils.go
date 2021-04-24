package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func SetApiKeyHeader(req *http.Request) error {
	//setting value of api-key header
	api_key := GetApiKey()
	if len(api_key) == 0 {
		return ErrorApiKeyMissing
	}
	req.Header.Set("api-key", api_key)
	return nil
}

//payloadReq is an util function to perform Post and Put requests
func payloadReq(ptr interface{}, method, pathToAdd string) ([]byte, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathArticle)

	//preparing payload
	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(ptr)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, payloadBuf)
	if err != nil {
		return nil, err
	}

	//setting value of api-key header
	if err := SetApiKeyHeader(req); err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(pathToAdd) > 0 {
		req.URL.Path += pathToAdd
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	//printing on the terminal
	fmt.Fprint(os.Stdout, string(b[:]))
	return b, nil
}
