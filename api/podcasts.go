package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RetrievePodcastEpisodes returns the podcast episodes
// API PATH: /podcast_episodes
// Method: GET
func RetrievePodcastEpisodes(queries *PodcastEpisodesQuery) (*PodcastResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathPodcast)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

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
	//fmt.Fprint(os.Stdout, string(b[:]))

	data := new(PodcastResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
