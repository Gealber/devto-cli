package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//RetrieveTagsIFollow returns the tags that I follow
// API PATH: /follows/tags
// Method: GET
func RetrieveTagsIFollow() (*FollowTagsResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathFollowsTags)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//setting value of api-key header
	if err := SetApiKeyHeader(req); err != nil {
		return nil, err
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
	fmt.Fprint(os.Stdout, string(b[:]))

	data := new(FollowTagsResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//RetrieveTags returns all available tags
// API PATH: /tags
// Method: GET
func RetrieveTags(queries *CommonQuery) (*TagsResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathTags)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//setting value of api-key header
	if err := SetApiKeyHeader(req); err != nil {
		return nil, err
	}
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
	fmt.Fprint(os.Stdout, string(b[:]))

	data := new(TagsResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
