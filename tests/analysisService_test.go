package tests

import (
	"context"
	"encoding/json"
	"net/http"
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
	testCases := make(map[string]TestData)
	basicParams := gointelowl.BasicAnalysisParams{
		User:                 1,
		Tlp:                  gointelowl.WHITE,
		RuntimeConfiguration: map[string]interface{}{},
		AnalyzersRequested:   []string{},
		ConnectorsRequested:  []string{},
		TagsLabels:           []string{},
	}
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
