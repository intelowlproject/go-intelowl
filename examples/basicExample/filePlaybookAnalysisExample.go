package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/intelowlproject/go-intelowl/gointelowl"
	"github.com/sirupsen/logrus"
)

// Configure example.go() to run this funtion.

func FilePlaybookAnalysis() {
	// Configuring the Client!
	clientOptions := gointelowl.ClientOptions{
		Url:         "http://localhost:80",
		Token:       "feaed162aefa6ac35bdbbcb2b93c4bdfb5db88c0",
		Certificate: "",
		Timeout:     0,
	}

	loggerParams := &gointelowl.LoggerParams{
		File:      nil,
		Formatter: &logrus.JSONFormatter{},
		Level:     logrus.DebugLevel,
	}

	// Making the client!
	client := gointelowl.NewClient(
		&clientOptions,
		nil,
		loggerParams,
	)

	ctx := context.Background()

	basicAnalysisParams := gointelowl.BasicAnalysisParams{
		User:                 1,
		Tlp:                  gointelowl.WHITE,
		RuntimeConfiguration: map[string]interface{}{},
		AnalyzersRequested:   []string{},
		ConnectorsRequested:  []string{},
		TagsLabels:           []string{},
	}

	file, err := os.Open("exampleFiles/sample.jpeg")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	observableAnalysisParams := gointelowl.FilePlaybookAnalysisParams{
		BasicAnalysisParams: basicAnalysisParams,
		PlaybookRequested:   "Sample_Static_Analysis",
		File:                file,
	}

	analyzerResponse, err := client.CreateFilePlaybookAnalysis(ctx, &observableAnalysisParams)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	} else {
		analyzerResponseJSON, _ := json.Marshal(analyzerResponse)

		fmt.Println("========== ANALYZER RESPONSE ==========")
		fmt.Println(string(analyzerResponseJSON))
		fmt.Println("========== ANALYZER RESPONSE END ==========")
	}
}
