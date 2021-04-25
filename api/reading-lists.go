package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//RetrieveReadingList returns he client to retrieve a
//list of readinglist reactions along with the related
//article for the authenticated user.
// API PATH: /readinglist
// Method: GET
func RetrieveReadingList(queries *ReadingListQuery) (*ReadingListResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathReadingList)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//setting value of api-key header
	if err := SetApiKeyHeader(req); err != nil {
		return nil, err
	}
	//adding queries
	addReadingListingQuery(req, queries)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	fmt.Fprint(os.Stdout, string(b[:]))

	data := new(ReadingListResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
