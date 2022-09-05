package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/intelowlproject/go-intelowl/constants"
	"github.com/intelowlproject/go-intelowl/gointelowl"
)

func TestJobServiceList(t *testing.T) {
	jobListJson := `{"count":8,"total_pages":1,"results":[{"id":74,"user":{"username":"hussain"},"tags":[],"process_time":102.05,"is_sample":false,"md5":"91c0c59c8f6fc9aa2dc99a89f2fd0ab5","observable_name":"string2","observable_classification":"generic","file_name":"","file_mimetype":"","status":"reported_with_fails","analyzers_requested":[],"connectors_requested":[],"analyzers_to_execute":["CryptoScamDB_CheckAPI","CRXcavator","CyberChef","Darksearch_Query","FileScan_Search","InQuest_IOCdb","InQuest_REPdb","InQuest_DFI","Phishstats","ThreatFox","YARAify_Generics","YETI"],"connectors_to_execute":["YETI"],"received_request_time":"2022-07-15T20:54:48.734361Z","finished_analysis_time":"2022-07-15T20:56:30.779578Z","tlp":"WHITE","errors":[]},{"id":73,"user":{"username":"hussain"},"tags":[],"process_time":73.68,"is_sample":false,"md5":"54099216c16afededa934713c7396aa7","observable_name":"sknknksndksnd","observable_classification":"generic","file_name":"","file_mimetype":"","status":"reported_with_fails","analyzers_requested":[],"connectors_requested":[],"analyzers_to_execute":["CryptoScamDB_CheckAPI","CRXcavator","CyberChef","Darksearch_Query","FileScan_Search","InQuest_IOCdb","InQuest_REPdb","InQuest_DFI","Phishstats","ThreatFox","YARAify_Generics","YETI"],"connectors_to_execute":["YETI"],"received_request_time":"2022-07-15T20:27:48.244790Z","finished_analysis_time":"2022-07-15T20:29:01.924146Z","tlp":"WHITE","errors":[]},{"id":72,"user":{"username":"hussain"},"tags":[],"process_time":87.87,"is_sample":false,"md5":"40ff44d9e619b17524bf3763204f9cbb","observable_name":"8.8.8.8","observable_classification":"ip","file_name":"","file_mimetype":"","status":"reported_with_fails","analyzers_requested":["Classic_DNS","CryptoScamDB_CheckAPI","Darksearch_Query","GoogleWebRisk","FileScan_Search","InQuest_IOCdb","GreyNoiseAlpha","GreyNoiseCommunity"],"connectors_requested":[],"analyzers_to_execute":["Classic_DNS","CryptoScamDB_CheckAPI","Darksearch_Query","GoogleWebRisk","FileScan_Search","InQuest_IOCdb","GreyNoiseCommunity"],"connectors_to_execute":["YETI"],"received_request_time":"2022-07-15T20:25:44.041286Z","finished_analysis_time":"2022-07-15T20:27:11.909898Z","tlp":"WHITE","errors":[]},{"id":71,"user":{"username":"hussain"},"tags":[],"process_time":null,"is_sample":false,"md5":"8fa14cdd754f91cc6554c9e71929cce7","observable_name":"f","observable_classification":"generic","file_name":"","file_mimetype":"","status":"killed","analyzers_requested":[],"connectors_requested":[],"analyzers_to_execute":["CryptoScamDB_CheckAPI","CRXcavator","CyberChef","Darksearch_Query","FileScan_Search","InQuest_IOCdb","InQuest_REPdb","InQuest_DFI","Phishstats","ThreatFox","YARAify_Generics","YETI"],"connectors_to_execute":["YETI"],"received_request_time":"2022-07-15T19:29:59.832445Z","finished_analysis_time":null,"tlp":"WHITE","errors":[]},{"id":70,"user":{"username":"hussain"},"tags":[],"process_time":32.67,"is_sample":false,"md5":"99754106633f94d350db34d548d6091a","observable_name":"fuck","observable_classification":"generic","file_name":"","file_mimetype":"","status":"reported_with_fails","analyzers_requested":[],"connectors_requested":[],"analyzers_to_execute":["CryptoScamDB_CheckAPI","CRXcavator","CyberChef","Darksearch_Query","FileScan_Search","InQuest_IOCdb","InQuest_REPdb","InQuest_DFI","Phishstats","ThreatFox","YARAify_Generics","YETI"],"connectors_to_execute":["YETI"],"received_request_time":"2022-07-15T19:29:10.045774Z","finished_analysis_time":"2022-07-15T19:29:42.713338Z","tlp":"WHITE","errors":[]},{"id":69,"user":{"username":"hussain"},"tags":[],"process_time":39.69,"is_sample":false,"md5":"40ff44d9e619b17524bf3763204f9cbb","observable_name":"8.8.8.8","observable_classification":"ip","file_name":"","file_mimetype":"","status":"reported_with_fails","analyzers_requested":[],"connectors_requested":[],"analyzers_to_execute":["Classic_DNS","CryptoScamDB_CheckAPI","Darksearch_Query","FireHol_IPList","FileScan_Search","GoogleWebRisk","GreyNoiseCommunity","InQuest_IOCdb","InQuest_REPdb","InQuest_DFI","MalwareBazaar_Google_Observable","Mnemonic_PassiveDNS","Phishstats","Pulsedive_Active_IOC","Robtex_IP_Query","Robtex_Reverse_PDNS_Query","Stratosphere_Blacklist","TalosReputation","ThreatFox","Threatminer_PDNS","Threatminer_Reports_Tagging","TorProject","URLhaus","UrlScan_Search","WhoIs_RipeDB_Search","YETI"],"connectors_to_execute":["YETI"],"received_request_time":"2022-07-15T19:26:47.638794Z","finished_analysis_time":"2022-07-15T19:27:27.325066Z","tlp":"WHITE","errors":[]},{"id":68,"user":{"username":"hussain"},"tags":[],"process_time":null,"is_sample":false,"md5":"40ff44d9e619b17524bf3763204f9cbb","observable_name":"8.8.8.8","observable_classification":"ip","file_name":"","file_mimetype":"","status":"killed","analyzers_requested":[],"connectors_requested":[],"analyzers_to_execute":["Classic_DNS","CryptoScamDB_CheckAPI","Darksearch_Query","FireHol_IPList","FileScan_Search","GoogleWebRisk","GreyNoiseCommunity","InQuest_IOCdb","InQuest_REPdb","InQuest_DFI","MalwareBazaar_Google_Observable","Mnemonic_PassiveDNS","Phishstats","Pulsedive_Active_IOC","Robtex_IP_Query","Robtex_Reverse_PDNS_Query","Stratosphere_Blacklist","TalosReputation","ThreatFox","Threatminer_PDNS","Threatminer_Reports_Tagging","TorProject","URLhaus","UrlScan_Search","WhoIs_RipeDB_Search","YETI"],"connectors_to_execute":["YETI"],"received_request_time":"2022-07-15T18:51:02.901898Z","finished_analysis_time":null,"tlp":"WHITE","errors":[]},{"id":67,"user":{"username":"hussain"},"tags":[],"process_time":82.88,"is_sample":false,"md5":"70047c83e6cfab6f85cf9fdf0cb4fdff","observable_name":"ransomware","observable_classification":"generic","file_name":"","file_mimetype":"","status":"reported_with_fails","analyzers_requested":[],"connectors_requested":[],"analyzers_to_execute":["CryptoScamDB_CheckAPI","CRXcavator","CyberChef","Darksearch_Query","FileScan_Search","InQuest_IOCdb","InQuest_REPdb","InQuest_DFI","Phishstats","ThreatFox","YARAify_Generics","YETI"],"connectors_to_execute":["YETI"],"received_request_time":"2022-07-15T18:33:10.719610Z","finished_analysis_time":"2022-07-15T18:34:33.601332Z","tlp":"WHITE","errors":[]}]}`
	jobList := gointelowl.JobListResponse{}
	err := json.Unmarshal([]byte(jobListJson), &jobList)
	if err != nil {
		t.Fatalf("Unexpected error - could not parse job list json")
	} else {
		// * table test case
		testCases := make(map[string]TestData)
		testCases["simple"] = TestData{
			Input:      nil,
			Data:       jobListJson,
			StatusCode: http.StatusOK,
			Want:       &jobList,
		}
		for name, testCase := range testCases {
			//* Subtest
			t.Run(name, func(t *testing.T) {
				client, apiHandler, closeServer := setup()
				defer closeServer()
				ctx := context.Background()
				apiHandler.Handle(constants.BASE_JOB_URL, serverHandler(t, testCase, "GET"))
				gottenJobList, err := client.JobService.List(ctx)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenJobList)
				}
			})
		}
	}
}

