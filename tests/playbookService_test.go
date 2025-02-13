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

func TestPlaybookServiceList(t *testing.T) {
	playbookListJson := `{
		"count": 2,
		"total_pages": 1,
		"results": [
			{
			"id": 10,
			"type": [
				"ip",
				"url",
				"domain"
			],
			"analyzers": [
				"CIRCLPassiveDNS",
				"DNSDB",
				"Mnemonic_PassiveDNS",
				"OTXQuery",
				"Robtex",
				"Threatminer",
				"Validin"
			],
			"connectors": [],
			"pivots": [],
			"visualizers": [
				"Passive_DNS"
			],
			"runtime_configuration": {
				"pivots": {},
				"analyzers": {},
				"connectors": {},
				"visualizers": {}
			},
			"scan_mode": 2,
			"scan_check_time": "1:00:00:00",
			"tags": [],
			"tlp": "AMBER",
			"weight": 15,
			"is_editable": false,
			"for_organization": false,
			"name": "Passive_DNS",
			"description": "A playbook that retrieve information from Passive DNS",
			"disabled": false,
			"starting": true,
			"owner": null
			},
			{
			"id": 1,
			"type": [
				"domain"
			],
			"analyzers": [
				"AdGuard",
				"Classic_DNS",
				"CloudFlare_DNS",
				"CloudFlare_Malicious_Detector",
				"DNS0_EU",
				"DNS0_EU_Malicious_Detector",
				"Google_DNS",
				"Quad9_DNS",
				"Quad9_Malicious_Detector",
				"UltraDNS_DNS",
				"UltraDNS_Malicious_Detector"
			],
			"connectors": [],
			"pivots": [],
			"visualizers": [
				"DNS"
			],
			"runtime_configuration": {
				"pivots": {},
				"analyzers": {},
				"connectors": {},
				"visualizers": {}
			},
			"scan_mode": 2,
			"scan_check_time": "1:00:00:00",
			"tags": [],
			"tlp": "AMBER",
			"weight": 8,
			"is_editable": false,
			"for_organization": false,
			"name": "Dns",
			"description": "Retrieve information from DNS about the domain",
			"disabled": false,
			"starting": true,
			"owner": null
			}
		]
	}`

	playbookList := gointelowl.PlaybookListResponse{}
	err := json.Unmarshal([]byte(playbookListJson), &playbookList)
	if err != nil {
		t.Fatalf("Unexpected error - could not parse playbook list JSON")
	}

	// * Table test case
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      nil,
		Data:       playbookListJson,
		StatusCode: http.StatusOK,
		Want:       &playbookList,
	}

	for name, testCase := range testCases {
		// * Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()

			ctx := context.Background()
			apiHandler.Handle(constants.BASE_PLAYBOOK_URL, serverHandler(t, testCase, "GET"))

			gottenPlaybookList, err := client.PlaybookService.ListPlaybooks(ctx)
			if err != nil {
				testError(t, testCase, err)
			} else {
				testWantData(t, testCase.Want, gottenPlaybookList)
			}
		})
	}
}

func TestGetPlaybookByName(t *testing.T) {

	playbookConfigJson := `{"id":1,"type":["domain"],"analyzers":["AdGuard","Classic_DNS","CloudFlare_DNS","CloudFlare_Malicious_Detector","DNS0_EU","DNS0_EU_Malicious_Detector","Google_DNS","Quad9_DNS","Quad9_Malicious_Detector","UltraDNS_DNS","UltraDNS_Malicious_Detector"],"connectors":[],"pivots":[],"visualizers":["DNS"],"runtime_configuration":{"pivots":{},"analyzers":{},"connectors":{},"visualizers":{}},"scan_mode":2,"scan_check_time":"1:00:00:00","tags":[],"tlp":"AMBER","weight":8,"is_editable":false,"for_organization":false,"name":"Dns","description":"Retrieve information from DNS about the domain","disabled":false,"starting":true,"owner":null}`
	playbookConfig := gointelowl.PlaybookConfig{}
	err := json.Unmarshal([]byte(playbookConfigJson), &playbookConfig)
	if err != nil {
		t.Fatalf("Unexpected error - could not parse playbook config JSON")
	}

	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      `Dns`,
		Data:       playbookConfigJson,
		StatusCode: http.StatusOK,
		Want:       &playbookConfig,
	}
	testCases["cantFind"] = TestData{
		Input:      "nonexistent",
		Data:       "404 page not found",
		StatusCode: http.StatusNotFound,
		Want: &gointelowl.Error{
			StatusCode: http.StatusNotFound,
			Message:    "404 page not found\n",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			// Mock the HTTP client and server
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			// Define the mock response
			url := fmt.Sprintf("%s%s", constants.BASE_PLAYBOOK_URL, testCase.Input)
			apiHandler.Handle(url, serverHandler(t, testCase, "GET"))
			gottenPlaybook, err := client.PlaybookService.GetPlaybookByName(ctx, string(testCase.Input.(string)))

			if err != nil {
				testError(t, testCase, err)
			} else {
				testWantData(t, testCase.Want, gottenPlaybook)
			}
		})
	}
}
