package gointelowl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UserDetails struct {
	Username string `json:"username"`
}

type Report struct {
	Name                 string                 `json:"name"`
	Status               string                 `json:"status"`
	Report               map[string]interface{} `json:"report"`
	Errors               []string               `json:"errors"`
	ProcessTime          float64                `json:"process_time"`
	StartTime            time.Time              `json:"start_time"`
	EndTime              time.Time              `json:"end_time"`
	RuntimeConfiguration map[string]interface{} `json:"runtime_configuration"`
	Type                 string                 `json:"type"`
}

type BaseJob struct {
	ID                       int         `json:"id"`
	User                     UserDetails `json:"user"`
	Tags                     []Tag       `json:"tags"`
	ProcessTime              float64     `json:"process_time"`
	IsSample                 bool        `json:"is_sample"`
	Md5                      string      `json:"md5"`
	ObservableName           string      `json:"observable_name"`
	ObservableClassification string      `json:"observable_classification"`
	FileName                 string      `json:"file_name"`
	FileMimetype             string      `json:"file_mimetype"`
	Status                   string      `json:"status"`
	AnalyzersRequested       []string    `json:"analyzers_requested" `
	ConnectorsRequested      []string    `json:"connectors_requested"`
	AnalyzersToExecute       []string    `json:"analyzers_to_execute"`
	ConnectorsToExecute      []string    `json:"connectors_to_execute"`
	ReceivedRequestTime      *time.Time  `json:"received_request_time"`
	FinishedAnalysisTime     *time.Time  `json:"finished_analysis_time"`
	Tlp                      string      `json:"tlp"`
	Errors                   []string    `json:"errors"`
}

// This is to represent a job
type Job struct {
	BaseJob
	AnalyzerReports  []Report               `json:"analyzer_reports"`
	ConnectorReports []Report               `json:"connector_reports"`
	Permission       map[string]interface{} `json:"permission"`
}

// This is to represent the jobs which come as a list
type JobList struct {
	BaseJob
}

type JobListResponse struct {
	Count      int       `json:"count"`
	TotalPages int       `json:"total_pages"`
	Results    []JobList `json:"results"`
}

// Service object to access the Job endpoints!
type JobService struct {
	client *IntelOwlClient
}

// Getting a list of all the jobs
//
//	GET /api/jobs
func (jobService *JobService) List(ctx context.Context) (*JobListResponse, error) {
	requestUrl := fmt.Sprintf(BASE_JOB_URL, jobService.client.options.Url)
	contentType := "application/json"
	method := "GET"
	request, err := jobService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return nil, err
	}
	successResp, err := jobService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	jobList := JobListResponse{}
	marashalError := json.Unmarshal(successResp.Data, &jobList)
	if marashalError != nil {
		return nil, marashalError
	}

	return &jobList, nil
}

// Get a Job through its respective ID
//
//	Endpoint: GET /api/jobs/{jobID}
func (jobService *JobService) Get(ctx context.Context, jobId uint64) (*Job, error) {
	requestUrl := fmt.Sprintf(SPECIFIC_JOB_URL, jobService.client.options.Url, jobId)
	contentType := "application/json"
	method := "GET"
	request, err := jobService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return nil, err
	}
	successResp, err := jobService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	jobResponse := Job{}
	unmarshalError := json.Unmarshal(successResp.Data, &jobResponse)
	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &jobResponse, nil
}

// Get the File Sample associated with a Job through its ID
//
//	GET /api/jobs/{jobID}/download_sample
func (jobService *JobService) DownloadSample(ctx context.Context, jobId uint64) ([]byte, error) {
	requestUrl := fmt.Sprintf(DOWNLOAD_SAMPLE_JOB_URL, jobService.client.options.Url, jobId)
	contentType := "application/json"
	method := "GET"
	request, err := jobService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return nil, err
	}
	successResp, err := jobService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	return successResp.Data, nil
}

