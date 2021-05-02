package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func GetApiKey() string {
	devtoDir := os.Getenv("DEVTO_DIR")
	filePath := path.Join(devtoDir, ".devto")
	key := os.Getenv("DEV_API_KEY")
	if len(key) == 0 {
		//ignoring error
		b, _ := ioutil.ReadFile(filePath)
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
func payloadReq(ctx context.Context, ptr interface{}, method, pathBase, pathToAdd string) ([]byte, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathBase)

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

	req = req.WithContext(ctx)

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
	//fmt.Fprint(os.Stdout, string(b[:]))
	return b, nil
}

func addQueries(req *http.Request, queries *GetArticleQuery) {
	if queries != nil {
		q := req.URL.Query()
		if queries.Page > 0 {
			q.Add("page", fmt.Sprintf("%d", queries.Page))
		}
		if queries.PerPage > 0 {
			q.Add("per_page", fmt.Sprintf("%d", queries.PerPage))
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
		if queries.Top > 0 {
			q.Add("top", fmt.Sprintf("%d", queries.Top))
		}
		if queries.CollectionID > 0 {
			q.Add("collection_id", fmt.Sprintf("%d", queries.CollectionID))
		}
		if len(queries.Username) > 0 {
			q.Add("username", queries.Username)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func addCommonQueries(req *http.Request, queries *CommonQuery) {
	if queries != nil {
		q := req.URL.Query()
		if queries.Page > 0 {
			q.Add("page", fmt.Sprintf("%d", queries.Page))
		}
		if queries.PerPage > 0 {
			q.Add("per_page", fmt.Sprintf("%d", queries.PerPage))
		}
		req.URL.RawQuery = q.Encode()
	}
}

func addCommentQuery(req *http.Request, queries *CommentQuery) {
	if queries != nil {
		q := req.URL.Query()
		if queries.AID > 0 {
			q.Add("a_id", fmt.Sprintf("%d", queries.AID))
		}
		if queries.PID > 0 {
			q.Add("p_id", fmt.Sprintf("%d", queries.PID))
		}
		req.URL.RawQuery = q.Encode()
	}
}

func addFollowerQuery(req *http.Request, queries *FollowersQuery) {
	if queries != nil {
		q := req.URL.Query()
		if queries.Page > 0 {
			q.Add("page", fmt.Sprintf("%d", queries.Page))
		}
		if queries.PerPage > 0 {
			q.Add("per_page", fmt.Sprintf("%d", queries.PerPage))
		}
		if len(queries.Sort) > 0 {
			q.Add("sort", queries.Sort)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func addListingsQuery(req *http.Request, queries *ListingQuery) {
	if queries != nil {
		q := req.URL.Query()
		if queries.Page > 0 {
			q.Add("page", fmt.Sprintf("%d", queries.Page))
		}
		if queries.PerPage > 0 {
			q.Add("per_page", fmt.Sprintf("%d", queries.PerPage))
		}
		if len(queries.Category) > 0 {
			q.Add("category", queries.Category)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func addOrganizationListingQuery(req *http.Request, queries *OrganizationListingQuery) {
	if queries != nil {
		q := req.URL.Query()
		if queries.Page > 0 {
			q.Add("page", fmt.Sprintf("%d", queries.Page))
		}
		if queries.PerPage > 0 {
			q.Add("per_page", fmt.Sprintf("%d", queries.PerPage))
		}
		if len(queries.Category) > 0 {
			q.Add("category", queries.Category)
		}
		req.URL.RawQuery = q.Encode()
	}
}

func addPodcastQuery(req *http.Request, queries *PodcastEpisodesQuery) {
	if queries != nil {
		q := req.URL.Query()
		if queries.Page > 0 {
			q.Add("page", fmt.Sprintf("%d", queries.Page))
		}
		if queries.PerPage > 0 {
			q.Add("per_page", fmt.Sprintf("%d", queries.PerPage))
		}
		if len(queries.Username) > 0 {
			q.Add("username", queries.Username)
		}
		req.URL.RawQuery = q.Encode()
	}
}
