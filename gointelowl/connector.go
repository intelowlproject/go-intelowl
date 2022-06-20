package gointelowl

type ConnectorConfiguration struct {
	BaseConfigurationType
	MaximumTlp string `json:"maximum_tlp"`
}
