package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//RetrieveArticles returns the articles of a given username
// API PATH: /articles
// Method: GET
func RetrieveArticles(username string, queries *GetArticleQuery) (*GetArticlesResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathArticle)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//setting value of api-key header
	if len(username) > 0 {
		q := req.URL.Query()
		q.Add("username", username)
		req.URL.RawQuery = q.Encode()
	}
	addQueries(req, queries)

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

	data := new(GetArticlesResponse)
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
	b, err := payloadReq(article, "PUT", "/"+id)
	if err != nil {
		return nil, err
	}

	data := &UpdateArticleResponse{}
	if len(b) > 0 {
		err = json.Unmarshal(b, data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

//CreateArticle create a new article
// API PATH: /articles
// Method: POST
func CreateArticle(article *ArticleCreate) (*ArticleCreatedResponse, error) {
	b, err := payloadReq(article, "POST", "")
	if err != nil {
		return nil, err
	}

	data := &ArticleCreatedResponse{}
	if len(b) > 0 {
		err = json.Unmarshal(b, data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

//RetrieveLatestArticles returns latest articles
// API PATH: /articles/latest
// Method: GET
func RetrieveLatestArticles(queries *GetLatestArticleQuery) (*GetArticlesResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathArticle+"/"+"latest")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	addLatesQueries(req, queries)

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

	data := new(GetArticlesResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//RetrieveArticleByID returns the article
// API PATH: /articles/{id}
// Method: GET
func RetrieveArticleByID(id string) (*ModifiedArticle, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathArticle)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//setting value of api-key header
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
	fmt.Fprint(os.Stdout, string(b[:]))

	data := new(ModifiedArticle)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