func TestJobServiceGet(t *testing.T) {
	jobJsonString := `{"id":72,"user":{"username":"hussain"},"tags":[],"process_time":87.87,"analyzer_reports":[{"name":"CryptoScamDB_CheckAPI","status":"SUCCESS","report":{"input":"8.8.8.8","result":{"type":"ip","status":"neutral","entries":[]},"success":true},"errors":[],"process_time":1.91,"start_time":"2022-07-15T20:25:45.681509Z","end_time":"2022-07-15T20:25:47.595517Z","runtime_configuration":{},"type":"analyzer"},{"name":"Darksearch_Query","status":"FAILED","report":{},"errors":["DarkSearchRequestException: "],"process_time":1.51,"start_time":"2022-07-15T20:25:45.681505Z","end_time":"2022-07-15T20:25:47.189974Z","runtime_configuration":{},"type":"analyzer"},{"name":"Classic_DNS","status":"SUCCESS","report":{"observable":"8.8.8.8","resolutions":["dns.google"]},"errors":[],"process_time":0.51,"start_time":"2022-07-15T20:25:47.601891Z","end_time":"2022-07-15T20:25:48.116356Z","runtime_configuration":{},"type":"analyzer"},{"name":"FileScan_Search","status":"SUCCESS","report":{"count":5,"items":[{"id":"bc6d3a15-b371-41f6-8b53-57333fc779be","date":"06/30/2022, 01:16:19","file":{"name":"test.bat","sha256":"b636ee9b411b5cc6ea5fae704f0889d05f509b9642574136f086c23220ce951a","mime_type":"application/x-bat","short_type":null},"tags":[],"state":"success_partial","matches":[{"origin":{"sha256":"b636ee9b411b5cc6ea5fae704f0889d05f509b9642574136f086c23220ce951a","filetype":null,"relation":"source","mime_type":"application/x-bat"},"matches":{"ip":[{"value":"8.8.8.8"}]}}],"verdict":"informational","scan_init":{"id":"62bcf6c787c294f96f8b67e7"},"updated_date":"06/30/2022, 01:16:58"},{"id":"7d849b04-3f66-42d3-9059-0c3f0d3d6af5","date":"06/30/2022, 01:12:01","file":{"name":"Anyrun_port80.bat","sha256":"2b0411d9f1bdc2c3906710bd62cb35288e50ee774f845a7665dbc86a0d6be40d","mime_type":"application/x-bat","short_type":"html"},"tags":[{"tag":{"name":"html","verdict":{"verdict":"INFORMATIONAL","confidence":1,"threatLevel":0.1},"synonyms":[],"descriptions":[]},"source":"MEDIA_TYPE","isRootTag":true,"sourceIdentifier":"2b0411d9f1bdc2c3906710bd62cb35288e50ee774f845a7665dbc86a0d6be40d"}],"state":"success_partial","matches":[{"origin":{"sha256":"2b0411d9f1bdc2c3906710bd62cb35288e50ee774f845a7665dbc86a0d6be40d","filetype":null,"relation":"source","mime_type":"application/x-bat"},"matches":{"ip":[{"value":"8.8.8.8"}]}}],"verdict":"informational","scan_init":{"id":"62bcf6cab0578633c8b24d36"},"updated_date":"06/30/2022, 01:21:09"},{"id":"aed83c14-0c79-44b4-96c4-7da063a6fd30","date":"06/30/2022, 01:05:21","file":{"name":"Anyrun_port80.bat","sha256":"2b0411d9f1bdc2c3906710bd62cb35288e50ee774f845a7665dbc86a0d6be40d","mime_type":"application/x-bat","short_type":"html"},"tags":[{"tag":{"name":"html","verdict":{"verdict":"INFORMATIONAL","confidence":1,"threatLevel":0.1},"synonyms":[],"descriptions":[]},"source":"MEDIA_TYPE","isRootTag":true,"sourceIdentifier":"2b0411d9f1bdc2c3906710bd62cb35288e50ee774f845a7665dbc86a0d6be40d"}],"state":"success_partial","matches":[{"origin":{"sha256":"2b0411d9f1bdc2c3906710bd62cb35288e50ee774f845a7665dbc86a0d6be40d","filetype":null,"relation":"source","mime_type":"application/x-bat"},"matches":{"ip":[{"value":"8.8.8.8"}]}}],"verdict":"informational","scan_init":{"id":"62bcf6be95a8514e298d7edd"},"retry_count":1,"updated_date":"07/01/2022, 01:53:30"},{"id":"4fc5a2a8-5554-4b7f-bb2a-8a0056629aa9","date":"01/02/2022, 00:34:33","file":{"name":"6e1d9a9c12395e4b505e1606cc0a6e26446412cd","sha256":"a258701294dffe74d811d173db94ec6cad2227c792a4233cd1dc2124544b9899","mime_type":"application/vnd.ms-excel.sheet.macroenabled.12","short_type":"xlsx"},"tags":[{"tag":{"name":"xlsx","verdict":{"verdict":"INFORMATIONAL","confidence":1,"threatLevel":0.1},"synonyms":[],"descriptions":[]},"source":"MEDIA_TYPE","isRootTag":true,"sourceIdentifier":"a258701294dffe74d811d173db94ec6cad2227c792a4233cd1dc2124544b9899"},{"tag":{"name":"html","verdict":{"verdict":"INFORMATIONAL","confidence":1,"threatLevel":0.1},"synonyms":[],"descriptions":[]},"source":"MEDIA_TYPE","isRootTag":true,"sourceIdentifier":"a258701294dffe74d811d173db94ec6cad2227c792a4233cd1dc2124544b9899"},{"tag":{"name":"fingerprint","verdict":{"verdict":"LIKELY_MALICIOUS","confidence":1,"threatLevel":0.75},"synonyms":[],"descriptions":[]},"source":"SIGNAL","isRootTag":true,"sourceIdentifier":"a258701294dffe74d811d173db94ec6cad2227c792a4233cd1dc2124544b9899"},{"tag":{"name":"stealer","verdict":{"verdict":"LIKELY_MALICIOUS","confidence":1,"threatLevel":0.75},"synonyms":[],"descriptions":[]},"source":"SIGNAL","isRootTag":true,"sourceIdentifier":"a258701294dffe74d811d173db94ec6cad2227c792a4233cd1dc2124544b9899"},{"tag":{"name":"macros","verdict":{"verdict":"INFORMATIONAL","confidence":1,"threatLevel":0.1},"synonyms":[],"descriptions":[]},"source":"SIGNAL","isRootTag":true,"sourceIdentifier":"a258701294dffe74d811d173db94ec6cad2227c792a4233cd1dc2124544b9899"}],"state":"success_partial","matches":[{"origin":{"sha256":"a258701294dffe74d811d173db94ec6cad2227c792a4233cd1dc2124544b9899","filetype":"xlsx","relation":"source","mime_type":"application/vnd.ms-excel.sheet.macroenabled.12"},"matches":{"ip":[{"value":"8.8.8.8"}]}}],"verdict":"suspicious","scan_init":{"id":"61d0f2da4ab5c44cbf5b1f85"},"updated_date":"01/02/2022, 00:38:53"},{"id":"595a3779-d327-4ac9-81b4-eb793479cae6","date":"08/23/2021, 01:28:03","file":{"name":"king.bat","sha256":"3f285b9294040d32eaaff486866372b79f1236bfb300d9821d20024142831f05","mime_type":"application/x-bat","short_type":null},"tags":[],"state":"success_partial","matches":[{"origin":{"sha256":"3f285b9294040d32eaaff486866372b79f1236bfb300d9821d20024142831f05","filetype":null,"relation":"source","mime_type":"application/x-bat"},"matches":{"ip":[{"value":"8.8.8.8"}]}}],"verdict":"informational","scan_init":{"id":"6122aec2d972e521533a7b28"},"updated_date":"08/23/2021, 01:28:03"}],"query":"OC44LjguOA==","method":"and","count_search_params":1},"errors":[],"process_time":8.41,"start_time":"2022-07-15T20:25:47.665641Z","end_time":"2022-07-15T20:25:56.079799Z","runtime_configuration":{},"type":"analyzer"},{"name":"GreyNoiseCommunity","status":"SUCCESS","report":{"ip":"8.8.8.8","link":"https://viz.greynoise.io/riot/8.8.8.8","name":"Google APIs and Services","riot":true,"noise":false,"message":"Success","last_seen":"2022-07-15","classification":"benign"},"errors":[],"process_time":1.4,"start_time":"2022-07-15T20:25:48.982345Z","end_time":"2022-07-15T20:25:50.385093Z","runtime_configuration":{},"type":"analyzer"},{"name":"InQuest_IOCdb","status":"SUCCESS","report":{"data":[],"success":true},"errors":["No API key retrieved"],"process_time":1.68,"start_time":"2022-07-15T20:25:49.007793Z","end_time":"2022-07-15T20:25:50.691163Z","runtime_configuration":{},"type":"analyzer"},{"name":"GoogleWebRisk","status":"FAILED","report":{},"errors":["/opt/deploy/intel_owl/configuration/service_account_keyfile.json should be an existing file. Check the docs on how to add this file to properly execute this analyzer"],"process_time":0.09,"start_time":"2022-07-15T20:27:11.329940Z","end_time":"2022-07-15T20:27:11.420128Z","runtime_configuration":{},"type":"analyzer"}],"connector_reports":[],"permissions":{"kill":true,"delete":true,"plugin_actions":true},"is_sample":false,"md5":"40ff44d9e619b17524bf3763204f9cbb","observable_name":"8.8.8.8","observable_classification":"ip","file_name":"","file_mimetype":"","status":"reported_with_fails","analyzers_requested":["Classic_DNS","CryptoScamDB_CheckAPI","Darksearch_Query","GoogleWebRisk","FileScan_Search","InQuest_IOCdb","GreyNoiseAlpha","GreyNoiseCommunity"],"connectors_requested":[],"analyzers_to_execute":["Classic_DNS","CryptoScamDB_CheckAPI","Darksearch_Query","GoogleWebRisk","FileScan_Search","InQuest_IOCdb","GreyNoiseCommunity"],"connectors_to_execute":["YETI"],"received_request_time":"2022-07-15T20:25:44.041286Z","finished_analysis_time":"2022-07-15T20:27:11.909898Z","tlp":"WHITE","errors":[]}`
	job := gointelowl.Job{}
	err := json.Unmarshal([]byte(jobJsonString), &job)
	if err != nil {
		t.Fatalf("Unexpected error - could not parse job list json")
		t.Fatalf("%v", err)
	} else {
		// *table test case
		testCases := make(map[string]TestData)
		testCases["simple"] = TestData{
			Input:      7,
			Data:       jobJsonString,
			StatusCode: http.StatusOK,
			Want:       &job,
		}
		testCases["cantFind"] = TestData{
			Input:      9000,
			Data:       `{"detail":"Not found."}`,
			StatusCode: http.StatusNotFound,
			Want: &gointelowl.IntelOwlError{
				StatusCode: http.StatusNotFound,
				Message:    `{"detail":"Not found."}`,
			},
		}
		for name, testCase := range testCases {
			//* Subtest
			t.Run(name, func(t *testing.T) {
				client, apiHandler, closeServer := setup()
				defer closeServer()
				ctx := context.Background()
				id, ok := testCase.Input.(int)
				if ok {
					jobId := uint64(id)
					testUrl := fmt.Sprintf(constants.SPECIFIC_JOB_URL, jobId)
					apiHandler.Handle(testUrl, serverHandler(t, testCase, "GET"))
					gottenJob, err := client.JobService.Get(ctx, jobId)
					if err != nil {
						testError(t, testCase, err)
					} else {
						testWantData(t, testCase.Want, gottenJob)
					}
				} else {
					t.Fatalf("Casting failed!")
				}
			})
		}
	}
}

