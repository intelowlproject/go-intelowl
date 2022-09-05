package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"testing"

	"github.com/intelowlproject/go-intelowl/constants"
	"github.com/intelowlproject/go-intelowl/gointelowl"
)

func TestConnectorServiceGetConfigs(t *testing.T) {
	connectorConfigJsonString := `{"MISP":{"name":"MISP","python_module":"misp.MISP","disabled":true,"description":"Automatically creates an event on your MISP instance, linking the successful analysis on IntelOwl","config":{"queue":"default","soft_time_limit":30},"secrets":{"api_key_name":{"env_var_key":"CONNECTOR_MISP_KEY","description":"API key for your MISP instance","required":true},"url_key_name":{"env_var_key":"CONNECTOR_MISP_URL","description":"URL of your MISP instance","required":true}},"params":{"ssl_check":{"value":true,"type":"bool","description":"Enable SSL certificate server verification. Change this if your MISP instance has not SSL enabled."},"debug":{"value":false,"type":"bool","description":"Enable debug logs."},"tlp":{"value":"white","type":"str","description":"Change this as per your organization's threat sharing conventions."}},"verification":{"configured":false,"error_message":"(api_key_name,url_key_name) not set; (0 of 2 satisfied)","missing_secrets":["api_key_name","url_key_name"]},"maximum_tlp":"WHITE"},"OpenCTI":{"name":"OpenCTI","python_module":"opencti.OpenCTI","disabled":true,"description":"Automatically creates an observable and a linked report on your OpenCTI instance, linking the successful analysis on IntelOwl","config":{"queue":"default","soft_time_limit":30},"secrets":{"api_key_name":{"env_var_key":"CONNECTOR_OPENCTI_KEY","description":"API key for your OpenCTI instance","required":true},"url_key_name":{"env_var_key":"CONNECTOR_OPENCTI_URL","description":"URL of your OpenCTI instance","required":true}},"params":{"ssl_verify":{"value":true,"type":"bool","description":"Enable SSL certificate server verification. Change this if your OpenCTI instance has not SSL enabled."},"proxies":{"value":{"http":"","https":""},"type":"dict","description":"Use these options to pass your request through a proxy server."},"tlp":{"value":{"type":"white","color":"#ffffff","x_opencti_order":1},"type":"dict","description":"Change this as per your organization's threat sharing conventions."}},"verification":{"configured":false,"error_message":"(api_key_name,url_key_name) not set; (0 of 2 satisfied)","missing_secrets":["api_key_name","url_key_name"]},"maximum_tlp":"WHITE"},"YETI":{"name":"YETI","python_module":"yeti.YETI","disabled":true,"description":"find or create observable on YETI, linking the successful analysis on IntelOwl.","config":{"queue":"default","soft_time_limit":30},"secrets":{"api_key_name":{"env_var_key":"CONNECTOR_YETI_KEY","description":"API key for your YETI instance","required":true},"url_key_name":{"env_var_key":"CONNECTOR_YETI_URL","description":"API URL of your YETI instance","required":true}},"params":{"verify_ssl":{"value":true,"type":"bool","description":"Enable SSL certificate server verification. Change this if your YETI instance has not SSL enabled."}},"verification":{"configured":false,"error_message":"(api_key_name,url_key_name) not set; (0 of 2 satisfied)","missing_secrets":["api_key_name","url_key_name"]},"maximum_tlp":"WHITE"}}`
	connectorConfigurationResponse := map[string]gointelowl.ConnectorConfig{}
	if unmarshalError := json.Unmarshal([]byte(connectorConfigJsonString), &connectorConfigurationResponse); unmarshalError != nil {
		t.Fatalf("Error: %s", unmarshalError)
	}
	connectorNames := make([]string, 0)
	// *getting all the analyzer key names!
	for connectorName := range connectorConfigurationResponse {
		connectorNames = append(connectorNames, connectorName)
	}
	// * sorting them alphabetically
	sort.Strings(connectorNames)
	connectorConfigurationList := []gointelowl.ConnectorConfig{}
	for _, connectorName := range connectorNames {
		connectorConfig := connectorConfigurationResponse[connectorName]
		connectorConfigurationList = append(connectorConfigurationList, connectorConfig)
	}
	// * table test cases
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      nil,
		Data:       connectorConfigJsonString,
		StatusCode: http.StatusOK,
		Want:       connectorConfigurationList,
	}
	for name, testCase := range testCases {
		// *Subtest
		t.Run(name, func(t *testing.T) {
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			apiHandler.Handle(constants.CONNECTOR_CONFIG_URL, serverHandler(t, testCase, "GET"))
			gottenConnectorConfigList, err := client.ConnectorService.GetConfigs(ctx)
			if err != nil {
				testError(t, testCase, err)
			} else {
				testWantData(t, testCase.Want, *gottenConnectorConfigList)
			}
		})
	}
}

func TestConnectorServiceHealthCheck(t *testing.T) {
	// * table test cases
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      "OpenCTI",
		Data:       `{"status": false}`,
		StatusCode: http.StatusOK,
		Want:       false,
	}
	testCases["connectorDoesntExist"] = TestData{
		Input:      "notAConnector",
		Data:       `{"errors": {"detail": "Connector doesn't exist"}}`,
		StatusCode: http.StatusBadRequest,
		Want: &gointelowl.IntelOwlError{
			StatusCode: http.StatusBadRequest,
			Message:    `{"errors": {"detail": "Connector doesn't exist"}}`,
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
				testUrl := fmt.Sprintf(constants.CONNECTOR_HEALTHCHECK_URL, input)
				apiHandler.Handle(testUrl, serverHandler(t, testCase, "GET"))
				status, err := client.ConnectorService.HealthCheck(ctx, input)
				if err != nil {
					testError(t, testCase, err)
				} else {
					testWantData(t, testCase.Want, status)
				}
			}
		})
	}
}
