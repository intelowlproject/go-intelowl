package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/intelowlproject/go-intelowl/gointelowl"
)

/*
For this example I'll be using the tag params!
*/
func main() {
	// Configuring the IntelOwlClient!
	clientOptions := gointelowl.IntelOwlClientOptions{
		Url:         "PUT-YOUR-INTELOWL-INSTANCE-URL-HERE",
		Token:       "PUT-YOUR-TOKEN-HERE",
		Certificate: "",
	}

	// Making the client!
	client := gointelowl.NewIntelOwlClient(
		&clientOptions,
		nil,
	)

	ctx := context.Background()

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
