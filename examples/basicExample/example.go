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
		Timeout:     0,
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

	basicAnalysisParams := gointelowl.BasicAnalysisParams{
		User:                 1,
		Tlp:                  gointelowl.WHITE,
		RuntimeConfiguration: map[string]interface{}{},
		AnalyzersRequested:   []string{},
		ConnectorsRequested:  []string{},
		TagsLabels:           []string{},
	}

	observableAnalysisParams := gointelowl.ObservableAnalysisParams{
		BasicAnalysisParams:      basicAnalysisParams,
		ObservableName:           "192.168.69.42",
		ObservableClassification: "ip",
	}

	analyzerResponse, err := client.CreateObservableAnalysis(ctx, &observableAnalysisParams)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	} else {
		analyzerResponseJSON, _ := json.Marshal(analyzerResponse)
		fmt.Println("JOB ID")
		fmt.Println(analyzerResponse.JobID)
		fmt.Println("JOB ID END")
		fmt.Println("========== ANALYZER RESPONSE ==========")
		fmt.Println(string(analyzerResponseJSON))
		fmt.Println("========== ANALYZER RESPONSE END ==========")
	}
}
