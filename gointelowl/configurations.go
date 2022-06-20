package gointelowl

type ConfigType struct {
	Queue         string `json:"queue"`
	SoftTimeLimit int    `json:"soft_time_limit"`
}

type Secret struct {
	EnvironmentVariableKey string `json:"env_var_key"`
	Description            string `json:"description"`
	Required               bool   `json:"required"`
}

type Parameter struct {
	Value       string `json:"value"`
	Type        int    `json:"type"`
	Description string `json:"description"`
}

type VerificationType struct {
	Configured     bool     `json:"configured"`
	ErrorMessage   string   `json:"error_message"`
	MissingSecrets []string `json:"missing_secrets"`
}

type BaseConfigurationType struct {
	Name         string                      `json:"name"`
	PythonModule string                      `json:"python_module"`
	Disabled     bool                        `json:"disabled"`
	Description  string                      `json:"description"`
	Config       ConfigType                  `json:"config"`
	Secrets      map[string]Secret           `json:"secrets"`
	Params       map[string]Parameter        `json:"params"`
	Verification map[string]VerificationType `json:"verification"`
}
