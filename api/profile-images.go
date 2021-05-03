package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RetrieveProfileImage retrieve a user or organization
//profile image information by its corresponding username
// API PATH: /profile_images
// Method: GET
func RetrieveProfileImage(ctx context.Context, username string) (*ProfileImageResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathProfileImage)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.URL.Path += "/" + username

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

	data := new(ProfileImageResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
