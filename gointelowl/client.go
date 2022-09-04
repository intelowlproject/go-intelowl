// go-intelowl provides an SDK to easily integrate intelowl with your own set of tools.

// go-intelowl makes it easy to automate, configure, and use intelowl with your own set of tools
// with its Idiomatic approach making an analysis is easy as just writing one line of code!
package gointelowl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// IntelOwlError represents an error that has occurred when communicating with IntelOwl.
type IntelOwlError struct {
	StatusCode int
	Message    string
	Response   *http.Response
}

// Error lets you implement the error interface.
// This is used for making custom go errors.
func (intelOwlError *IntelOwlError) Error() string {
	errorMessage := fmt.Sprintf("Status Code: %d \n Error: %s", intelOwlError.StatusCode, intelOwlError.Message)
	return errorMessage
}

// newIntelOwlError lets you easily create new IntelOwlErrors.
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

// IntelOwlClientOptions represents the fields needed to configure and use the IntelOwlClient
type IntelOwlClientOptions struct {
	Url   string `json:"url"`
	Token string `json:"token"`
	// Certificate represents your SSL cert: path to the cert file!
	Certificate string `json:"certificate"`
	// Timeout is in seconds
	Timeout uint64 `json:"timeout"`
}

// IntelOwlClient handles all the communication with your IntelOwl instance.
type IntelOwlClient struct {
	options          *IntelOwlClientOptions
	client           *http.Client
	TagService       *TagService
	JobService       *JobService
	AnalyzerService  *AnalyzerService
	ConnectorService *ConnectorService
	UserService      *UserService
	Logger           *IntelOwlLogger
}

// TLP represents an enum for the TLP attribute used in IntelOwl's REST API.
//
// IntelOwl docs: https://intelowl.readthedocs.io/en/latest/Usage.html#tlp-support
type TLP int

// Values of the TLP enum.
const (
	WHITE TLP = iota + 1
	GREEN
	AMBER
	RED
)

// TLPVALUES represents a map to easily access the TLP values.
var TLPVALUES = map[string]int{
	"WHITE": 1,
	"GREEN": 2,
	"AMBER": 3,
	"RED":   4,
}

// Overriding the String method to get the string representation of the TLP enum
func (tlp TLP) String() string {
	switch tlp {
	case WHITE:
		return "WHITE"
	case GREEN:
		return "GREEN"
	case AMBER:
		return "AMBER"
	case RED:
		return "RED"
	}
	return "WHITE"
}

// ParseTLP is used to easily make a TLP enum
func ParseTLP(s string) TLP {
	s = strings.TrimSpace(s)
	value, ok := TLPVALUES[s]
	if !ok {
		return TLP(0)
	}
	return TLP(value)
}

// Implementing the MarshalJSON interface to make our custom Marshal for the enum
func (tlp TLP) MarshalJSON() ([]byte, error) {
	return json.Marshal(tlp.String())
}

// Implementing the UnmarshalJSON interface to make our custom Unmarshal for the enum
func (tlp *TLP) UnmarshalJSON(data []byte) (err error) {
	var tlpString string
	if err := json.Unmarshal(data, &tlpString); err != nil {
		return err
	}
	if *tlp = ParseTLP(tlpString); err != nil {
		return err
	}
	return nil
}

// NewIntelOwlClient lets you easily create a new IntelOwlClient by providing IntelOwlClientOptions, http.Clients, and LoggerParams.
func NewIntelOwlClient(options *IntelOwlClientOptions, httpClient *http.Client, loggerParams *LoggerParams) IntelOwlClient {

	var timeout time.Duration

	if options.Timeout == 0 {
		timeout = time.Duration(10) * time.Second
	} else {
		timeout = time.Duration(options.Timeout) * time.Second
	}

	// configuring the http.Client
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: timeout,
		}
	}

	// configuring the client
	client := IntelOwlClient{
		options: options,
		client:  httpClient,
	}

	// Adding the services
	client.TagService = &TagService{
		client: &client,
	}
	client.JobService = &JobService{
		client: &client,
	}
	client.AnalyzerService = &AnalyzerService{
		client: &client,
	}
	client.ConnectorService = &ConnectorService{
		client: &client,
	}
	client.UserService = &UserService{
		client: &client,
	}

	// configuring the logger!
	client.Logger = &IntelOwlLogger{}
	client.Logger.Init(loggerParams)

	return client
}

// NewIntelOwlClientThroughJsonFile lets you create a new IntelOwlClient through a JSON file that contains your IntelOwlClientOptions
func NewIntelOwlClientThroughJsonFile(filePath string, httpClient *http.Client, loggerParams *LoggerParams) (*IntelOwlClient, error) {
	optionsBytes, err := os.ReadFile(filePath)
	if err != nil {
		errorMessage := fmt.Sprintf("Could not read %s", filePath)
		intelOwlError := newIntelOwlError(400, errorMessage, nil)
		return nil, intelOwlError
	}

	intelOwlClientOptions := &IntelOwlClientOptions{}
	if unmarshalError := json.Unmarshal(optionsBytes, &intelOwlClientOptions); unmarshalError != nil {
		return nil, unmarshalError
	}

	intelOwlClient := NewIntelOwlClient(intelOwlClientOptions, httpClient, loggerParams)

	return &intelOwlClient, nil
}

// buildRequest is used for building requests.
func (client *IntelOwlClient) buildRequest(ctx context.Context, method string, contentType string, body io.Reader, url string) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType)

	tokenString := fmt.Sprintf("token %s", client.options.Token)

	request.Header.Set("Authorization", tokenString)
	return request, nil
}

// newRequest is used for making requests.
func (client *IntelOwlClient) newRequest(ctx context.Context, request *http.Request) (*successResponse, error) {
	response, err := client.client.Do(request)

	// Checking for context errors such as reaching the deadline and/or Timeout
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	defer response.Body.Close()

	msgBytes, err := ioutil.ReadAll(response.Body)
	statusCode := response.StatusCode
	if err != nil {
		errorMessage := fmt.Sprintf("Could not convert JSON response. Status code: %d", statusCode)
		intelOwlError := newIntelOwlError(statusCode, errorMessage, response)
		return nil, intelOwlError
	}

	if statusCode < http.StatusOK || statusCode >= http.StatusBadRequest {
		errorMessage := string(msgBytes)
		intelOwlError := newIntelOwlError(statusCode, errorMessage, response)
		return nil, intelOwlError
	}

	sucessResp := successResponse{
		StatusCode: statusCode,
		Data:       msgBytes,
	}

	return &sucessResp, nil
}
