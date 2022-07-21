package gointelowl

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
)

type AnalyzerConfig struct {
	BaseConfigurationType
	Type                  string   `json:"type"`
	ExternalService       bool     `json:"external_service"`
	LeaksInfo             bool     `json:"leaks_info"`
	DockerBased           bool     `json:"docker_based"`
	RunHash               bool     `json:"run_hash"`
	RunHashType           string   `json:"run_hash_type"`
	SupportedFiletypes    []string `json:"supported_filetypes"`
	NotSupportedFiletypes []string `json:"not_supported_filetypes"`
	ObservableSupported   []string `json:"observable_supported"`
}

type AnalyzerService struct {
	client *IntelOwlClient
}

/*
* Desc: Getting the Analyzer Configurations
* Endpoint: GET "/api/get_analyzer_configs"
 */
func (analyzerService *AnalyzerService) GetConfigs(ctx context.Context) (*[]AnalyzerConfig, error) {
	requestUrl := fmt.Sprintf("%s/api/get_analyzer_configs", analyzerService.client.options.Url)
	contentType := "application/json"
	method := "GET"
	request, err := analyzerService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return nil, err
	}

	successResp, err := analyzerService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	analyzerConfigurationResponse := map[string]AnalyzerConfig{}
	if unmarshalError := json.Unmarshal(successResp.Data, &analyzerConfigurationResponse); unmarshalError != nil {
		return nil, unmarshalError
	}

	analyzerNames := make([]string, 0)
	// *getting all the analyzer key names!
	for analyzerName := range analyzerConfigurationResponse {
		analyzerNames = append(analyzerNames, analyzerName)
	}
	// * sorting them alphabetically
	sort.Strings(analyzerNames)
	analyzerConfigurationList := []AnalyzerConfig{}
	for _, analyzerName := range analyzerNames {
		analyzerConfig := analyzerConfigurationResponse[analyzerName]
		analyzerConfigurationList = append(analyzerConfigurationList, analyzerConfig)
	}
	return &analyzerConfigurationList, nil
}

/*
* Desc: Getting the Analyzer Configurations
* Endpoint: GET "/api/analyzer/{NameOfAnalyzer}/healthcheck"
 */
func (analyzerService *AnalyzerService) HealthCheck(ctx context.Context, analyzerName string) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/analyzer/%s/healthcheck", analyzerService.client.options.Url, analyzerName)
	contentType := "application/json"
	method := "GET"
	request, err := analyzerService.client.buildRequest(ctx, method, contentType, nil, requestUrl)
	if err != nil {
		return false, err
	}
	status := StatusResponse{}
	successResp, err := analyzerService.client.newRequest(ctx, request)
	if err != nil {
		return false, err
	}
	if unmarshalError := json.Unmarshal(successResp.Data, &status); unmarshalError != nil {
		return false, unmarshalError
	}
	return status.Status, nil
}
