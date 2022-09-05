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

func TestAnalyzerServiceGetConfigs(t *testing.T) {
	analyzerConfigJsonString := `{
		"APKiD_Scan_APK_DEX_JAR": {
			"name": "APKiD_Scan_APK_DEX_JAR",
			"python_module": "apkid.APKiD",
			"disabled": false,
			"description": "APKiD identifies many compilers, packers, obfuscators, and other weird stuff from an APK or DEX file.",
			"config": {
				"queue": "default",
				"soft_time_limit": 400
			},
			"secrets": {},
			"params": {},
			"verification": {
				"configured": true,
				"error_message": null,
				"missing_secrets": []
			},
			"type": "file",
			"external_service": false,
			"leaks_info": false,
			"docker_based": true,
			"run_hash": false,
			"supported_filetypes": [
				"application/zip",
				"application/java-archive",
				"application/vnd.android.package-archive",
				"application/x-dex"
			],
			"not_supported_filetypes": [],
			"observable_supported": []
		}
	}`
	serverErrorString := `{"error": "Error occurred by the server"}`
	badGatewayErrorString := `{"code": 502,"message": "Bad Gateway"}`
	analyzerConfigurationResponse := map[string]gointelowl.AnalyzerConfig{}
	if unmarshalError := json.Unmarshal([]byte(analyzerConfigJsonString), &analyzerConfigurationResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	analyzerConfigurationList := []gointelowl.AnalyzerConfig{}
	for _, analyzerConfig := range analyzerConfigurationResponse {
		analyzerConfigurationList = append(analyzerConfigurationList, analyzerConfig)
	}
	// * table test cases
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      nil,
		Data:       analyzerConfigJsonString,
		StatusCode: http.StatusOK,
		Want:       analyzerConfigurationList,
	}
	testCases["serverError"] = TestData{
		Input:      nil,
		Data:       serverErrorString,
		StatusCode: http.StatusInternalServerError,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusInternalServerError,
			Message:    serverErrorString,
		},
	}
	testCases["badGateway"] = TestData{
		Input:      nil,
		Data:       badGatewayErrorString,
		StatusCode: http.StatusBadGateway,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusBadGateway,
			Message:    badGatewayErrorString,
		},
	}
	for name, testCase := range testCases {
		// *Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			apiHandler.Handle(constants.ANALYZER_CONFIG_URL, serverHandler(t, testCase, "GET"))
			gottenAnalyzerConfigList, err := client.AnalyzerService.GetConfigs(ctx)
			if err != nil {
				testError(t, testCase, err)
			} else {
				testWantData(t, testCase.Want, (*gottenAnalyzerConfigList))
			}
		})
	}
}

func TestAnalyzerServiceHealthCheck(t *testing.T) {
	// * table test cases
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      "Floss",
		Data:       `{"status": true}`,
		StatusCode: http.StatusOK,
		Want:       true,
	}
	testCases["analyzerDoesntExist"] = TestData{
		Input:      "notAnAnalyzer",
		Data:       `{"errors": {"detail": "Analyzer doesn't exist"}}`,
		StatusCode: http.StatusBadRequest,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusBadRequest,
			Message:    `{"errors": {"detail": "Analyzer doesn't exist"}}`,
		},
	}
	for name, testCase := range testCases {
		// *Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			input, ok := testCase.Input.(string)
			if ok {
				testUrl := fmt.Sprintf(constants.ANALYZER_HEALTHCHECK_URL, input)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "GET"))
				status, err := client.AnalyzerService.HealthCheck(ctx, input)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, status)
				}
			}
		})
	}
}
