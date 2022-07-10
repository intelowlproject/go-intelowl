package gointelowl
<<<<<<< HEAD

import (
	"context"
=======
import (
	"context"
	"errors"
>>>>>>> fdde6f08176939476d9ae8c7efed93b8aed0eb51
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type IntelOwlError struct {
	StatusCode int
	Data       []byte
<<<<<<< HEAD
	Response   *http.Response
=======
>>>>>>> fdde6f08176939476d9ae8c7efed93b8aed0eb51
}

// * implementing the error interface
func (intelOwlError *IntelOwlError) Error() string {
	errorMessage := fmt.Sprintf("Status Code: %d \n Error: %s", intelOwlError.StatusCode, string(intelOwlError.Data))
	return errorMessage
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
	options    *IntelOwlClientOptions
	client     *http.Client
	TagService *TagService
}

<<<<<<< HEAD
func NewIntelOwlClient(options *IntelOwlClientOptions, httpClient *http.Client) IntelOwlClient {
=======
func MakeNewIntelOwlClient(options *IntelOwlClientOptions, httpClient *http.Client) IntelOwlClient {
>>>>>>> fdde6f08176939476d9ae8c7efed93b8aed0eb51

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
<<<<<<< HEAD
	client.TagService = &TagService{
		client: &client,
	}
	return client
}

func (client *IntelOwlClient) newRequest(ctx context.Context, request *http.Request) (*successResponse, error) {
=======
	client.TagService = &TagService{client: &client}
	return client
}

func (client *IntelOwlClient) makeRequest(ctx context.Context, request *http.Request) (*successResponse, error) {
>>>>>>> fdde6f08176939476d9ae8c7efed93b8aed0eb51
	request = request.WithContext(ctx)

	request.Header.Set("Content-Type", "application/json")

	tokenString := fmt.Sprintf("token %s", client.options.Token)

	request.Header.Set("Authorization", tokenString)
	response, err := client.client.Do(request)

<<<<<<< HEAD
	// * Checking for context errors such as reaching the deadline and/or Timeout
=======
>>>>>>> fdde6f08176939476d9ae8c7efed93b8aed0eb51
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
<<<<<<< HEAD
			errorMessageBytes := []byte(errorMessage)
			intelOwlError := IntelOwlError{
				StatusCode: statusCode,
				Data:       errorMessageBytes,
				Response:   response,
			}
			return nil, &intelOwlError
=======
			return nil, errors.New(errorMessage)
>>>>>>> fdde6f08176939476d9ae8c7efed93b8aed0eb51
		}
		intelOwlError := IntelOwlError{
			StatusCode: statusCode,
			Data:       msgBytes,
<<<<<<< HEAD
			Response:   response,
=======
>>>>>>> fdde6f08176939476d9ae8c7efed93b8aed0eb51
		}
		return nil, &intelOwlError
	}

	msgBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
<<<<<<< HEAD
		errorMessage := fmt.Sprintf("Could not convert JSON response. Status code: %d", statusCode)
		errorMessageBytes := []byte(errorMessage)
		intelOwlError := IntelOwlError{
			StatusCode: statusCode,
			Data:       errorMessageBytes,
			Response:   response,
		}
		return nil, &intelOwlError
=======
		fmt.Printf("Could not convert JSON response. Status code: %d", statusCode)
>>>>>>> fdde6f08176939476d9ae8c7efed93b8aed0eb51
	}
	sucessResp := successResponse{
		StatusCode: statusCode,
		Data:       msgBytes,
	}
	return &sucessResp, nil
}
