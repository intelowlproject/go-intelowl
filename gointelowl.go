package gointelowl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Struct for IntelOwl.
// Initiate it using Token, Instance, Certificate
type IntelOwlClient struct {
	BaseURL     string
	Token       string
	URL         string
	Certificate string
	HTTPClient  *http.Client
}

type Response struct {
	Code    int
	Data    string
	Success bool
}

func CreateClient(Token string, BaseURL string) *IntelOwlClient {
	return &IntelOwlClient{
		Token:       Token,
		BaseURL:     BaseURL,
		Certificate: "",
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *IntelOwlClient) sendRequest(req *http.Request) (Response, error) {
	response := Response{
		Success: false,
	}

	req.Header.Add("Authorization", "Token "+c.Token)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("User-Agent", "IntelOwlClient/3.0.1")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return response, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	sb := string(body)

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusForbidden || res.StatusCode == http.StatusUnauthorized {
			return response, fmt.Errorf("UnAuthorized. Status Code: %d", res.StatusCode)
		} else if res.StatusCode != http.StatusNoContent {
			return response, fmt.Errorf("Unknown error. Status Code: %d", res.StatusCode)
		}
	}

	response = Response{
		Code:    res.StatusCode,
		Data:    sb,
		Success: true,
	}

	return response, nil
}
