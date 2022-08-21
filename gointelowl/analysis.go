package gointelowl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// This serves as the basis for the analysis parameters: Observable and File analysis.
type BasicAnalysisParams struct {
	User                 int                    `json:"user"`
	Tlp                  TLP                    `json:"tlp"`
	RuntimeConfiguration map[string]interface{} `json:"runtime_configuration"`
	AnalyzersRequested   []string               `json:"analyzers_requested"`
	ConnectorsRequested  []string               `json:"connectors_requested"`
	TagsLabels           []string               `json:"tags_labels"`
}

// Observable analysis parameters struct to easily make an observable analysis!
type ObservableAnalysisParams struct {
	BasicAnalysisParams
	ObservableName           string `json:"observable_name"`
	ObservableClassification string `json:"classification"`
}

// Multiple observable analysis parameters struct to easily analyze multiple observables
type MultipleObservableAnalysisParams struct {
	BasicAnalysisParams
	Observables [][]string `json:"observables"`
}

// File Analysis parameters struct to easily make an file analysis!
// This is gonna be passed to multiform data! so no JSON tags
type FileAnalysisParams struct {
	BasicAnalysisParams
	File *os.File
}

// Multiple File Analysis parameters struct to easily analyze multiple files
// This is gonna be passed to multiform data! so no JSON tags
type MultipleFileAnalysisParams struct {
	BasicAnalysisParams
	Files []*os.File
}

// The response struct for making an analysis!
type AnalysisResponse struct {
	JobID             int      `json:"job_id"`
	Status            string   `json:"status"`
	Warnings          []string `json:"warnings"`
	AnalyzersRunning  []string `json:"analyzers_running"`
	ConnectorsRunning []string `json:"connectors_running"`
}

// The response struct when you analyze multiple observables or files
type MultipleAnalysisResponse struct {
	Count   int                `json:"count"`
	Results []AnalysisResponse `json:"results"`
}

