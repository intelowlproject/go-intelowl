package gointelowl

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestIntelOwlClient(t *testing.T) {
	Token := "ddaa4c2fdf53c213de4b862de1c4518e"
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
	Token := "ddaa4c2fdf53c213de4b862de1c4518e"
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

	for _, job := range jobs {
		fmt.Println(job.id) // Iterate over first job ID.
	}
	log.Println("GetAllJobs() works!")
}

func TestGetAllTags(t *testing.T) {
	Token := "ddaa4c2fdf53c213de4b862de1c4518e"
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

	for _, tag := range tags {
		fmt.Println(tag.id) // Iterate over first tag ID.
	}
	log.Println("GetAllTags() works!")

}
