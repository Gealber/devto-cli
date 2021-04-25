package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//RetrieveListings returns the published listings
// API PATH: /listings
// Method: GET
func RetrieveListings(queries *ListingQuery) (*ListingResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathListings)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	addListingsQuery(req, queries)

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

//CreateListing create a new listing
// API PATH: /listings
// Method: POST
func CreateListing(listing *ListingCreate) (*ListingType, error) {
	b, err := payloadReq(listing, "POST", pathListings, "")
	if err != nil {
		return nil, err
	}

	data := &ListingType{}
	if len(b) > 0 {
		err = json.Unmarshal(b, data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

//UpdateListing update a listing
// API PATH: /listings/{id}
// Method: PUT
func UpdateListing(id string, listing *ListingUpdate) (*ModifiedArticle, error) {
	b, err := payloadReq(listing, "PUT", pathListings, "/"+id)
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

//RetrieveListingByID returns a listing by id
// API PATH: /listings/{id}
// Method: GET
func RetrieveListingsByID(id string) (*ListingResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathListings)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
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
	fmt.Fprint(os.Stdout, string(b[:]))

	data := new(ListingResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
