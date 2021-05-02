package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RetrieveArticles returns the articles of a given username
// API PATH: /articles
// Method: GET
func RetrieveArticles(ctx context.Context, username string, queries *GetArticleQuery) (*GetArticlesResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathArticle)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

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
	//fmt.Fprint(os.Stdout, string(b[:]))

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
func UpdateArticle(ctx context.Context, id string, article *ArticleEdit) (*ModifiedArticle, error) {
	b, err := payloadReq(ctx, article, "PUT", pathArticle, "/"+id)
	if err != nil {
		return nil, err
	}

	data := &ModifiedArticle{}
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
func CreateArticle(ctx context.Context, article *ArticleCreate) (*ModifiedArticle, error) {
	b, err := payloadReq(ctx, article, "POST", pathArticle, "")
	if err != nil {
		return nil, err
	}

	data := &ModifiedArticle{}
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
func RetrieveLatestArticles(ctx context.Context, queries *CommonQuery) (*GetArticlesResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathArticle+"/"+"latest")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

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
	//fmt.Fprint(os.Stdout, string(b[:]))

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
func RetrieveArticleByID(ctx context.Context, id string) (*ModifiedArticle, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathArticle)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

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
	//fmt.Fprint(os.Stdout, string(b[:]))

	data := new(ModifiedArticle)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//RetrieveArticlesVideo returns the articles with videos
// API PATH: /videos
// Method: GET
func RetrieveArticlesVideo(ctx context.Context, id string) (*ArticlesVideoResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, "/videos")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Fprint(os.Stdout, string(b[:]))

	data := new(ArticlesVideoResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//RetrieveMeArticles returns the articles of the authenticated user
// API PATHs:
// * /articles/me
// * /articles/me/published,
// * /articles/me/unpublished
// * /articles/me/all
// Method: GET
func RetrieveMeArticles(ctx context.Context, queries *CommonQuery, pathToAdd string) (*GetArticlesMeResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathArticle)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	//setting value of api-key header
	if err := SetApiKeyHeader(req); err != nil {
		return nil, err
	}
	addCommonQueries(req, queries)
	req.URL.Path += "/me" + pathToAdd

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Fprint(os.Stdout, string(b[:]))

	data := new(GetArticlesMeResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
