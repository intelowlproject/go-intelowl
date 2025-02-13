package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/intelowlproject/go-intelowl/constants"
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
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			apiHandler.Handle(constants.ANALYZE_OBSERVABLE_URL, serverHandler(t, testCase, "POST"))
			observableParams, ok := testCase.Input.(gointelowl.ObservableAnalysisParams)
			if ok {
				gottenAnalysisResponse, err := client.CreateObservableAnalysis(ctx, &observableParams)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenAnalysisResponse)
				}
			}
		})
	}
}

func TestCreateObservablePlaybookAnalysis(t *testing.T) {
	playbookAnalysisJsonString := `{"results":[{"job_id":3002,"analyzers_running":["AdGuard","Classic_DNS","CloudFlare_DNS","CloudFlare_Malicious_Detector","DNS0_EU","DNS0_EU_Malicious_Detector","Google_DNS","Quad9_DNS","Quad9_Malicious_Detector","UltraDNS_DNS","UltraDNS_Malicious_Detector"],"connectors_running":[],"visualizers_running":["DNS"],"playbook_running":"Dns","investigation":null,"status":"accepted","already_exists":false}],"count":1}`
	playbookAnalysisResponse := gointelowl.MultipleAnalysisResponse{}
	if unmarshalError := json.Unmarshal([]byte(playbookAnalysisJsonString), &playbookAnalysisResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}

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
		Input: gointelowl.ObservablePlaybookAnalysisParams{
			BasicAnalysisParams:      basicAnalysisParams,
			ObservableName:           "series9.io",
			ObservableClassification: "domain",
			PlaybookRequested:        "Dns",
		},
		Data:       playbookAnalysisJsonString,
		StatusCode: http.StatusOK,
		Want:       &playbookAnalysisResponse,
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			apiHandler.Handle(constants.ANALYZE_OBSERVABLE_PLAYBOOK_URL, serverHandler(t, testCase, "POST"))
			ctx := context.Background()
			playbookObservableParams, ok := testCase.Input.(gointelowl.ObservablePlaybookAnalysisParams)
			if ok {
				gottenMultipleAnalysisResponse, err := client.CreateObservablePlaybookAnalysis(ctx, &playbookObservableParams)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenMultipleAnalysisResponse)
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
			client, apiHandler, closeServer := setup()
			defer closeServer()
			apiHandler.Handle(constants.ANALYZE_MULTIPLE_OBSERVABLES_URL, serverHandler(t, testCase, "POST"))
			ctx := context.Background()
			multipleObservableParams, ok := testCase.Input.(gointelowl.MultipleObservableAnalysisParams)
			if ok {
				gottenMultipleAnalysisResponse, err := client.CreateMultipleObservableAnalysis(ctx, &multipleObservableParams)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenMultipleAnalysisResponse)
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
			client, apiHandler, closeServer := setup()
			defer closeServer()
			apiHandler.Handle(constants.ANALYZE_FILE_URL, serverHandler(t, testCase, "POST"))
			ctx := context.Background()
			fileAnalysisParams, ok := testCase.Input.(gointelowl.FileAnalysisParams)
			if ok {
				gottenFileAnalysisResponse, err := client.CreateFileAnalysis(ctx, &fileAnalysisParams)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenFileAnalysisResponse)
				}
			}
		})
	}
}

func TestCreateFilesPlaybookAnalysis(t *testing.T) {
	playbookAnalysisJsonString := `{
  "results": [
    {
      "job_id": 3003,
      "analyzers_running": [
        "ClamAV",
        "Cymru_Hash_Registry_Get_File",
        "File_Info",
        "HashLookupServer_Get_File",
        "MalwareBazaar_Get_File",
        "Strings_Info",
        "YARAify_File_Search",
        "Yara"
      ],
      "connectors_running": [],
      "visualizers_running": [],
      "playbook_running": "Sample_Static_Analysis",
      "investigation": null,
      "status": "accepted",
      "already_exists": true
    }
  ],
  "count": 1
}
`
	playbookAnalysisResponse := gointelowl.MultipleAnalysisResponse{}
	if unmarshalError := json.Unmarshal([]byte(playbookAnalysisJsonString), &playbookAnalysisResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	fileDir := "./testFiles/"
	fileName := "fileForAnalysis.txt"
	filePath := path.Join(fileDir, fileName)
	file, _ := os.Open(filePath)
	defer file.Close()
	basicAnalysisParams := gointelowl.BasicAnalysisParams{
		User:                 1,
		Tlp:                  gointelowl.WHITE,
		RuntimeConfiguration: map[string]interface{}{},
		AnalyzersRequested:   []string{},
		ConnectorsRequested:  []string{},
		TagsLabels:           []string{},
	}
	playbookFileParams := &gointelowl.FilePlaybookAnalysisParams{
		BasicAnalysisParams: basicAnalysisParams,
		PlaybookRequested:   "Sample_Static_Analysis",
		File:                file,
	}
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      playbookFileParams,
		Data:       playbookAnalysisJsonString,
		StatusCode: http.StatusOK,
		Want:       playbookAnalysisResponse,
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			apiHandler.Handle(constants.ANALYZE_FILE_PLAYBOOK_URL, serverHandler(t, testCase, "POST"))
			ctx := context.Background()
			playbookFilesAnalysisParams, ok := testCase.Input.(gointelowl.FilePlaybookAnalysisParams)
			if ok {
				gottenFilePlaybookAnalysisResponse, err := client.CreateFilePlaybookAnalysis(ctx, &playbookFilesAnalysisParams)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenFilePlaybookAnalysisResponse)
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
			client, apiHandler, closeServer := setup()
			defer closeServer()
			apiHandler.Handle(constants.ANALYZE_MULTIPLE_FILES_URL, serverHandler(t, testCase, "POST"))
			ctx := context.Background()
			multipleFilesAnalysisParams, ok := testCase.Input.(gointelowl.MultipleFileAnalysisParams)
			if ok {
				gottenMultipleFilesAnalysisResponse, err := client.CreateMultipleFileAnalysis(ctx, &multipleFilesAnalysisParams)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenMultipleFilesAnalysisResponse)
				}
			}
		})
	}

}
