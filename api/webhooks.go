package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RetrieveWebhooks list of webhooks they have previously registered.
// API PATH: /webhooks
// Method: GET
func RetrieveWebhooks(ctx context.Context) (*WebhooksResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathWebhooks)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	//setting value of api-key header
	if err := SetApiKeyHeader(req); err != nil {
		return nil, err
	}

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

	data := new(WebhooksResponse)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//CreateWebhook create a new webhook
// API PATH: /webhooks
// Method: POST
func CreateWebhook(ctx context.Context, listing *WebhooksCreateType) (*WebhookCreatedResponse, error) {
	b, err := payloadReq(ctx, listing, "POST", pathWebhooks, "")
	if err != nil {
		return nil, err
	}

	data := &WebhookCreatedResponse{}
	if len(b) > 0 {
		err = json.Unmarshal(b, data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

//RetrieveWebhookByID returns a webhook by id
// API PATH: /webhook/{id}
// Method: GET
func RetrieveWebhookByID(ctx context.Context, id string) (*WebhookTypeBasic, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathWebhooks)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

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
	err = extractError(b)
	if err != nil {
		return nil, err
	}

	data := new(WebhookTypeBasic)
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//DeleteWebhook delete a webhook by id
// API PATH: /webhook/{id}
// Method: DELETE
func DeleteWebhook(ctx context.Context, id string) (*WebhookTypeBasic, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", baseURL, pathWebhooks)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

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
	err = extractError(b)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