// Desc: Analyze an observable(IP, String, Hash)
//
//	Endpoint: POST /api/analyze_observable
func (client *IntelOwlClient) CreateObservableAnalysis(ctx context.Context, params *ObservableAnalysisParams) (*AnalysisResponse, error) {
	requestUrl := fmt.Sprintf(ANALYZE_OBSERVABLE_URL, client.options.Url)
	method := "POST"
	contentType := "application/json"
	jsonData, _ := json.Marshal(params)
	body := bytes.NewBuffer(jsonData)

	request, err := client.buildRequest(ctx, method, contentType, body, requestUrl)
	if err != nil {
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

// Desc: Analyze multiple observables
//
//	Endpoint: POST /api/analyze_multiple_observables
func (client *IntelOwlClient) CreateMultipleObservableAnalysis(ctx context.Context, params *MultipleObservableAnalysisParams) (*MultipleAnalysisResponse, error) {
	requestUrl := fmt.Sprintf(ANALYZE_MULTIPLE_OBSERVABLES_URL, client.options.Url)

	method := "POST"
	contentType := "application/json"
	jsonData, _ := json.Marshal(params)
	body := bytes.NewBuffer(jsonData)

	request, err := client.buildRequest(ctx, method, contentType, body, requestUrl)
	if err != nil {
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

// Desc: Analyze a File (.txt, .jpeg, .csv)
//
//	Endpoint: POST /api/analyze_file
func (client *IntelOwlClient) CreateFileAnalysis(ctx context.Context, fileAnalysisParams *FileAnalysisParams) (*AnalysisResponse, error) {
	requestUrl := fmt.Sprintf(ANALYZE_FILE_URL, client.options.Url)
	// * Making the multiform data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// * Adding the TLP field
	writeTlpError := writer.WriteField("tlp", fileAnalysisParams.Tlp.String())
	if writeTlpError != nil {
		return nil, writeTlpError
	}
	// * Adding the runtimeconfiguration field
	runTimeConfigurationJson, marshalError := json.Marshal(fileAnalysisParams.RuntimeConfiguration)
	if marshalError != nil {
		return nil, marshalError
	}
	runTimeConfigurationJsonString := string(runTimeConfigurationJson)
	writeRuntimeError := writer.WriteField("runtime_configuration", runTimeConfigurationJsonString)
	if writeRuntimeError != nil {
		return nil, writeRuntimeError
	}

	// * Adding the requested analyzers
	for _, analyzer := range fileAnalysisParams.AnalyzersRequested {
		writeAnalyzerError := writer.WriteField("analyzers_requested", analyzer)
		if writeAnalyzerError != nil {
			return nil, writeAnalyzerError
		}
	}

	// * Adding the requested connectors
	for _, connector := range fileAnalysisParams.ConnectorsRequested {
		writeConnectorError := writer.WriteField("connectors_requested", connector)
		if writeConnectorError != nil {
			return nil, writeConnectorError
		}
	}

	// * Adding the tag labels
	for _, tagLabel := range fileAnalysisParams.TagsLabels {
		writeTagLabelError := writer.WriteField("tags_labels", tagLabel)
		if writeTagLabelError != nil {
			return nil, writeTagLabelError
		}
	}

	// * Adding the file!
	filePart, _ := writer.CreateFormFile("file", filepath.Base(fileAnalysisParams.File.Name()))
	_, writeFileError := io.Copy(filePart, fileAnalysisParams.File)
	if writeFileError != nil {
		writer.Close()
		return nil, writeFileError
	}
	writer.Close()

	//* building the request!
	contentType := writer.FormDataContentType()
	method := "POST"
	request, err := client.buildRequest(ctx, method, contentType, body, requestUrl)
	if err != nil {
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

// Desc: Analyze multiple files (.txt, .jpeg, .csv)
//
//	Endpoint: POST /api/analyze_mutliple_files
func (client *IntelOwlClient) CreateMultipleFileAnalysis(ctx context.Context, fileAnalysisParams *MultipleFileAnalysisParams) (*MultipleAnalysisResponse, error) {
	requestUrl := fmt.Sprintf(ANALYZE_MULTIPLE_FILES_URL, client.options.Url)
	// * Making the multiform data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// * Adding the TLP field
	writeTlpError := writer.WriteField("tlp", fileAnalysisParams.Tlp.String())
	if writeTlpError != nil {
		return nil, writeTlpError
	}
	// * Adding the runtimeconfiguration field
	runTimeConfigurationJson, marshalError := json.Marshal(fileAnalysisParams.RuntimeConfiguration)
	if marshalError != nil {
		return nil, marshalError
	}
	runTimeConfigurationJsonString := string(runTimeConfigurationJson)
	writeRuntimeError := writer.WriteField("runtime_configuration", runTimeConfigurationJsonString)
	if writeRuntimeError != nil {
		return nil, writeRuntimeError
	}

	// * Adding the requested analyzers
	for _, analyzer := range fileAnalysisParams.AnalyzersRequested {
		writeAnalyzerError := writer.WriteField("analyzers_requested", analyzer)
		if writeAnalyzerError != nil {
			return nil, writeAnalyzerError
		}
	}

	// * Adding the requested connectors
	for _, connector := range fileAnalysisParams.ConnectorsRequested {
		writeConnectorError := writer.WriteField("connectors_requested", connector)
		if writeConnectorError != nil {
			return nil, writeConnectorError
		}
	}

	// * Adding the tag labels
	for _, tagLabel := range fileAnalysisParams.TagsLabels {
		writeTagLabelError := writer.WriteField("tags_labels", tagLabel)
		if writeTagLabelError != nil {
			return nil, writeTagLabelError
		}
	}

	// * Adding the files!
	for _, file := range fileAnalysisParams.Files {
		filePart, _ := writer.CreateFormFile("files", filepath.Base(file.Name()))
		_, writeFileError := io.Copy(filePart, file)
		if writeFileError != nil {
			writer.Close()
			return nil, writeFileError
		}
	}
	writer.Close()

	//* building the request!
	contentType := writer.FormDataContentType()
	method := "POST"
	request, err := client.buildRequest(ctx, method, contentType, body, requestUrl)
	if err != nil {
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
