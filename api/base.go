package api

import "errors"

var (
	//api base URL
	baseURL            = "https://dev.to/api"
	pathArticle        = "/articles"
	pathComment        = "/comments"
	ErrorApiKeyMissing = errors.New("API_KEY is missing")
	ErrorIDMissing     = errors.New("id is missing")
)
