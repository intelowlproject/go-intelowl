> ⚠️ This repository is incomplete and not ready to be used.

# GoIntelOwl

IntelOwl client library/SDK written in [Go](https://golang.org/).


## Installation

```bash
$ go get github.com/intelowlproject/go-intelowl
```

## Using as a library / SDK

```go
import (
	"github.com/intelowlproject/go-intelowl"
)
client := gointelowl.IntelOwlClient{
		Token:       "<your_api_key>",
		URL:         "<your_intelowl_instance_url>",
		Certificate: "optional<path_to_pem_file>",
}
```

## FAQ

#### Generate API key
You need a valid API key to interact with the IntelOwl server. 
Keys should be created from the admin interface of [IntelOwl](https://github.com/intelowlproject/intelowl): you have to go in the *Durin* section (click on `Auth tokens`) and generate a key there.

## Checklist

- [ ] `/api/jobs`
- [ ] `/api/tags`
- [ ] send analysis request APIs
- [ ] `/api/ask_analysis_availability`
- [ ] Job actions like: `download_sample`, `kill`, `retry`
- [ ] `/api/get_analyzer_configs`
- [ ] `/api/get_connector_configs`