// Delete a Job through its ID
//
//	DELETE /api/jobs/{jobID}
func (jobService *JobService) Delete(ctx context.Context, jobId uint64) (bool, error) {
	requestUrl := fmt.Sprintf(SPECIFIC_JOB_URL, jobService.client.options.Url, jobId)
	contentType := "application/json"
	method := "DELETE"
	request, err := jobService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return false, err
	}
	successResp, err := jobService.client.newRequest(ctx, request)
	if err != nil {
		return false, err
	}
	if successResp.StatusCode == http.StatusNoContent {
		return true, nil
	}
	return false, nil
}

// Stop a running job through its ID
//
//	PATCH /api/jobs/{jobID}/kill
func (jobService *JobService) Kill(ctx context.Context, jobId uint64) (bool, error) {
	requestUrl := fmt.Sprintf(KILL_JOB_URL, jobService.client.options.Url, jobId)
	contentType := "application/json"
	method := "PATCH"
	request, err := jobService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return false, err
	}
	successResp, err := jobService.client.newRequest(ctx, request)
	if err != nil {
		return false, err
	}
	if successResp.StatusCode == http.StatusNoContent {
		return true, nil
	}
	return false, nil
}

// Stop a running analyzer on a job that is being processed through its ID and the analyzer's name
//
//	PATCH /api/jobs/{jobID}/analyzer/{nameOfAnalyzer}/kill
func (jobService *JobService) KillAnalyzer(ctx context.Context, jobId uint64, analyzerName string) (bool, error) {
	requestUrl := fmt.Sprintf(KILL_ANALYZER_JOB_URL, jobService.client.options.Url, jobId, analyzerName)
	contentType := "application/json"
	method := "PATCH"
	request, err := jobService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return false, err
	}
	successResp, err := jobService.client.newRequest(ctx, request)
	if err != nil {
		return false, err
	}
	if successResp.StatusCode == http.StatusNoContent {
		return true, nil
	}
	return false, nil
}

// Re-run a selected analyzer on a job that is being processed through its ID and the analyzer's name
//
//	PATCH /api/jobs/{jobID}/analyzer/{nameOfAnalyzer}/retry
func (jobService *JobService) RetryAnalyzer(ctx context.Context, jobId uint64, analyzerName string) (bool, error) {
	requestUrl := fmt.Sprintf(RETRY_ANALYZER_JOB_URL, jobService.client.options.Url, jobId, analyzerName)
	contentType := "application/json"
	method := "PATCH"
	request, err := jobService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return false, err
	}
	successResp, err := jobService.client.newRequest(ctx, request)
	if err != nil {
		return false, err
	}
	if successResp.StatusCode == http.StatusNoContent {
		return true, nil
	}
	return false, nil
}

// Stopping a running connector on a job that is being processed through its ID and the connector's name
//
//	PATCH /api/jobs/{jobID}/connector/{nameOfConnector}/kill
func (jobService *JobService) KillConnector(ctx context.Context, jobId uint64, connectorName string) (bool, error) {
	requestUrl := fmt.Sprintf(KILL_CONNECTOR_JOB_URL, jobService.client.options.Url, jobId, connectorName)
	contentType := "application/json"
	method := "PATCH"
	request, err := jobService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return false, err
	}
	successResp, err := jobService.client.newRequest(ctx, request)
	if err != nil {
		return false, err
	}
	if successResp.StatusCode == http.StatusNoContent {
		return true, nil
	}
	return false, nil
}

// Re-run a selected connector on a job that is being processed through its ID and the connector's name
//
//	PATCH /api/jobs/{jobID}/connector/{nameOfConnector}/retry
func (jobService *JobService) RetryConnector(ctx context.Context, jobId uint64, connectorName string) (bool, error) {
	requestUrl := fmt.Sprintf(RETRY_CONNECTOR_JOB_URL, jobService.client.options.Url, jobId, connectorName)
	contentType := "application/json"
	method := "PATCH"
	request, err := jobService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return false, err
	}
	successResp, err := jobService.client.newRequest(ctx, request)
	if err != nil {
		return false, err
	}
	if successResp.StatusCode == http.StatusNoContent {
		return true, nil
	}
	return false, nil
}
