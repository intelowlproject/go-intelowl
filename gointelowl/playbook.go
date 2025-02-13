package gointelowl

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/intelowlproject/go-intelowl/constants"
)

type PlaybookConfig struct {
	ID                   int64                  `json:"id"`
	Name                 string                 `json:"name"`
	Type                 []string               `json:"type"` // ChoiceArrayField equivalent
	Analyzers            []string               `json:"analyzers"`
	Connectors           []string               `json:"connectors"`
	Pivots               []string               `json:"pivots"`
	RuntimeConfiguration map[string]interface{} `json:"runtime_configuration"`
	ScanMode             int64                  `json:"scan_mode"`
	ScanCheckTime        string                 `json:"scan_check_time"`
	Tags                 []string               `json:"tags"`
	TLP                  string                 `json:"tlp"`
	Starting             bool                   `json:"starting"`
	Owner                string                 `json:"owner"` // OwnershipAbstractModel equivalent
	Disabled             bool                   `json:"disabled"`
}

type PlaybookListResponse struct {
	Count      int              `json:"count"`
	TotalPages int              `json:"total_pages"`
	Results    []PlaybookConfig `json:"results"`
}

type PlaybookService struct {
	client *Client
}

func (playbookService *PlaybookService) ListPlaybooks(ctx context.Context) (*PlaybookListResponse, error) {
	requestUrl := playbookService.client.options.Url + constants.BASE_PLAYBOOK_URL
	contentType := constants.ContentTypeJSON
	method := http.MethodGet
	request, err := playbookService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return nil, err
	}
	successResp, err := playbookService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	playbookList := PlaybookListResponse{}
	marshalError := json.Unmarshal(successResp.Data, &playbookList)
	if marshalError != nil {
		return nil, marshalError
	}

	return &playbookList, nil
}

func (playbookService *PlaybookService) GetPlaybookByName(ctx context.Context, playbookName string) (*PlaybookConfig, error) {
	requestUrl := playbookService.client.options.Url + constants.BASE_PLAYBOOK_URL + "/" + playbookName
	contentType := constants.ContentTypeJSON
	method := http.MethodGet
	request, err := playbookService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return nil, err
	}
	successResp, err := playbookService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	playbook := PlaybookConfig{}
	marshalError := json.Unmarshal(successResp.Data, &playbook)
	if marshalError != nil {
		return nil, marshalError
	}

	return &playbook, nil
}
