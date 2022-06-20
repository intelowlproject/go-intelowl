package gointelowl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AnalyzerConfiguration struct {
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

type Analyzer struct {
	client *IntelOwlClient
}

func (analyzer *Analyzer) GetConfigs(ctx context.Context) (*[]AnalyzerConfiguration, error) {
	requestUrl := fmt.Sprintf("%s/api/get_analyzer_configs", analyzer.client.options.Url)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	successResp, err := analyzer.client.makeRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	analyzerConfigurationResponse := map[string]AnalyzerConfiguration{}
	json.Unmarshal(successResp.Data, &analyzerConfigurationResponse)
	analyzerConfigurationList := []AnalyzerConfiguration{}
	for _, analyzerConfig := range analyzerConfigurationResponse {
		analyzerConfigurationList = append(analyzerConfigurationList, analyzerConfig)
	}
	return &analyzerConfigurationList, nil

	// return &analyzerConfigurationList, nil
	// analyzerConfigurationList := []AnalyzerConfiguration{}
	// for _, analyzerConfig := range analyzerConfigurationResponse {
	// 	analyzerConfigurationList = append(analyzerConfigurationList, analyzerConfig)
	// }
	// requestUrl := fmt.Sprintf("%s/api/get_analyzer_configs", analyzer.client.options.Url)
	// fmt.Println(requestUrl)
	// request, err := http.NewRequest("GET", requestUrl, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// analyzerConfigurationResponse := map[string]AnalyzerConfiguration{}
	// if err := analyzer.client.makeRequest(ctx, request, &analyzerConfigurationResponse); err != nil {
	// 	return nil, err
	// }
	// analyzerConfigurationList := []AnalyzerConfiguration{}
	// for _, analyzerConfig := range analyzerConfigurationResponse {
	// 	analyzerConfigurationList = append(analyzerConfigurationList, analyzerConfig)
	// }

	// return &analyzerConfigurationList, nil
}

func (analyzerConfig *AnalyzerConfiguration) Display() error {
	jsonBytes, err := json.Marshal(analyzerConfig)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonBytes))
	return nil
}
