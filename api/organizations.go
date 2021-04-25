package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//RetrieveOrganizationByUsername returns an organization by its username
// API PATH: /organizations/{username}
// Method: GET
func RetrieveOrganizationByUsername(username string) (*OrganizationResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathOrganizations)
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
	fmt.Fprint(os.Stdout, string(b[:]))

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
func RetrieveUsersOnOrganization(username string, queries *OrganizationQuery) (*UserOnOrganizationResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathOrganizations)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.URL.Path += "/" + username + "/users"
	addOrganizationQuery(req, queries)

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
func RetrieveListingOnOrganization(username string, queries *OrganizationListingQuery) (*ListingResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathOrganizations)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

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
	fmt.Fprint(os.Stdout, string(b[:]))

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
func RetrieveArticlesOnOrganization(username string, queries *OrganizationQuery) (*GetArticlesResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathOrganizations)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.URL.Path += "/" + username + "/articles"
	addOrganizationQuery(req, queries)

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
