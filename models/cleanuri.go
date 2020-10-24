package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CleanUriResponse struct {
	ResultURL string `json:"result_url"`
}

type CleanUriPayload struct {
	URL string `json:"url"`
}

type CleanUriAPI struct {}

func (c *CleanUriAPI) Shorten(url string) (string, error) {
	serializedPayload, err := json.Marshal(&CleanUriPayload{
		URL: url,
	})

	if err != nil {
		return "", err
	}

	resp, err := http.Post("https://cleanuri.com/api/v1/shorten", "application/json", bytes.NewReader(serializedPayload))
	if err != nil {
		resp.Body.Close()
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return "", err
	}

	response := &CleanUriResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		resp.Body.Close()
		return "", err
	}

	return response.ResultURL, err
}