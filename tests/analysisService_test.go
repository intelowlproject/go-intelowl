package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/intelowlproject/go-intelowl/gointelowl"
)

func TestCreateObservableAnalysis(t *testing.T) {
	analysisJsonString := `{"job_id":260,"status":"accepted","warnings":[],"analyzers_running":["Classic_DNS","CryptoScamDB_CheckAPI","Darksearch_Query","FireHol_IPList","FileScan_Search","GoogleWebRisk","GreyNoiseCommunity","InQuest_IOCdb","InQuest_REPdb","InQuest_DFI","MalwareBazaar_Google_Observable","Mnemonic_PassiveDNS","Phishstats","Pulsedive_Active_IOC","Robtex_IP_Query","Robtex_Reverse_PDNS_Query","Stratosphere_Blacklist","TalosReputation","ThreatFox","Threatminer_PDNS","Threatminer_Reports_Tagging","TorProject","URLhaus","UrlScan_Search","WhoIs_RipeDB_Search","YETI"],"connectors_running":["YETI"]}`
	analysisResponse := gointelowl.AnalysisResponse{}
	if unmarshalError := json.Unmarshal([]byte(analysisJsonString), &analysisResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	basicParams := gointelowl.BasicAnalysisParams{
		User:                 1,
		Tlp:                  gointelowl.WHITE,
		RuntimeConfiguration: map[string]interface{}{},
		AnalyzersRequested:   []string{},
		ConnectorsRequested:  []string{},
		TagsLabels:           []string{},
	}
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input: gointelowl.ObservableAnalysisParams{
			BasicAnalysisParams:      basicParams,
			ObservableName:           "192.168.69.42",
			ObservableClassification: "",
		},
		Data:       analysisJsonString,
		StatusCode: http.StatusOK,
		Want:       &analysisResponse,
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testServer := NewTestServer(&testCase)
			defer testServer.Close()
			client := NewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			observableParams, ok := testCase.Input.(gointelowl.ObservableAnalysisParams)
			if ok {
				gottenAnalysisResponse, err := client.CreateObservableAnalysis(ctx, &observableParams)
				if testCase.StatusCode < http.StatusOK || testCase.StatusCode >= http.StatusBadRequest {
					diff := cmp.Diff(testCase.Want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
					if diff != "" {
						t.Fatalf(diff)
					}
				} else {
					diff := cmp.Diff(testCase.Want, gottenAnalysisResponse)
					if diff != "" {
						t.Fatalf(diff)
					}
				}
			}
		})
	}
}

func TestCreateMultipleObservableAnalysis(t *testing.T) {
	multiAnalysisJsonString := `{"count":2,"results":[{"job_id":263,"status":"accepted","warnings":[],"analyzers_running":["Classic_DNS","CryptoScamDB_CheckAPI","Darksearch_Query","FireHol_IPList","FileScan_Search","GoogleWebRisk","GreyNoiseCommunity","InQuest_IOCdb","InQuest_REPdb","InQuest_DFI","MalwareBazaar_Google_Observable","Mnemonic_PassiveDNS","Phishstats","Pulsedive_Active_IOC","Robtex_IP_Query","Robtex_Reverse_PDNS_Query","Stratosphere_Blacklist","TalosReputation","ThreatFox","Threatminer_PDNS","Threatminer_Reports_Tagging","TorProject","URLhaus","UrlScan_Search","WhoIs_RipeDB_Search","YETI"],"connectors_running":["YETI"]},{"job_id":264,"status":"accepted","warnings":[],"analyzers_running":["Classic_DNS","CryptoScamDB_CheckAPI","Darksearch_Query","FireHol_IPList","FileScan_Search","GoogleWebRisk","GreyNoiseCommunity","InQuest_IOCdb","InQuest_REPdb","InQuest_DFI","MalwareBazaar_Google_Observable","Mnemonic_PassiveDNS","Phishstats","Pulsedive_Active_IOC","Robtex_IP_Query","Robtex_Reverse_PDNS_Query","Stratosphere_Blacklist","TalosReputation","ThreatFox","Threatminer_PDNS","Threatminer_Reports_Tagging","TorProject","URLhaus","UrlScan_Search","WhoIs_RipeDB_Search","YETI"],"connectors_running":["YETI"]}]}`
	multiAnalysisResponse := gointelowl.MultipleAnalysisResponse{}
	if unmarshalError := json.Unmarshal([]byte(multiAnalysisJsonString), &multiAnalysisResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	observables := make([][]string, 2)
	observables[0] = make([]string, 2)
	observables[0][0] = "ip"
	observables[0][1] = "8.8.8.8"
	observables[1] = make([]string, 2)
	observables[1][0] = "ip"
	observables[1][1] = "8.8.8.7"
	basicAnalysisParams := gointelowl.BasicAnalysisParams{
		User:                 1,
		Tlp:                  gointelowl.WHITE,
		RuntimeConfiguration: map[string]interface{}{},
		AnalyzersRequested:   []string{},
		ConnectorsRequested:  []string{},
		TagsLabels:           []string{},
	}

	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input: gointelowl.MultipleObservableAnalysisParams{
			BasicAnalysisParams: basicAnalysisParams,
			Observables:         observables,
		},
		Data:       multiAnalysisJsonString,
		StatusCode: http.StatusOK,
		Want:       &multiAnalysisResponse,
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testServer := NewTestServer(&testCase)
			defer testServer.Close()
			client := NewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			multipleObservableParams, ok := testCase.Input.(gointelowl.MultipleObservableAnalysisParams)
			if ok {
				gottenMultipleAnalysisResponse, err := client.CreateMultipleObservableAnalysis(ctx, &multipleObservableParams)
				if testCase.StatusCode < http.StatusOK || testCase.StatusCode >= http.StatusBadRequest {
					diff := cmp.Diff(testCase.Want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
					if diff != "" {
						t.Fatalf(diff)
					}
				} else {
					diff := cmp.Diff(testCase.Want, gottenMultipleAnalysisResponse)
					if diff != "" {
						t.Fatalf(diff)
					}
				}
			}
		})
	}

}

