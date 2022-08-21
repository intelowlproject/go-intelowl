package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/intelowlproject/go-intelowl/gointelowl"
	"github.com/sirupsen/logrus"
)

func main() {

	// Configuring the IntelOwlClient!
	clientOptions := gointelowl.IntelOwlClientOptions{
		Url:         "PUT-YOUR-INTELOWL-INSTANCE-URL-HERE",
		Token:       "PUT-YOUR-TOKEN-HERE",
		Certificate: "",
	}

	loggerParams := &gointelowl.LoggerParams{
		File:      nil,
		Formatter: &logrus.JSONFormatter{},
		Level:     logrus.DebugLevel,
	}

	// Making the client!
	client := gointelowl.NewIntelOwlClient(
		&clientOptions,
		nil,
		loggerParams,
	)

	ctx := context.Background()

	/*
		Now we can use the client to commnicate with your intelowl instance via the service objects!
		For this example I want to Display my tags list and create a new tag!
	*/

	fmt.Println("Getting the tag list!")

	// Getting the tag list!
	tagList, err := client.TagService.List(ctx)
	// checking for any pesky errors if there's any error it'll return an IntelOwlError
	if err != nil {
		fmt.Println(err)
	} else {
		// Iterating through the list unless its empty in that case create some using TagService.Create()!
		for _, tag := range *tagList {
			tagJson, err := json.Marshal(tag)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(tagJson))
			}
		}
	}

	// making the tag parameters!
	tagParams := gointelowl.TagParams{
		Label: "your super duper cool tag label!",
		Color: "#ffb703",
	}
	createdTag, err := client.TagService.Create(ctx, &tagParams)
	if err != nil {
		fmt.Println(err)
	} else {
		tagJson, err := json.Marshal(createdTag)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(tagJson))
		}
	}

}
