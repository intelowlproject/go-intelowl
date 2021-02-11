package gointelowl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Struct for IntelOwl.
// Initiate it using Token, Instance, Certificate
type IntelOwlClient struct {
	Token       string
	URL         string
	Certificate string
}

// Take url and token and place get request and return response or log the error if any and exit the program.
func buildAndMakeGetRequest(url string, token string) *http.Response {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Add("Authorization", "Token "+token)
	request.Header.Add("User-Agent", "IntelOwlClient/3.0.1")
	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(url, response.StatusCode)
	return response
}

// Take url, token and body, and place post request and return response or log the error if any and exit the program
func buildAndMakePostRequest(url string, token string, body []byte) *http.Response {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Add("Authorization", "Token "+token)
	request.Header.Add("User-Agent", "IntelOwlClient/3.0.1")
	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(url, response.StatusCode)
	return response
}

// Fetch list of all tags.
//        Endpoint: ``/api/tags``
// It Returns Slice of Tags
func (client *IntelOwlClient) GetAllTags() []map[string]string {
	url := client.URL + "/api/tags"
	response := buildAndMakeGetRequest(url, client.Token)
	var tags []map[string]string
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(data, &tags)
	if err != nil {
		log.Fatalln(err)
	}
	return tags
}

//  Fetch list of all jobs.
//		  Endpoint: ``/api/jobs``
//  Returns:
//	      []map[string]string: Slice of Jobs
func (client *IntelOwlClient) GetAllJobs() []map[string]string {
	url := client.URL + "/api/jobs"
	response := buildAndMakeGetRequest(url, client.Token)
	var jobs []map[string]string
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(data, &jobs)
	if err != nil {
		log.Fatalln(err)
	}
	return jobs
}

// Get current state of `analyzer_config.json` from the IntelOwl instance.
//        Endpoint: ``/api/get_analyzer_configs``
func (client *IntelOwlClient) GetAnalyzerConfigs() string {
	url := client.URL + "/api/get_analyzer_configs"
	response := buildAndMakeGetRequest(url, client.Token)
	analyzerConfig, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(analyzerConfig)
}
