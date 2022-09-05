package constants

// These represent tag endpoints URL
const (
	BASE_TAG_URL     = "/api/tags"
	SPECIFIC_TAG_URL = BASE_TAG_URL + "/%d"
)

// These represent job endpoints URL
const (
	BASE_JOB_URL            = "/api/jobs"
	SPECIFIC_JOB_URL        = BASE_JOB_URL + "/%d"
	DOWNLOAD_SAMPLE_JOB_URL = SPECIFIC_JOB_URL + "/download_sample"
	KILL_JOB_URL            = SPECIFIC_JOB_URL + "/kill"
	KILL_ANALYZER_JOB_URL   = SPECIFIC_JOB_URL + "/analyzer/%s/kill"
	RETRY_ANALYZER_JOB_URL  = SPECIFIC_JOB_URL + "/analyzer/%s/retry"
	KILL_CONNECTOR_JOB_URL  = SPECIFIC_JOB_URL + "/connector/%s/kill"
	RETRY_CONNECTOR_JOB_URL = SPECIFIC_JOB_URL + "/connector/%s/retry"
)

// These represent analyzer endpoints URL
const (
	ANALYZER_CONFIG_URL      = "/api/get_analyzer_configs"
	ANALYZER_HEALTHCHECK_URL = "/api/analyzer/%s/healthcheck"
)

// These represent connector endpoints URL
const (
	CONNECTOR_CONFIG_URL      = "/api/get_connector_configs"
	CONNECTOR_HEALTHCHECK_URL = "/api/connector/%s/healthcheck"
)

// These represent analyze endpoints URL
const (
	ANALYZE_OBSERVABLE_URL           = "/api/analyze_observable"
	ANALYZE_MULTIPLE_OBSERVABLES_URL = "/api/analyze_multiple_observables"
	ANALYZE_FILE_URL                 = "/api/analyze_file"
	ANALYZE_MULTIPLE_FILES_URL       = "/api/analyze_multiple_files"
)

// These represent me endpoints URL

const (
	BASE_ME_URL                         = "/api/me"
	USER_DETAILS_URL                    = BASE_ME_URL + "/access"
	ORGANIZATION_URL                    = BASE_ME_URL + "/organization"
	INVITE_TO_ORGANIZATION_URL          = ORGANIZATION_URL + "/invite"
	REMOVE_MEMBER_FROM_ORGANIZATION_URL = ORGANIZATION_URL + "/remove_member"
)
