package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RetrieveOrganizationByUsername returns an organization by its username
// API PATH: /organizations/{username}
// Method: GET
func RetrieveOrganizationByUsername(ctx context.Context, username string) (*OrganizationResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathOrganizations)
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
	//fmt.Fprint(os.Stdout, string(b[:]))

	data := new(OrganizationResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//RetrieveUsersOnOrganization returns list of users on a given organization
// API PATH: /organizations/{username}/users
// Method: GET
func RetrieveUsersOnOrganization(ctx context.Context, username string, queries *CommonQuery) (*UserOnOrganizationResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathOrganizations)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.URL.Path += "/" + username + "/users"
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

	data := new(UserOnOrganizationResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//RetrieveListingOnOrganization returns list of listings on a given organization
// API PATH: /organizations/{username}/listings
// Method: GET
func RetrieveListingOnOrganization(ctx context.Context, username string, queries *OrganizationListingQuery) (*ListingResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathOrganizations)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.URL.Path += "/" + username + "/listings"
	addOrganizationListingQuery(req, queries)

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

	data := new(ListingResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//RetrieveArticlesOnOrganization returns list of articles belonging a given organization
// API PATH: /organizations/{username}/articles
// Method: GET
func RetrieveArticlesOnOrganization(ctx context.Context, username string, queries *CommonQuery) (*GetArticlesResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathOrganizations)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.URL.Path += "/" + username + "/articles"
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