func TestCreateFileAnalysis(t *testing.T) {
	analysisJsonString := `{"job_id":269,"status":"accepted","warnings":[],"analyzers_running":["File_Info"],"connectors_running":["YETI"]}`
	analysisResponse := gointelowl.AnalysisResponse{}
	if unmarshalError := json.Unmarshal([]byte(analysisJsonString), &analysisResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	fileName := "fileForAnalysis.txt"
	fileDir := "./testFiles/"
	filePath := path.Join(fileDir, fileName)
	file, _ := os.Open(filePath)
	defer file.Close()
	basicAnalysisParams := gointelowl.BasicAnalysisParams{
		User:                 1,
		Tlp:                  gointelowl.WHITE,
		RuntimeConfiguration: map[string]interface{}{},
		AnalyzersRequested:   []string{"File_Info"},
		ConnectorsRequested:  []string{},
		TagsLabels:           []string{},
	}
	fileParams := &gointelowl.FileAnalysisParams{
		BasicAnalysisParams: basicAnalysisParams,
		File:                file,
	}
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      fileParams,
		Data:       analysisJsonString,
		StatusCode: http.StatusOK,
		Want:       &analysisResponse,
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testServer := NewTestServer(&testCase)
			defer testServer.Close()
			client := NewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			fileAnalysisParams, ok := testCase.Input.(gointelowl.FileAnalysisParams)
			if ok {
				gottenFileAnalysisResponse, err := client.CreateFileAnalysis(ctx, &fileAnalysisParams)
				if testCase.StatusCode < http.StatusOK || testCase.StatusCode >= http.StatusBadRequest {
					diff := cmp.Diff(testCase.Want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
					if diff != "" {
						t.Fatalf(diff)
					}
				} else {
					diff := cmp.Diff(testCase.Want, gottenFileAnalysisResponse)
					if diff != "" {
						t.Fatalf(diff)
					}
				}
			}
		})
	}
}

func TestCreateMultipleFilesAnalysis(t *testing.T) {
	multiAnalysisJsonString := `{"count":2,"results":[{"job_id":270,"status":"accepted","warnings":[],"analyzers_running":["File_Info"],"connectors_running":["YETI"]},{"job_id":271,"status":"accepted","warnings":[],"analyzers_running":["File_Info"],"connectors_running":["YETI"]}]}`
	multiAnalysisResponse := gointelowl.MultipleAnalysisResponse{}
	if unmarshalError := json.Unmarshal([]byte(multiAnalysisJsonString), &multiAnalysisResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	fileDir := "./testFiles/"
	fileName := "fileForAnalysis.txt"
	filePath := path.Join(fileDir, fileName)
	file, _ := os.Open(filePath)
	defer file.Close()
	fileName2 := "fileForAnalysis.txt"
	filePath2 := path.Join(fileDir, fileName2)
	file2, _ := os.Open(filePath2)
	defer file2.Close()
	filesArray := make([]*os.File, 2)
	filesArray[0] = file
	filesArray[1] = file2
	basicAnalysisParams := gointelowl.BasicAnalysisParams{
		User:                 1,
		Tlp:                  gointelowl.WHITE,
		RuntimeConfiguration: map[string]interface{}{},
		AnalyzersRequested:   []string{"File_Info"},
		ConnectorsRequested:  []string{},
		TagsLabels:           []string{},
	}
	multipleFileParams := &gointelowl.MultipleFileAnalysisParams{
		BasicAnalysisParams: basicAnalysisParams,
		Files:               filesArray,
	}
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      multipleFileParams,
		Data:       multiAnalysisJsonString,
		StatusCode: http.StatusOK,
		Want:       multiAnalysisResponse,
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testServer := NewTestServer(&testCase)
			defer testServer.Close()
			client := NewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			multipleFilesAnalysisParams, ok := testCase.Input.(gointelowl.MultipleFileAnalysisParams)
			if ok {
				gottenMultipleFilesAnalysisResponse, err := client.CreateMultipleFileAnalysis(ctx, &multipleFilesAnalysisParams)
				if testCase.StatusCode < http.StatusOK || testCase.StatusCode >= http.StatusBadRequest {
					diff := cmp.Diff(testCase.Want, err, cmpopts.IgnoreFields(gointelowl.IntelOwlError{}, "Response"))
					if diff != "" {
						t.Fatalf(diff)
					}
				} else {
					diff := cmp.Diff(testCase.Want, gottenMultipleFilesAnalysisResponse)
					if diff != "" {
						t.Fatalf(diff)
					}
				}
			}
		})
	}

}
