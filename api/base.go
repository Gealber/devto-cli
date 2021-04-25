package api

import "errors"

var (
	//api base URL
	baseURL           = "https://dev.to/api"
	pathArticle       = "/articles"
	pathComment       = "/comments"
	pathFollowsTags   = "/follows/tags"
	pathTags          = "/tags"
	pathFollowers     = "/followers/users"
	pathListings      = "/listings"
	pathOrganizations = "/organizations"
	pathPodcast       = "/podcast_episodes"
	pathReadingList   = "/readinglist"

	ErrorApiKeyMissing = errors.New("API_KEY is missing")
	ErrorIDMissing     = errors.New("id is missing")
)
