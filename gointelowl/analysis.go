package gointelowl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type BasicAnalysisParams struct {
	User                 int                    `json:"user"`
	Tlp                  string                 `json:"tlp"`
	RuntimeConfiguration map[string]interface{} `json:"runtime_configuration"`
	AnalyzersRequested   []string               `json:"analyzers_requested"`
	ConnectorsRequested  []string               `json:"connectors_requested"`
	TagsLabels           []string               `json:"tags_labels"`
}

type ObservableAnalysisParams struct {
	BasicAnalysisParams
	ObservableName           string `json:"observable_name"`
	ObservableClassification string `json:"classification"`
}

type MultipleObservableAnalysisParams struct {
	BasicAnalysisParams
	Observables [][]string `json:"observables"`
}

type FileAnalysisParams struct {
	BasicAnalysisParams
}

type AnalysisResponse struct {
	JobID             int      `json:"job_id"`
	Status            string   `json:"status"`
	Warnings          []string `json:"warnings"`
	AnalyzersRunning  []string `json:"analyzers_running"`
	ConnectorsRunning []string `json:"connectors_running"`
}

type MultipleAnalysisResponse struct {
	Count   int                `json:"count"`
	Results []AnalysisResponse `json:"results"`
}

func (client *IntelOwlClient) CreateObservableAnalysis(ctx context.Context, params *ObservableAnalysisParams) (*AnalysisResponse, error) {
	requestUrl := fmt.Sprintf("%s/api/analyze_observable", client.options.Url)

	jsonData, _ := json.Marshal(params)

	fmt.Println("JSON DATA")
	fmt.Println(string(jsonData))

	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("???")
		return nil, err
	}

	analysisResponse := AnalysisResponse{}
	successResp, err := client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	if unmarshalError := json.Unmarshal(successResp.Data, &analysisResponse); unmarshalError != nil {
		return nil, unmarshalError
	}
	return &analysisResponse, nil

}

func (client *IntelOwlClient) CreateMultipleObservableAnalysis(ctx context.Context, params *MultipleObservableAnalysisParams) (*MultipleAnalysisResponse, error) {
	requestUrl := fmt.Sprintf("%s/api/analyze_multiple_observables", client.options.Url)

	jsonData, _ := json.Marshal(params)

	fmt.Println("JSON DATA")
	fmt.Println(string(jsonData))

	request, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("???")
		return nil, err
	}

	multipleAnalysisResponse := MultipleAnalysisResponse{}
	successResp, err := client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	if unmarshalError := json.Unmarshal(successResp.Data, &multipleAnalysisResponse); unmarshalError != nil {
		return nil, unmarshalError
	}
	return &multipleAnalysisResponse, nil
}

// func (client *IntelOwlClient) CreateFileAnalysis(ctx context.Context, *FileAnalysisParams) (*AnalysisResponse, error) {

// }
