// go-intelowl provides an SDK to easily integrate intelowl with your own set of tools.

// go-intelowl makes it easy to automate, configure, and use intelowl with your own set of tools
// with its Idiomatic approach making an analysis is easy as just writing one line of code!
package gointelowl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Error represents an error that has occurred when communicating with IntelOwl.
type Error struct {
	StatusCode int
	Message    string
	Response   *http.Response
}

// Error lets you implement the error interface.
// This is used for making custom go errors.
func (intelOwlError *Error) Error() string {
	errorMessage := fmt.Sprintf("Status Code: %d \n Error: %s", intelOwlError.StatusCode, intelOwlError.Message)
	return errorMessage
}

// newError lets you easily create new Errors.
func newError(statusCode int, message string, response *http.Response) *Error {
	return &Error{
		StatusCode: statusCode,
		Message:    message,
		Response:   response,
	}
}

type successResponse struct {
	StatusCode int
	Data       []byte
}

// ClientOptions represents the fields needed to configure and use the Client
type ClientOptions struct {
	Url   string `json:"url"`
	Token string `json:"token"`
	// Certificate represents your SSL cert: path to the cert file!
	Certificate string `json:"certificate"`
	// Timeout is in seconds
	Timeout uint64 `json:"timeout"`
}

// Client handles all the communication with your IntelOwl instance.
type Client struct {
	options          *ClientOptions
	client           *http.Client
	TagService       *TagService
	JobService       *JobService
	AnalyzerService  *AnalyzerService
	ConnectorService *ConnectorService
	UserService      *UserService
	Logger           *Logger
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
func (tlp *TLP) MarshalJSON() ([]byte, error) {
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

// NewClient lets you easily create a new Client by providing ClientOptions, http.Clients, and LoggerParams.
func NewClient(options *ClientOptions, httpClient *http.Client, loggerParams *LoggerParams) Client {

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
	client := Client{
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
	client.Logger = &Logger{}
	client.Logger.Init(loggerParams)

	return client
}

// NewClientFromJsonFile lets you create a new Client through a JSON file that contains your ClientOptions
func NewClientFromJsonFile(filePath string, httpClient *http.Client, loggerParams *LoggerParams) (*Client, error) {
	optionsBytes, err := os.ReadFile(filePath)
	if err != nil {
		errorMessage := fmt.Sprintf("Could not read %s", filePath)
		intelOwlError := newError(400, errorMessage, nil)
		return nil, intelOwlError
	}

	intelOwlClientOptions := &ClientOptions{}
	if unmarshalError := json.Unmarshal(optionsBytes, &intelOwlClientOptions); unmarshalError != nil {
		return nil, unmarshalError
	}

	intelOwlClient := NewClient(intelOwlClientOptions, httpClient, loggerParams)

	return &intelOwlClient, nil
}

// buildRequest is used for building requests.
func (client *Client) buildRequest(ctx context.Context, method string, contentType string, body io.Reader, url string) (*http.Request, error) {
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
func (client *Client) newRequest(ctx context.Context, request *http.Request) (*successResponse, error) {
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

	msgBytes, err := io.ReadAll(response.Body)
	statusCode := response.StatusCode
	if err != nil {
		errorMessage := fmt.Sprintf("Could not convert JSON response. Status code: %d", statusCode)
		intelOwlError := newError(statusCode, errorMessage, response)
		return nil, intelOwlError
	}

	if statusCode < http.StatusOK || statusCode >= http.StatusBadRequest {
		errorMessage := string(msgBytes)
		intelOwlError := newError(statusCode, errorMessage, response)
		return nil, intelOwlError
	}

	sucessResp := successResponse{
		StatusCode: statusCode,
		Data:       msgBytes,
	}

	return &sucessResp, nil
}
