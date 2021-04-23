package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var path = "/articles"

//RetrieveArticles returns the articles of a given username
// API PATH: /articles
// Method: GET
func RetrieveArticles(username string) (*GetArticlesResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//setting value of api-key header
	if len(username) > 0 {
		q := req.URL.Query()
		q.Add("user", username)
		req.URL.RawQuery = q.Encode()
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

	data := &GetArticlesResponse{}
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//UpdateArticle update an specified article
// API PATH: /articles/{id}
// Method: PUT
func UpdateArticle(id string, article *ArticleEdit) (*UpdateArticleResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, path)

	//preparing payload
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(article)
	req, err := http.NewRequest("PUT", url, payloadBuf)
	if err != nil {
		return nil, err
	}

	//setting value of api-key header
	req.Header.Set("api-key", GetApiKey())
	req.Header.Set("Content-Type", "application/json")
	req.URL.Path += id

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	data := &UpdateArticleResponse{}
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
