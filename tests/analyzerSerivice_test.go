package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/intelowlproject/go-intelowl/gointelowl"
)

type AnalyzerTest struct {
	input      interface{}
	data       string
	statusCode int
	want       interface{}
}

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
	analyzerConfigurationResponse := map[string]gointelowl.AnalyzerConfiguration{}
	if unmarshalError := json.Unmarshal([]byte(analyzerConfigJsonString), &analyzerConfigurationResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	analyzerConfigurationList := []gointelowl.AnalyzerConfiguration{}
	for _, analyzerConfig := range analyzerConfigurationResponse {
		analyzerConfigurationList = append(analyzerConfigurationList, analyzerConfig)
	}
	// * table test cases
	testCases := make(map[string]AnalyzerTest)
	testCases["simple"] = AnalyzerTest{
		input:      nil,
		data:       analyzerConfigJsonString,
		statusCode: http.StatusOK,
		want:       analyzerConfigurationList,
	}
	for name, testCase := range testCases {
		// *Subtest
		t.Run(name, func(t *testing.T) {
			testData := TestData{
				StatusCode: testCase.statusCode,
				Data:       testCase.data,
			}
			testServer := NewTestServer(&testData)
			defer testServer.Close()
			client := NewTestIntelOwlClient(testServer.URL)
			ctx := context.Background()
			gottenAnalyzerConfigList, err := client.AnalyzerService.GetConfigs(ctx)
			if err != nil {
				t.Fatalf("Error listing tags: %v", err)
			}
			diff := cmp.Diff(testCase.want, (*gottenAnalyzerConfigList))
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
