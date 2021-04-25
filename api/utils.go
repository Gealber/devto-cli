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

func addQueries(req *http.Request, queries *GetArticleQuery) {
	if queries != nil {
		q := req.URL.Query()
		if len(queries.Page) > 0 {
			q.Add("page", queries.Page)
		}
		if len(queries.PerPage) > 0 {
			q.Add("per_page", queries.PerPage)
		}
		if len(queries.Tag) > 0 {
			q.Add("tag", queries.Tag)
		}
		if len(queries.Tags) > 0 {
			q.Add("tags", queries.Tags)
		}
		if len(queries.State) > 0 {
			q.Add("state", queries.State)
		}
		if len(queries.Top) > 0 {
			q.Add("top", queries.Top)
		}
		if len(queries.CollectionID) > 0 {
			q.Add("collection_id", queries.CollectionID)
		}
		if len(queries.Username) > 0 {
			q.Add("username", queries.Username)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func addLatesQueries(req *http.Request, queries *GetLatestArticleQuery) {
	if queries != nil {
		q := req.URL.Query()
		if len(queries.Page) > 0 {
			q.Add("page", queries.Page)
		}
		if len(queries.PerPage) > 0 {
			q.Add("per_page", queries.PerPage)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func addCommentQuery(req *http.Request, queries *CommentQuery) {
	if queries != nil {
		q := req.URL.Query()
		if len(queries.AID) > 0 {
			q.Add("a_id", queries.AID)
		}
		if len(queries.PID) > 0 {
			q.Add("p_id", queries.PID)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func addFollowerQuery(req *http.Request, queries *FollowersQuery) {
	if queries != nil {
		q := req.URL.Query()
		if len(queries.Page) > 0 {
			q.Add("page", queries.Page)
		}
		if len(queries.PerPage) > 0 {
			q.Add("per_page", queries.PerPage)
		}
		if len(queries.Sort) > 0 {
			q.Add("sort", queries.Sort)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func addTagsQuery(req *http.Request, queries *TagsQuery) {
	if queries != nil {
		q := req.URL.Query()
		if len(queries.Page) > 0 {
			q.Add("page", queries.Page)
		}
		if len(queries.PerPage) > 0 {
			q.Add("per_page", queries.PerPage)
		}
		req.URL.RawQuery = q.Encode()
	}
}
