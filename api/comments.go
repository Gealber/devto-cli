package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RetrieveComments returns the comments of a given article or podcast
// API PATH: /comments?a_id=<id> or /comments?p_id=<id>
// Method: GET
func RetrieveComments(ctx context.Context, queries *CommentQuery) (*CommentsResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathComment)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	addCommentQuery(req, queries)

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

	data := new(CommentsResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//RetrieveComment returns the comment
// API PATH: /comments/{id}
// Method: GET
func RetrieveComment(ctx context.Context, id string) (*CommentType, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathComment)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	if len(id) == 0 {
		return nil, ErrorIDMissing
	}
	req.URL.Path += "/" + id

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

	data := new(CommentType)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
