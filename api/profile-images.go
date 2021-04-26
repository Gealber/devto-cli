package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RetrieveProfileImage retrieve a user or organization
//profile image information by its corresponding username
// API PATH: /profile_images
// Method: GET
func RetrieveProfileImage(username string) (*ProfileImageResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathProfileImage)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

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
	//fmt.Fprint(os.Stdout, string(b[:]))

	data := new(ProfileImageResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
