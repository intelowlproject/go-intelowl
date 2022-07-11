package gointelowl

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type IntelOwlError struct {
	StatusCode int
	Message    string
	Response   *http.Response
}

// * implementing the error interface
func (intelOwlError *IntelOwlError) Error() string {
	errorMessage := fmt.Sprintf("Status Code: %d \n Error: %s", intelOwlError.StatusCode, intelOwlError.Message)
	return errorMessage
}

func newIntelOwlError(statusCode int, message string, response *http.Response) *IntelOwlError {
	return &IntelOwlError{
		StatusCode: statusCode,
		Message:    message,
		Response:   response,
	}
}

type successResponse struct {
	StatusCode int
	Data       []byte
}

type IntelOwlClientOptions struct {
	Url   string
	Token string
	// * so basically your SSL cert: path to the cert file!
	Certificate string
	Timeout     time.Duration
}

type IntelOwlClient struct {
	options         *IntelOwlClientOptions
	client          *http.Client
	AnalyzerService *AnalyzerService
}

func NewIntelOwlClient(options *IntelOwlClientOptions, httpClient *http.Client) IntelOwlClient {

	if options.Timeout == 0 {
		options.Timeout = time.Duration(10) * time.Second
	}

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: options.Timeout,
		}
	}
	client := IntelOwlClient{
		options: options,
		client:  httpClient,
	}
	client.AnalyzerService = &AnalyzerService{
		client: &client,
	}
	return client
}

func (client *IntelOwlClient) newRequest(ctx context.Context, request *http.Request) (*successResponse, error) {
	request = request.WithContext(ctx)

	request.Header.Set("Content-Type", "application/json")

	tokenString := fmt.Sprintf("token %s", client.options.Token)

	request.Header.Set("Authorization", tokenString)
	response, err := client.client.Do(request)

	// * Checking for context errors such as reaching the deadline and/or Timeout
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	defer response.Body.Close()

	statusCode := response.StatusCode
	if statusCode < http.StatusOK || statusCode >= http.StatusBadRequest {
		msgBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			errorMessage := fmt.Sprintf("Could not convert JSON response. Status code: %d", statusCode)
			intelOwlError := newIntelOwlError(statusCode, errorMessage, response)
			return nil, intelOwlError
		}
		errorMessage := string(msgBytes)
		intelOwlError := newIntelOwlError(statusCode, errorMessage, response)
		return nil, intelOwlError
	}

	msgBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errorMessage := fmt.Sprintf("Could not convert JSON response. Status code: %d", statusCode)
		intelOwlError := newIntelOwlError(statusCode, errorMessage, response)
		return nil, intelOwlError
	}
	sucessResp := successResponse{
		StatusCode: statusCode,
		Data:       msgBytes,
	}
	return &sucessResp, nil
}
