package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RetrievePodcastEpisodes returns the podcast episodes
// API PATH: /podcast_episodes
// Method: GET
func RetrievePodcastEpisodes(ctx context.Context, queries *PodcastEpisodesQuery) (*PodcastResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathPodcast)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	addPodcastQuery(req, queries)

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

	data := new(PodcastResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
