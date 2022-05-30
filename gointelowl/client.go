package gointelowl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type errorResponse struct {
	StatusCode int
	Data       interface{}
}

type successResponse struct {
	Data interface{}
}

type IntelOwlClientOptions struct {
	Url         string
	Token       string
	Certificate string
}

type IntelOwlClient struct {
	options *IntelOwlClientOptions
	client  *http.Client
	Tag     *Tag
}

func MakeNewIntelOwlClient(options *IntelOwlClientOptions, httpClient *http.Client) IntelOwlClient {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: time.Duration(10) * time.Second}
	}
	client := IntelOwlClient{
		options: options,
		client:  httpClient,
	}
	client.Tag = &Tag{client: &client}
	return client
}

func (client *IntelOwlClient) makeRequest(ctx context.Context, request *http.Request, typeOfData interface{}) error {
	request = request.WithContext(ctx)

	request.Header.Set("Content-Type", "application/json")

	tokenString := fmt.Sprintf("token %s", client.options.Token)

	request.Header.Set("Authorization", tokenString)
	response, err := client.client.Do(request)

	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		return err
	}

	defer response.Body.Close()

	statusCode := response.StatusCode
	if statusCode < http.StatusOK || statusCode >= http.StatusBadRequest {
		msgBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			errorMessage := fmt.Sprintf("Could not convert JSON response. Status code: %d", statusCode)
			return errors.New(errorMessage)
		}
		errorResp := errorResponse{
			StatusCode: statusCode,
			Data:       string(msgBytes),
		}
		errorJson, err := json.Marshal(errorResp)
		if err != nil {
			errorMessage := fmt.Sprintf("Could not convert error response into JSON. Status code: %d", statusCode)
			return errors.New(errorMessage)
		}
		return errors.New(string(errorJson))
	}

	sucessResp := successResponse{
		Data: typeOfData,
	}
	msgBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Could not convert JSON response. Status code: %d", statusCode)
	}
	if err := json.Unmarshal(msgBytes, &sucessResp.Data); err != nil {
		errorMessage := "could not convert parse JSON"
		return errors.New(errorMessage)
	}
	return nil
}