func TestJobServiceDownloadSample(t *testing.T) {
	sampleString := "This is the sample"
	doesNotHaveASampleResponseJsonString := `{"errors":{"detail":"Requested job does not have a sample associated with it."}}`
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      1,
		Data:       sampleString,
		StatusCode: http.StatusOK,
		Want:       []byte(sampleString),
	}
	testCases["doesNotHaveASample"] = TestData{
		Input:      2,
		Data:       doesNotHaveASampleResponseJsonString,
		StatusCode: http.StatusBadRequest,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusBadRequest,
			Message:    doesNotHaveASampleResponseJsonString,
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			id, ok := testCase.Input.(int)
			if ok {
				jobId := uint64(id)
				testUrl := fmt.Sprintf(constants.DOWNLOAD_SAMPLE_JOB_URL, jobId)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "GET"))
				gottenSample, err := client.JobService.DownloadSample(ctx, jobId)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, gottenSample)
				}
			} else {
				t.Fatalf("Casting failed!")
			}
		})
	}
}

func TestJobServiceDelete(t *testing.T) {
	// *table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      1,
		Data:       "",
		StatusCode: http.StatusNoContent,
		Want:       true,
	}
	notFoundJson := `{"detail":"Not found."}`
	testCases["notFound"] = TestData{
		Input:      300,
		Data:       notFoundJson,
		StatusCode: http.StatusNotFound,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusNotFound,
			Message:    notFoundJson,
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			id, ok := testCase.Input.(int)
			if ok {
				jobId := uint64(id)
				testUrl := fmt.Sprintf(constants.SPECIFIC_JOB_URL, jobId)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "DELETE"))
				isDeleted, err := client.JobService.Delete(ctx, jobId)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, isDeleted)
				}
			}
		})
	}
}

