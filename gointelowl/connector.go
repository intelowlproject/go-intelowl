package gointelowl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type ConnectorConfig struct {
	BaseConfigurationType
	MaximumTlp TLP `json:"maximum_tlp"`
}

type ConnectorService struct {
	client *IntelOwlClient
}

func (connectorService *ConnectorService) GetConfigs(ctx context.Context) (*[]ConnectorConfig, error) {
	requestUrl := fmt.Sprintf("%s/api/get_connector_configs", connectorService.client.options.Url)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	successResp, err := connectorService.client.newRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	connectorConfigurationResponse := map[string]ConnectorConfig{}
	if unmarshalError := json.Unmarshal(successResp.Data, &connectorConfigurationResponse); unmarshalError != nil {
		return nil, unmarshalError
	}

	connectorNames := make([]string, 0)
	// *getting all the analyzer key names!
	for connectorName := range connectorConfigurationResponse {
		connectorNames = append(connectorNames, connectorName)
	}
	// * sorting them alphabetically
	sort.Strings(connectorNames)
	connectorConfigurationList := []ConnectorConfig{}
	for _, connectorName := range connectorNames {
		connectorConfig := connectorConfigurationResponse[connectorName]
		connectorConfigurationList = append(connectorConfigurationList, connectorConfig)
	}
	return &connectorConfigurationList, nil
}

func (connectorService *ConnectorService) HealthCheck(ctx context.Context, connectorName string) (bool, error) {
	requestUrl := fmt.Sprintf("%s/api/connector/%s/healthcheck", connectorService.client.options.Url, connectorName)
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return false, err
	}
	status := StatusResponse{}
	successResp, err := connectorService.client.newRequest(ctx, request)
	if err != nil {
		return false, err
	}
	if unmarshalError := json.Unmarshal(successResp.Data, &status); unmarshalError != nil {
		return false, unmarshalError
	}
	return status.Status, nil
}
