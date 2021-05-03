package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RetrieveReadingList returns he client to retrieve a
//list of readinglist reactions along with the related
//article for the authenticated user.
// API PATH: /readinglist
// Method: GET
func RetrieveReadingList(ctx context.Context, queries *CommonQuery) (*ReadingListResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathReadingList)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	//setting value of api-key header
	if err := SetApiKeyHeader(req); err != nil {
		return nil, err
	}
	//adding queries
	addCommonQueries(req, queries)

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

	data := new(ReadingListResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