func TestJobServiceKill(t *testing.T) {
	// *table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      1,
		Data:       "",
		StatusCode: http.StatusNoContent,
		Want:       true,
	}
	testCases["notFound"] = TestData{
		Input:      300,
		Data:       `{"detail":"Not found."}`,
		StatusCode: http.StatusNotFound,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusNotFound,
			Message:    `{"detail":"Not found."}`,
		},
	}
	testCases["jobNotRunning"] = TestData{
		Input:      71,
		Data:       `{"errors":{"detail":"Job is not running"}}`,
		StatusCode: http.StatusBadRequest,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusBadRequest,
			Message:    `{"errors":{"detail":"Job is not running"}}`,
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			id, ok := testCase.Input.(int)
			if ok {
				jobId := uint64(id)
				testUrl := fmt.Sprintf(constants.KILL_JOB_URL, jobId)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "PATCH"))
				isDeleted, err := client.JobService.Kill(ctx, jobId)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, isDeleted)
				}
			}
		})
	}
}

type input struct {
	Name string
	Id   uint64
}

func TestJobServiceKillAnalyzer(t *testing.T) {
	// *table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input: input{
			Name: "Phishstats",
			Id:   71,
		},
		Data:       "",
		StatusCode: http.StatusNoContent,
		Want:       true,
	}
	testCases["notFound"] = TestData{
		Input: input{
			Name: "NotAnAnalyzer",
			Id:   71,
		},
		Data:       `{"errors":{"analyzer report":"Not found."}}`,
		StatusCode: http.StatusNotFound,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusNotFound,
			Message:    `{"errors":{"analyzer report":"Not found."}}`,
		},
	}
	testCases["analyzerNotRunning"] = TestData{
		Input: input{
			Name: "Phishstats",
			Id:   71,
		},
		Data:       `{"errors":{"detail":"Plugin call is not running or pending"}}`,
		StatusCode: http.StatusBadRequest,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusBadRequest,
			Message:    `{"errors":{"detail":"Plugin call is not running or pending"}}`,
		},
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			inputData, ok := testCase.Input.(input)
			if ok {
				testUrl := fmt.Sprintf(constants.KILL_ANALYZER_JOB_URL, inputData.Id, inputData.Name)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "PATCH"))
				isDeleted, err := client.JobService.KillAnalyzer(ctx, inputData.Id, inputData.Name)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, isDeleted)
				}
			}
		})
	}
}

