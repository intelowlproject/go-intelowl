package gointelowl

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestIntelOwlClient(t *testing.T) {
	Token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ"
	URL := "http://localhost:8080"
	client := IntelOwlClient{
		Token:       Token,
		URL:         URL,
		Certificate: "",
	}
	if client.URL != URL {
		log.Fatalln("Expected Url : ", URL, " got : ", client.URL)
	}
	if client.Token != Token {
		log.Fatalln("Expected Token : ", Token, " got : ", client.Token)
	}

	if reflect.TypeOf(client).String() != "gointelowl.IntelOwlClient" {
		log.Fatalln("Expected type: gointelowl.IntelOwlClient, but got : ", reflect.TypeOf(client).String())
	}
}

func TestBuildAndMakeGetRequest(t *testing.T) {
	URL := "https://google.com"
	response := buildAndMakeGetRequest(URL, "")
	if response.StatusCode != http.StatusOK {
		log.Fatalln("Expected status code : ", http.StatusOK, " got : ", response.StatusCode)
	}
}

func TestBuildAndMakePostRequest(t *testing.T) {
	URL := "https://reqres.in/api/users"
	response := buildAndMakePostRequest(URL, "", []byte("{}"))
	if response.StatusCode != http.StatusCreated {
		log.Fatalln("Expected status code : ", http.StatusOK, " got : ", response.StatusCode)
	}
}

func TestGetAllJobs(t *testing.T) {
	client := IntelOwlClient{
		Token:       "95830471d0bc0c0595228f890c20beea",
		URL:         "http://localhost:80",
		Certificate: "",
	}
	jobs := client.GetAllJobs()
	fmt.Println(len(jobs))
	if len(jobs) == 0 {
		log.Fatalln("Expected jobs to be greater than 0")
	}
}
