package gointelowl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type JobTag struct {
	id    int    `json:"id"`
	label string `json:"label"`
	color string `json:"color"`
}

type JobTags []JobTag

type Job struct {
	id                        int     `json:"id"`
	is_sample                 bool    `json:"is_sample"`
	observable_name           string  `json:"observable_name"`
	observable_classification string  `json:"observable_classification"`
	file_name                 string  `json:"file_name"`
	file_mimetype             string  `json:"file_mimetype"`
	status                    string  `json:"status"`
	tags                      JobTags `json:"tags"`
	process_time              int     `json:"process_time"`
	no_of_analyzers_executed  string  `json:"no_of_analyzers_executed"`
	no_of_connectors_executed string  `json:"no_of_connectors_executed"`
}

type JobList []Job

func (c *IntelOwlClient) GetAllJobs(ctx context.Context) (JobList, error) {
	var jobs_json JobList

	// Make API request to /api/jobs
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/jobs", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	responded, err := c.sendRequest(req)

	if err != nil {
		return nil, err
	}

	jobs_bytes := []byte(responded.Data)
	json.Unmarshal(jobs_bytes, &jobs_json)

	return jobs_json, nil
}