func TestJobServiceRetryAnalyzer(t *testing.T) {
	// *table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input: input{
			Name: "Phishstats",
			Id:   71,
		},
		Data:       "",
		StatusCode: http.StatusNoContent,
		Want:       true,
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			inputData, ok := testCase.Input.(input)
			if ok {
				testUrl := fmt.Sprintf(constants.RETRY_ANALYZER_JOB_URL, inputData.Id, inputData.Name)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "PATCH"))
				retry, err := client.JobService.RetryAnalyzer(ctx, inputData.Id, inputData.Name)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, retry)
				}
			}
		})
	}
}

func TestJobServiceKillConnector(t *testing.T) {
	// *table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input: input{
			Name: "YETI",
			Id:   71,
		},
		Data:       "",
		StatusCode: http.StatusNoContent,
		Want:       true,
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			inputData, ok := testCase.Input.(input)
			if ok {
				testUrl := fmt.Sprintf(constants.KILL_CONNECTOR_JOB_URL, inputData.Id, inputData.Name)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "PATCH"))
				isDeleted, err := client.JobService.KillConnector(ctx, inputData.Id, inputData.Name)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, isDeleted)
				}
			}
		})
	}
}

func TestJobServiceRetryConnector(t *testing.T) {
	// *table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input: input{
			Name: "YETI",
			Id:   71,
		},
		Data:       "",
		StatusCode: http.StatusNoContent,
		Want:       true,
	}
	for name, testCase := range testCases {
		//* Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			inputData, ok := testCase.Input.(input)
			if ok {
				testUrl := fmt.Sprintf(constants.RETRY_CONNECTOR_JOB_URL, inputData.Id, inputData.Name)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "PATCH"))
				retry, err := client.JobService.RetryConnector(ctx, inputData.Id, inputData.Name)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, retry)
				}
			}
		})
	}
}
