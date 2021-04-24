package api

import "errors"

var (
	//api base URL
	baseURL            = "https://dev.to/api"
	pathArticle        = "/articles"
	ErrorApiKeyMissing = errors.New("API_KEY is missing")
)
