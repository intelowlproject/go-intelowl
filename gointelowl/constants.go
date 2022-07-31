package gointelowl

//Tag endpoints URL
const (
	BASE_TAG_URL     = "%s/api/tags"
	SPECIFIC_TAG_URL = BASE_TAG_URL + "/%d"
)

// Job endpoints URL
const (
	BASE_JOB_URL            = "%s/api/jobs"
	SPECIFIC_JOB_URL        = BASE_JOB_URL + "/%d"
	DOWNLOAD_SAMPLE_JOB_URL = SPECIFIC_JOB_URL + "/download_sample"
	KILL_JOB_URL            = SPECIFIC_JOB_URL + "/kill"
	KILL_ANALYZER_JOB_URL   = SPECIFIC_JOB_URL + "/analyzer/%s/kill"
	RETRY_ANALYZER_JOB_URL  = SPECIFIC_JOB_URL + "/analyzer/%s/retry"
	KILL_CONNECTOR_JOB_URL  = SPECIFIC_JOB_URL + "/connector/%s/kill"
	RETRY_CONNECTOR_JOB_URL = SPECIFIC_JOB_URL + "/connector/%s/retry"
)

// Analyzer endpoints URL
const (
	ANALYZER_CONFIG_URL      = "%s/api/get_analyzer_configs"
	ANALYZER_HEALTHCHECK_URL = "%s/api/analyzer/%s/healthcheck"
)

// Connector endpoints URL
const (
	CONNECTOR_CONFIG_URL      = "%s/api/get_connector_configs"
	CONNECTOR_HEALTHCHECK_URL = "%s/api/connector/%s/healthcheck"
)

// Analysis endpoints URL
const (
	ANALYZE_OBSERVABLE_URL           = "%s/api/analyze_observable"
	ANALYZE_MULTIPLE_OBSERVABLES_URL = "%s/api/analyze_multiple_observables"
	ANALYZE_FILE_URL                 = "%s/api/analyze_file"
	ANALYZE_MULTIPLE_FILES_URL       = "%s/api/analyze_multiple_files"
)
