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

type BasicAnalysisParams struct {
	User                 int                    `json:"user"`
	Tlp                  TLP                    `json:"tlp"`
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

//* This is gonna be passed to multiform data! so no JSON tags
type FileAnalysisParams struct {
	BasicAnalysisParams
	File *os.File
}

type MultipleFileAnalysisParams struct {
	BasicAnalysisParams
	Files []*os.File
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

/*
* Desc: Create an analysis of an observable(IP, String, Hash)
* Endpoint: POST /api/analyze_observable
 */
func (client *IntelOwlClient) CreateObservableAnalysis(ctx context.Context, params *ObservableAnalysisParams) (*AnalysisResponse, error) {
	requestUrl := fmt.Sprintf("%s/api/analyze_observable", client.options.Url)
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

/*
* Desc: Create analysis of many observables
* Endpoint: POST /api/analyze_multiple_observables
 */
func (client *IntelOwlClient) CreateMultipleObservableAnalysis(ctx context.Context, params *MultipleObservableAnalysisParams) (*MultipleAnalysisResponse, error) {
	requestUrl := fmt.Sprintf("%s/api/analyze_multiple_observables", client.options.Url)

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

/*
* Desc: Create an analysis of a File (.txt, .jpeg, .csv)
* Endpoint: POST /api/analyze_file
 */
func (client *IntelOwlClient) CreateFileAnalysis(ctx context.Context, fileAnalysisParams *FileAnalysisParams) (*AnalysisResponse, error) {
	requestUrl := fmt.Sprintf("%s/api/analyze_file", client.options.Url)
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

/*
* Desc: Create an analysis of multiple Files (.txt, .jpeg, .csv)
* Endpoint: POST /api/analyze_mutliple_files
 */
func (client *IntelOwlClient) CreateMultipleFileAnalysis(ctx context.Context, fileAnalysisParams *MultipleFileAnalysisParams) (*MultipleAnalysisResponse, error) {
	requestUrl := fmt.Sprintf("%s/api/analyze_multiple_files", client.options.Url)
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
