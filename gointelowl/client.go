package gointelowl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
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
	options          *IntelOwlClientOptions
	client           *http.Client
	TagService       *TagService
	JobService       *JobService
	AnalyzerService  *AnalyzerService
	ConnectorService *ConnectorService
	Logger           *IntelOwlLogger
}

// * enum for TLP attribute used in the IntelOwl API
type TLP int

// * making the values of TLP
const (
	WHITE TLP = iota + 1
	GREEN
	AMBER
	RED
)

// * To easily access the TLP enum values
var TLPVALUES = map[string]int{
	"WHITE": 1,
	"GREEN": 2,
	"AMBER": 3,
	"RED":   4,
}

// * Overriding the String method to get the string representation of the enum
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

// * To easily make the TLP enum
func ParseTLP(s string) TLP {
	s = strings.TrimSpace(s)
	value, ok := TLPVALUES[s]
	if !ok {
		return TLP(0)
	}
	return TLP(value)
}

// * Implementing the MarshalJSON interface to make our custom Marshal for the enum
func (tlp TLP) MarshalJSON() ([]byte, error) {
	return json.Marshal(tlp.String())
}

// * Implementing the UnmarshalJSON interface to make our custom Unmarshal for the enum
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

func NewIntelOwlClient(options *IntelOwlClientOptions, httpClient *http.Client, loggerParams *LoggerParams) IntelOwlClient {

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

	client.Logger = &IntelOwlLogger{}
	client.Logger.Init(loggerParams)

	return client
}

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

func (client *IntelOwlClient) newRequest(ctx context.Context, request *http.Request) (*successResponse, error) {
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
