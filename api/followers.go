package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RetrieveFollowers returns my followers
// API PATH: /followers/users
// Method: GET
func RetrieveFollowers(ctx context.Context, query *FollowersQuery) (*FollowersResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathFollowers)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	//setting value of api-key header
	if err := SetApiKeyHeader(req); err != nil {
		return nil, err
	}
	addFollowerQuery(req, query)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = extractError(b)
	if err != nil {
		return nil, err
	}

	data := new(FollowersResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
