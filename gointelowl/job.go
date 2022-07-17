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

type Job struct {
	ID                       int                    `json:"id"`
	User                     UserDetails            `json:"user"`
	Tags                     []Tag                  `json:"tags"`
	ProcessTime              float64                `json:"process_time"`
	AnalyzerReports          []Report               `json:"analyzer_reports,omitempty"`
	ConnectorReports         []Report               `json:"connector_reports,omitempty"`
	Permission               map[string]interface{} `json:"permission,omitempty"`
	IsSample                 bool                   `json:"is_sample"`
	Md5                      string                 `json:"md5"`
	ObservableName           string                 `json:"observable_name"`
	ObservableClassification string                 `json:"observable_classification"`
	FileName                 string                 `json:"file_name"`
	FileMimetype             string                 `json:"file_mimetype"`
	Status                   string                 `json:"status"`
	AnalyzersRequested       []string               `json:"analyzers_requested" `
	ConnectorsRequested      []string               `json:"connectors_requested"`
	AnalyzersToExecute       []string               `json:"analyzers_to_execute"`
	ConnectorsToExecute      []string               `json:"connectors_to_execute"`
	ReceivedRequestTime      *time.Time             `json:"received_request_time"`
	FinishedAnalysisTime     *time.Time             `json:"finished_analysis_time"`
	Tlp                      string                 `json:"tlp"`
	Errors                   []string               `json:"errors"`
}

type JobList struct {
	Count      int   `json:"count"`
	TotalPages int   `json:"total_pages"`
	Results    []Job `json:"results"`
}

type JobService struct {
	client *IntelOwlClient
}

/*
* Desc: Get a list of all the jobs
* Endpoint: GET /api/jobs
 */
func (jobService *JobService) List(ctx context.Context) (*JobList, error) {
	requestUrl := fmt.Sprintf("%s/api/jobs", jobService.client.options.Url)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	successResp, err := jobService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	jobList := JobList{}
	marashalError := json.Unmarshal(successResp.Data, &jobList)
	if marashalError != nil {
		return nil, marashalError
	}

	return &jobList, nil
}

/*
* Desc: Get a Job through its respective ID
* Endpoint: GET /api/jobs/{jobID}
 */
func (jobService *JobService) Get(ctx context.Context, jobId uint64) (*Job, error) {
	requestUrl := fmt.Sprintf("%s/api/jobs/%d", jobService.client.options.Url, jobId)
	request, err := http.NewRequest("GET", requestUrl, nil)
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

/*
* Desc: Get the File Sample associated with a Job through its ID
* Endpoint: GET /api/jobs/{jobID}/download_sample
 */
func (jobService *JobService) DownloadSample(ctx context.Context, jobId uint64) ([]byte, error) {
	requestUrl := fmt.Sprintf("%s/api/jobs/%d/download_sample", jobService.client.options.Url, jobId)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}
	successResp, err := jobService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	return successResp.Data, nil
}

/*
* Desc: Delete a Job through its ID
* Endpoint: DELETE /api/jobs/{jobID}
 */
func (jobService *JobService) Delete(ctx context.Context, jobId uint64) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/jobs/%d", jobService.client.options.Url, jobId)
	request, err := http.NewRequest("DELETE", requestUrl, nil)
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

/*
* Desc: Stop a running job through its ID
* Endpoint: PATCH /api/jobs/{jobID}/kill
 */
func (jobService *JobService) Kill(ctx context.Context, jobId uint64) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/jobs/%d/kill", jobService.client.options.Url, jobId)
	request, err := http.NewRequest("PATCH", requestUrl, nil)
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

/*
* Desc: Stop a running analyzer on a job that is being processed through its ID and the analyzer's name
* Endpoint: PATCH /api/jobs/{jobID}/analyzer/{nameOfAnalyzer}/kill
 */
func (jobService *JobService) KillAnalyzer(ctx context.Context, jobId uint64, analyzerName string) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/jobs/%d/analyzer/%s/kill", jobService.client.options.Url, jobId, analyzerName)
	request, err := http.NewRequest("PATCH", requestUrl, nil)
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

/*
* Desc: Re-run a selected analyzer on a job that is being processed through its ID and the analyzer's name
* Endpoint: PATCH /api/jobs/{jobID}/analyzer/{nameOfAnalyzer}/retry
 */
func (jobService *JobService) RetryAnalyzer(ctx context.Context, jobId uint64, analyzerName string) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/jobs/%d/analyzer/%s/retry", jobService.client.options.Url, jobId, analyzerName)
	request, err := http.NewRequest("PATCH", requestUrl, nil)
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

/*
* Desc: Stop a running connector on a job that is being processed through its ID and the connector's name
* Endpoint: PATCH /api/jobs/{jobID}/connector/{nameOfConnector}/kill
 */
func (jobService *JobService) KillConnector(ctx context.Context, jobId uint64, connectorName string) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/jobs/%d/connector/%s/kill", jobService.client.options.Url, jobId, connectorName)
	request, err := http.NewRequest("PATCH", requestUrl, nil)
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

/*
* Desc: Re-run a selected connector on a job that is being processed through its ID and the connector's name
* Endpoint: PATCH /api/jobs/{jobID}/connector/{nameOfConnector}/retry
 */
func (jobService *JobService) RetryConnector(ctx context.Context, jobId uint64, connectorName string) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/jobs/%d/connector/%s/retry", jobService.client.options.Url, jobId, connectorName)
	request, err := http.NewRequest("PATCH", requestUrl, nil)
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
