package gointelowl

import (
	"context"
	"log"
	"reflect"
	"testing"
)

func TestIntelOwlClient(t *testing.T) {
	Token := ""
	URL := "http://localhost"
	client := CreateClient(
		Token,
		URL,
	)

	if client.BaseURL != URL {
		log.Fatalln("Expected Url : ", URL, " got : ", client.URL)
	}
	if client.Token != Token {
		log.Fatalln("Expected Token : ", Token, " got : ", client.Token)
	}

	if reflect.TypeOf(client).String() != "*gointelowl.IntelOwlClient" {
		log.Fatalln("Expected type: gointelowl.IntelOwlClient, but got : ", reflect.TypeOf(client).String())
	}
}

func TestGetAllJobs(t *testing.T) {
	Token := ""
	URL := "http://localhost"
	client := CreateClient(
		Token,
		URL,
	)

	ctx := context.Background()
	jobs, err := client.GetAllJobs(ctx)
	if err != nil {
		log.Fatalln("Error while testing GetAllJobs: ", err)
	}

	log.Println("ID of the first job is:", jobs[0].id)
	log.Println("GetAllJobs() works!")
}

func TestGetAllTags(t *testing.T) {
	Token := ""
	URL := "http://localhost"
	client := CreateClient(
		Token,
		URL,
	)

	ctx := context.Background()
	tags, err := client.GetAllTags(ctx)
	if err != nil {
		log.Fatalln("Error while testing GetAllJobs: ", err)
	}

	log.Println("ID of the first tag is:", tags[0].id)
	log.Println("GetAllTags() works!")

}
