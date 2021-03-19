package gointelowl

import (
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
